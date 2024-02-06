package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/repositories"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/utils"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/validator"
)

func Carts(w http.ResponseWriter, r *http.Request) {
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

func Cart(w http.ResponseWriter, r *http.Request) {
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

func Checkout(w http.ResponseWriter, r *http.Request) {
}
