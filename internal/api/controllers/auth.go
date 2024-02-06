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
	"golang.org/x/crypto/bcrypt"
)

// Register a user
// @Tags auth v1
// @Summary Register user
// @Description Register a new user
// @Accept json
// @Param q body models.UserRegisterRequest
// @Produce json
// @Success 201 {object} models.UserRegisterResponse
// @Router /auth/register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	var body models.UserRegisterRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("Error parse body: %v", err)
		utils.SendJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.ValidateRegisterBody(body); err != nil {
		utils.SendJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashPassword, err := utils.HashString(body.Password)
	if err != nil {
		utils.SendJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	insertUserData := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: hashPassword,
		Fullname: body.Fullname,
	}
	s := repositories.GetUserRepository()

	if err := s.Create(insertUserData); err != nil {
		log.Printf("Failed create user: %v", err)
		utils.SendJsonError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	utils.SendJsonSuccess(w, models.UserRegisterResponse{
		Message: "Account registerd",
	}, http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var body models.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("Error parse body: %v", err)
		utils.SendJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.ValidateLoginRequest(body); err != nil {
		utils.SendJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	s := repositories.GetUserRepository()
	res, err := s.GetByUsername(body.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.SendJsonError(w, "Username not found", http.StatusNotFound)
			return
		}
		utils.SendJsonError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(body.Password)); err != nil {
		utils.SendJsonError(w, "Password not match", http.StatusBadRequest)
		return
	}

	jwtClaims := map[string]interface{}{
		"user_id":  res.ID,
		"username": res.Username,
		"email":    res.Email,
	}
	_, token, err := utils.GetAuth().Encode(jwtClaims)
	if err != nil {
		log.Printf("Error signing token: %v", err)
	}
	utils.SendJsonSuccess(w, map[string]interface{}{"token": token}, http.StatusOK)
}
