package utils_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/utils"
)

func TestSendJsonError(t *testing.T) {
	recorder := httptest.NewRecorder()

	message := "Sample error message"
	statusCode := http.StatusBadGateway
	utils.SendJsonError(recorder, message, statusCode)

	if recorder.Code != statusCode {
		t.Errorf("Expected status code %d, got %d", statusCode, recorder.Code)
	}

	contentType := recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}

	var response models.ErrorResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshaling json response: %v", err)
	}

	if response.Error != message {
		t.Errorf("Expected error message %q, got %q", message, response)
	}
}

func TestSendJsonSuccess(t *testing.T) {
	recorder := httptest.NewRecorder()

	data := models.Product{
		Name:        "Sample product name",
		Stock:       123,
		Price:       13467.00,
		Description: "Product description",
	}
	statusCode := http.StatusOK
	utils.SendJsonSuccess(recorder, data, statusCode)

	if recorder.Code != statusCode {
		t.Errorf("Expected status code %d, got %d", statusCode, recorder.Code)
	}

	contentType := recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}

	var response models.Product
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshaling json response: %v", err)
	}

	if !(response.Name == data.Name && response.Stock == data.Stock && response.Price == data.Price && response.Description == data.Description) {
		t.Errorf("Expected error message %+v, got %+v", data, response)
	}
}
