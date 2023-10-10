package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/golang-jwt/jwt/v4"

	storage "github.com/notaduck/fitter-go/db"
	"github.com/notaduck/fitter-go/models"
	"github.com/notaduck/fitter-go/mq"
)

// APIServer provides an API server for handling activity-related requests.
type APIServer struct {
	listenAddr string
	storage    storage.Storage
	mq         mq.RabbitMQ
}

// NewAPIServer creates a new instance of APIServer.
func NewAPIServer(listenAddr string, storage storage.Storage, mq mq.RabbitMQ) *APIServer {
	// func NewAPIServer(listenAddr string, storage storage.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		storage:    storage,
		mq:         mq,
	}
}

// Run starts the API server.
func (s *APIServer) Run() {

	// http.HandleFunc("/activity", (makeHTTPHandleFunc(s.handleActivity)))
	http.HandleFunc("/user", makeHTTPHandleFunc(s.handleUser))
	http.HandleFunc("/login", makeHTTPHandleFunc(s.handleLogin))
	http.HandleFunc("/private", EnsureValidToken()(makeHTTPHandleFunc(s.handlePrivate)))
	http.HandleFunc("/activity", EnsureValidToken()(makeHTTPHandleFunc(s.handleActivity)))
	// CORS middleware

	// Start the server with the CORS middleware
	log.Println("JSON API server running on port:", ":3030")

	err := http.ListenAndServe(":3030", nil)

	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func (s *APIServer) handlePrivate(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "POST" {

		user := s.getUserFromContext(r)
		fmt.Println(user)
		return WriteJSON(w, http.StatusAccepted, 1337)
	}

	return fmt.Errorf("method not allowed %s", r.Method)

}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {

	// Publish a message

	// Check the HTTP request method, only allow POST requests
	if r.Method != "POST" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	// Create a variable to hold the request data
	var req models.LoginRequest

	// Decode the JSON request body into the 'req' variable
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// If there's an error decoding the JSON, return it as an error
		return err
	}

	// Check if required fields are present in the 'req' object

	// simple and naive req validation.
	if req.Email == "" || req.Password == "" {
		return fmt.Errorf("both 'Number' and 'Password' are required")
	}

	// You can add more specific validation logic here if needed,
	// such as checking the range of 'req.Number' or other conditions.

	// Retrieve the user account based on the 'req.Number'
	acc, err := s.storage.GetUserByEmail(req.Email)
	if err != nil {
		// If there's an error retrieving the user account, return it as an error
		return err
	}

	// Check if the provided password matches the user's password
	if !acc.ValidPassword(req.Password) {
		// If the password doesn't match, return an authentication error
		return fmt.Errorf("not authenticated")
	}

	// Create a JWT token for the authenticated user
	token, err := createJWT(acc)
	if err != nil {
		// If there's an error creating the JWT token, return it as an error
		return err
	}

	// Prepare the response with the JWT token and user ID
	resp := models.LoginResponse{
		Token:  token,
		Number: int64(acc.ID),
	}

	// Write the response as JSON with an HTTP status code of 200 (OK)
	return WriteJSON(w, http.StatusOK, resp)
}

// apiFunc represents a function that handles API requests.
type apiFunc func(http.ResponseWriter, *http.Request) error

// ApiError represents an API error response.
type ApiError struct {
	Error string `json:"error"`
}

// WriteJSON writes a JSON response with the given status code and data.
func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// makeHTTPHandleFunc creates an HTTP handler function from an API function.
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func createJWT(user *models.User) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"userId":    user.ID,
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func permissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
}

func withJWTAuth(handlerFunc http.HandlerFunc, s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type jwtMapClaim struct {
			jwt.MapClaims
			userId    int
			expiresAt int16
		}
		fmt.Println("calling JWT auth middleware")

		tokenString := r.Header.Get("x-jwt-token")
		token, err := validateJWT(tokenString)
		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			fmt.Errorf("couldn't extract claims")
			permissionDenied(w)
		}

		userId := int(claims["userId"].(float64))

		_ = userId

		if err != nil {
			permissionDenied(w)
			return
		}
		if !token.Valid {
			permissionDenied(w)
			return
		}

		user, err := s.GetUserById(userId)
		if err != nil {
			permissionDenied(w)
			return
		}

		if err != nil {
			WriteJSON(w, http.StatusForbidden, ApiError{Error: "invalid token"})
			return
		}

		// Create a context with the user information
		ctx := context.WithValue(r.Context(), "user", user)

		// Replace the request's context with the new context containing user information
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func (a *APIServer) getUserFromContext(r *http.Request) *models.User {

	claims := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	user, err := a.storage.GetUserByAuth0Id(claims.RegisteredClaims.Subject)

	if err != nil {
		log.Println("No d found found")
	}
	return user
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := models.JwtClaims{}
	_ = claims

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
}

func readMultipartFileToBytes(fileHeader *multipart.FileHeader) ([]byte, error) {
	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content into a byte slice
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

// =============================
// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope        string `json:"scope"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	ShouldReject bool   `json:"shouldReject,omitempty"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func EnsureValidToken() func(http.HandlerFunc) http.HandlerFunc {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{"https://localhost:3030", "2CN35CezPd8oNsAKupxsO9sW2VC8SmbH"},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)

	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			middleware.CheckJWT(http.HandlerFunc(next)).ServeHTTP(w, r)
		}
	}
}

func inspectUser(user interface{}) {
	// Use reflection to get the type of the user object
	userType := reflect.TypeOf(user)

	// Check if the user is a pointer and get the underlying type
	if userType.Kind() == reflect.Ptr {
		userType = userType.Elem()
	}

	// Print the name of the type
	fmt.Printf("Type: %s\n", userType.Name())

	// Iterate through the fields of the type
	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		fmt.Printf("Field Name: %s, Field Type: %s\n", field.Name, field.Type)
	}
}
