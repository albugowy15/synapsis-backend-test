package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/repositories"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/utils"
)

// Checkout cart
// @Tags carts v1
// @Summary Checkout cart
// @Description Checkout all product in cart and create transaction
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param message body models.CheckoutRequest true "choouse payment method when checkout"
// @Produce json
// @Success 201 {object} models.MessageResponse
// @Success 400 {object} models.ErrorResponse
// @Success 404 {object} models.ErrorResponse
// @Success 500 {object} models.ErrorResponse
// @Router /carts/checkout [post]
func Checkout(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetJwtClaim(r)
	if err != nil {
		utils.SendJsonError(w, "error parsing token", http.StatusBadRequest)
		return
	}
	var body models.CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendJsonError(w, "error parsing request body", http.StatusBadRequest)
		return
	}
	if body.PaymentMethod == 0 {
		utils.SendJsonError(w, "payment_method cannot 0 or empty", http.StatusBadRequest)
		return
	}

	// validate paymebt method id
	s := repositories.GetPaymentMethodRepository()
	_, err = s.GetById(body.PaymentMethod)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.SendJsonError(w, "payment method not found", http.StatusBadRequest)
			return
		}
		log.Println("Error get payment method: ", err)
		utils.SendJsonError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// check have a cart
	c := repositories.GetCartRepository()
	if err := c.CreateCart(userId); err != nil {
		utils.SendJsonError(w, "internal server error", http.StatusInternalServerError)
		return
	}
	// check at least one product in cart
	productInCart, err := c.GetByUserId(userId)
	if err != nil {
		log.Println("error get cart by id ", err)
		utils.SendJsonError(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if len(productInCart) <= 0 {
		utils.SendJsonError(w, "you must have at least one product in your cart", http.StatusBadRequest)
		return
	}
	// insert into payment_transaction

	pt := repositories.GetPaymentTransactionRepository()
	if err := pt.Add(userId, body.PaymentMethod); err != nil {
		log.Println("Error insert to payment transaction: ", err)
		utils.SendJsonError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	utils.SendJsonSuccess(w, map[string]interface{}{"message": "succesfully checkout"}, http.StatusCreated)
}
