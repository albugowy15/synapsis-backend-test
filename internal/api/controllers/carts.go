package controllers

import (
	"log"
	"net/http"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/repositories"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/utils"
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
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

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
