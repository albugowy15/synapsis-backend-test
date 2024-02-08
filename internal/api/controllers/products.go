package controllers

import (
	"log"
	"net/http"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/repositories"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/utils"
)

// Get All Products
// @Tags products v1
// @Summary Get all products information of filter with category
// @Description Get all products information of filter with category
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Success 400 {object} models.ErrorResponse
// @Success 404 {object} models.ErrorResponse
// @Success 500 {object} models.ErrorResponse
// @Router /products [get]
func GetProducts(w http.ResponseWriter, r *http.Request) {
	s := repositories.GetProductRepository()
	query := r.URL.Query()
	categoryParam := query.Get("category")
	var err error
	if categoryParam != "" {
		var products []models.Product
		products, err = s.GetByCategory(categoryParam)
		if err != nil {
			log.Printf("Error find products with category %s: %v", categoryParam, err)
			utils.SendJsonError(w, err.Error(), http.StatusNotFound)
			return
		}
		response := models.ArrayResponse{
			TotalItems: len(products),
			Data:       products,
		}
		utils.SendJsonSuccess(w, response, http.StatusOK)
		return
	}

	var products []models.Product
	products, err = s.All()
	if err != nil {
		log.Printf("Error get all products: %v", err)
		utils.SendJsonError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response := models.ArrayResponse{
		TotalItems: len(products),
		Data:       products,
	}
	utils.SendJsonSuccess(w, response, http.StatusOK)
}
