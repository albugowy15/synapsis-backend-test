package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/repositories"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/utils"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/validator"
	"github.com/go-chi/chi/v5"
)

// Get all products in cart
// @Tags carts v1
// @Summary Get all products in cart
// @Description Get all products in cart
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce json
// @Success 200 {object} models.ArrayResponse
// @Success 400 {object} models.ErrorResponse
// @Success 404 {object} models.ErrorResponse
// @Success 500 {object} models.ErrorResponse
// @Router /carts [get]
func GetCarts(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetJwtClaim(r)
	if err != nil {
		utils.SendJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}
	s := repositories.GetCartRepository()
	carts, err := s.GetByUserId(userId)
	if err != nil {
		log.Fatalf("Error get carts for user %d: %v", userId, err)
		utils.SendJsonError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	response := models.ArrayResponse{
		TotalItems: len(carts),
		Data:       carts,
	}
	utils.SendJsonSuccess(w, response, http.StatusOK)
}

// Add product to cart
// @Tags carts v1
// @Summary Add product to cart
// @Description Add product to cart
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param message body models.AddCartRequest true "Add to Cart Request"
// @Produce json
// @Success 201 {object} models.MessageResponse
// @Success 400 {object} models.ErrorResponse
// @Success 404 {object} models.ErrorResponse
// @Success 500 {object} models.ErrorResponse
// @Router /carts [post]
func AddCart(w http.ResponseWriter, r *http.Request) {
	var body models.AddCartRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println("error parsing body: ", err)
		utils.SendJsonError(w, "error parsing body", http.StatusBadRequest)
		return
	}

	if err := validator.ValidateAddToCartRequest(body); err != nil {
		utils.SendJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := utils.GetJwtClaim(r)
	if err != nil {
		log.Println("error parsing token: ", err)
		utils.SendJsonError(w, "error parsing token", http.StatusBadRequest)
		return
	}

	p := repositories.GetProductRepository()
	_, err = p.GetById(body.ProductId)
	if err != nil && err == sql.ErrNoRows {
		log.Println("product not exist: ", err)
		utils.SendJsonError(w, "product not exist", http.StatusNotFound)
		return
	}
	s := repositories.GetCartRepository()
	isProductInCart := s.CheckProductExistInCart(body.ProductId, userId)
	if isProductInCart {
		utils.SendJsonError(w, "this product already in your shopping cart", http.StatusBadRequest)
		return
	}
	isStockAvailable, err := s.CheckStockAvailable(body.ProductId, body.Quantity)
	if err != nil {
		log.Println("check stock error: ", err)
		utils.SendJsonError(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if !isStockAvailable {
		utils.SendJsonError(w, "product stock is less than the desired quantity", http.StatusBadRequest)
		return
	}

	err = s.Add(body.ProductId, body.Quantity, userId)
	if err != nil {
		log.Println("add cart error: ", err)
		utils.SendJsonError(w, "internal server error", http.StatusInternalServerError)
		return
	}
	utils.SendJsonSuccess(w, map[string]interface{}{"message": "succcessfully add product to shopping cart"}, http.StatusCreated)
}

// Delete product in cart
// @Tags carts v1
// @Summary Delete product in cart
// @Description Delete product in cart
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param product_id path int true "Product ID"
// @Produce json
// @Success 201 {object} models.MessageResponse
// @Success 400 {object} models.ErrorResponse
// @Success 404 {object} models.ErrorResponse
// @Success 500 {object} models.ErrorResponse
// @Router /carts/{product_id} [delete]
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetJwtClaim(r)
	if err != nil {
		log.Printf("Error get userId from jwt: %v", err)
		utils.SendJsonError(w, "error parsing token", http.StatusBadRequest)
		return
	}
	productIdStr := chi.URLParam(r, "product_id")
	if productIdStr == "" {
		utils.SendJsonError(w, "missing product_id in url params", http.StatusBadRequest)
	}
	productId, err := strconv.ParseInt(productIdStr, 10, 64)
	if err != nil {
		log.Printf("Error convert product id from string into int64: %v", err)
		utils.SendJsonError(w, "product_id must integer", http.StatusBadRequest)
		return
	}

	// do delete
	// validate product id exist
	s := repositories.GetCartRepository()
	isProductInCart := s.CheckProductExistInCart(productId, userId)
	if !isProductInCart {
		utils.SendJsonError(w, "product id is not in your cart", http.StatusBadRequest)
		return
	}

	p := repositories.GetProductRepository()
	_, err = p.GetById(productId)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.SendJsonError(w, fmt.Sprintf("product with id %d not found", productId), http.StatusBadRequest)
			return
		}
		log.Println("error get product id ", err)
		utils.SendJsonError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := s.DeleteProductFromCart(productId, userId); err != nil {
		utils.SendJsonError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	utils.SendJsonSuccess(w, map[string]interface{}{"message": "product succcessfully deleted from shopping cart"}, http.StatusOK)
}
