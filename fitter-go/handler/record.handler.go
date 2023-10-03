package handler

import (
	"fmt"
	"net/http"
)

func (s *APIServer) recordHandler(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return s.getRecords(w, r)
	}

	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) getRecords(w http.ResponseWriter, r *http.Request) error {

	return nil

}
