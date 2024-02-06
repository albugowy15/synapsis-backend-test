package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/repositories"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/utils"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/validator"
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
	/* Step register
		 * 1. Dapatkan username, email, fullname, dan password dari body
	  * 2. Validasi username unique, minimal, maksimal, no spasi, hanya huruf besar, kecil, dan angka
	* 3. Validasi email valid, unique, maksimal karakter
	* 4. Validasi fullname maksimal karakter
	* 5. Validasi password tidak ada whitespace, terdiri dari huruf, symbol, dan angka
	* 6. Hash password
	* 7. Insert ke db
	* 8. Kembalikan response
	*/
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

	hashPassword, err := utils.HashStr(body.Password)
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
}
