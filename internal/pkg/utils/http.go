package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
)

func SendJsonError(w http.ResponseWriter, message string, statusCode int) {
	response := models.ErrorResponse{
		Error: message,
	}
	json, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error marshal json error response: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
}

func SendJsonSuccess(w http.ResponseWriter, response any, statusCode int) {
	json, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshal json success response: %v", err)
		http.Error(w, "internal server errro", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
}
