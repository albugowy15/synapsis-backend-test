package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/repositories"
)

// Get All Products
// @Tags products v1
// @Summary Get all products information of filter with category
// @Description Get all products information of filter with category
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func Products(w http.ResponseWriter, r *http.Request) {
	s := repositories.GetProductRepository()
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
}
