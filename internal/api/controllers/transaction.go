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
