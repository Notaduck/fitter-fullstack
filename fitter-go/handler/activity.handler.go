package handler

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/notaduck/fitter-go/models"
)

func (s *APIServer) handleActivity(w http.ResponseWriter, r *http.Request) error {
	enableCors(&w)
	if r.Method == "GET" {
		return s.getActivities(w, r)
	}

	if r.Method == "POST" {
		return s.createActivities(w, r)
	}
	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) getActivities(w http.ResponseWriter, r *http.Request) error {

	user := s.getUserFromContext(r)

	fmt.Println("user", user)

	id := r.URL.Query().Get("id")

	if id != "" {

		ac := models.GetActivityDTO{}

		num, err := strconv.Atoi(id)

		if err != nil {
			// log.Fatalln("The id must be a valid number")
			return fmt.Errorf("The id [%s] must be a valid number.", id)
		}

		log.Default().Println("GET A SINGLE ACTIVITY WITH", num)
		activity, err := s.storage.GetActivity(1, num)

		ac.ID = activity.ID
		ac.Distance = activity.Distance
		ac.Elevation = activity.Elevation
		ac.Timestamp = activity.Timestamp
		ac.TotalRideTime = activity.TotalRideTime

		if err != nil {
			return err
		}

		records, err := s.storage.GetRecordMsgs(activity.ID)

		ac.Records = records

		if err != nil {
			return err
		}

		return WriteJSON(w, http.StatusOK, ac)

	}

	activities, err := s.storage.GetActivities(1)

	if err != nil {
		fmt.Println(err)
	}

	return WriteJSON(w, http.StatusOK, activities)

}

func (s *APIServer) createActivities(w http.ResponseWriter, r *http.Request) error {

	user := s.getUserFromContext(r)

	fmt.Println("user", user)

	err := r.ParseMultipartForm(10 << 20) // 10 MB limit for uploaded files

	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, "Please select one or more .fit files to upload and try again! 📁🚀")
	}

	files := r.MultipartForm.File["files"]

	for _, fileHeader := range files {
		// Open the file

		file, err := fileHeader.Open()

		if err != nil {
			http.Error(w, "Unable to open file", http.StatusInternalServerError)
			return err
		}
		defer file.Close()

		// Check if the file has a .fit extension
		if filepath.Ext(fileHeader.Filename) != ".fit" {
			http.Error(w, "Invalid file type. Only .fit files are allowed.", http.StatusBadRequest)
			return err
		}

		if err != nil {
			log.Fatal(err)
		}

		type FitterBody struct {
			FitFile []byte
			Size    int64
			UserId  int
		}

		bs := make([]byte, fileHeader.Size)

		_, err = bufio.NewReader(file).Read(bs)

		if err != nil && err != io.EOF {
			fmt.Println(err)
			return &json.SyntaxError{}
		}

		messageData := FitterBody{
			FitFile: bs,
			Size:    fileHeader.Size,
			UserId:  1,
		}

		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(messageData)

		err = s.mq.Publish("fit_queue", reqBodyBytes.Bytes())

		if err != nil {
			fmt.Println(err)

		}

	}

	if err != nil {
		panic(err)
	}

	_ = files

	return WriteJSON(w, http.StatusOK, 1337)
}
