package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/notaduck/fitter-go/models"
)

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateUser(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {

	// err := s.mq.Publish("", "my_exchange", "Hello, RabbitMQ2!")
	// if err != nil {
	// 	log.Fatalf("Failed to publish message: %v", err)
	// }

	return nil
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {

	req := new(models.CreateUserDTO)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	user, err := models.NewUser(req.FirstName, req.LastName, req.Password, req.Username, req.Email)

	if err != nil {
		return err
	}

	if err := s.storage.CreateUser(user); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}
