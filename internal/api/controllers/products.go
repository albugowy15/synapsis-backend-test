package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/repositories"
)

func Products(w http.ResponseWriter, r *http.Request) {
	s := repositories.GetProductRepository()
	switch r.Method {
	case http.MethodGet:
		query := r.URL.Query()
		categoryParam := query.Get("category")
		var products []models.Product
		var err error
		if categoryParam != "" {
			products, err = s.GetByCategory(categoryParam)
			if err != nil {
				log.Printf("Error find products with category %s: %v", categoryParam, err)
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
		}

		products, err = s.All()
		if err != nil {
			log.Printf("Error get all products: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(products)
		if err != nil {
			log.Printf("Error marshal response json: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
		return

	case http.MethodPost:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	case http.MethodPut:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	case http.MethodDelete:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
