package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/repositories"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/utils"
	"github.com/go-chi/chi/v5"
)

// Get All User Transactions
// @Tags transactions v1
// @Summary Get all user transactions
// @Description Get all user transactions
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce json
// @Success 200 {array} models.ArrayResponse
// @Success 400 {object} models.ErrorResponse
// @Success 404 {object} models.ErrorResponse
// @Success 500 {object} models.ErrorResponse
// @Router /transactions [get]
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetJwtClaim(r)
	if err != nil {
		utils.SendJsonError(w, "error parsing token", http.StatusBadRequest)
		return
	}
	t := repositories.GetPaymentTransactionRepository()
	transactions, err := t.All(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.SendJsonError(w, "you dont have any transaction", http.StatusNotFound)
			return
		}
		log.Println(err)
		utils.SendJsonError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	response := models.ArrayResponse{
		TotalItems: len(transactions),
		Data:       transactions,
	}
	utils.SendJsonSuccess(w, response, http.StatusOK)
}

// Get Transaction Detail
// @Tags transactions v1
// @Summary Get transactions detail
// @Description Get transactions detail
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param transaction_id path int true "Transaction ID"
// @Produce json
// @Success 200 {array} models.PaymentTransactionDetailSwagger
// @Success 400 {object} models.ErrorResponse
// @Success 404 {object} models.ErrorResponse
// @Success 500 {object} models.ErrorResponse
// @Router /transactions/{transaction_id} [get]
func GetTransactionById(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetJwtClaim(r)
	if err != nil {
		utils.SendJsonError(w, "error parsing token", http.StatusBadRequest)
		return
	}

	transactionIdStr := chi.URLParam(r, "transaction_id")
	transactionId, err := strconv.ParseInt(transactionIdStr, 10, 64)
	if err != nil {
		utils.SendJsonError(w, "Transaction id must be integer", http.StatusBadRequest)
		return
	}
	if transactionId == 0 {
		utils.SendJsonError(w, "Transaction cannot be 0", http.StatusBadRequest)
		return
	}

	t := repositories.GetPaymentTransactionRepository()
	transaction, err := t.GetById(transactionId, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.SendJsonError(w, "transaction not found", http.StatusBadRequest)
			return
		}
		log.Println("error get transaction by id: ", err)
		utils.SendJsonError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	utils.SendJsonSuccess(w, transaction, http.StatusOK)
}
