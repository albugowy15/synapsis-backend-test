package repositories

import (
	"fmt"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/db"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
)

type UserRepository struct{}

var userRepository *UserRepository

func GetUserRepository() *UserRepository {
	if userRepository == nil {
		userRepository = &UserRepository{}
	}
	return userRepository
}

func (r *UserRepository) GetById(id string) (models.User, error) {
	var user models.User
	if err := db.
		GetDB().
		QueryRow(
			"SELECT id, username, email, fullname, password, address, phone_number FROM users WHERE id = $1",
			id,
		).Scan(&user.ID, &user.Username, &user.Email, &user.Fullname, &user.Password, &user.Address, &user.PhoneNumber); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) GetByUsername(username string) (models.User, error) {
	var user models.User
	if err := db.
		GetDB().
		QueryRow(
			"SELECT id, username, email, fullname, password, address, phone_number FROM users WHERE username = $1",
			username,
		).Scan(&user.ID, &user.Username, &user.Email, &user.Fullname, &user.Password, &user.Address, &user.PhoneNumber); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	if err := db.
		GetDB().
		QueryRow(
			"SELECT id, username, email, fullname, password, address, phone_number FROM users WHERE email = $1",
			email,
		).Scan(&user.ID, &user.Username, &user.Email, &user.Fullname, &user.Password, &user.Address, &user.PhoneNumber); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) Create(user models.User) error {
	_, err := db.GetDB().Exec("INSERT INTO users (username, email, fullname, password, address, phone_number) VALUES ($1, $2, $3, $4, $5, $6)",
		user.Username, user.Email, user.Fullname, user.Password, nil, nil,
	)
	if err != nil {
		return fmt.Errorf("error inserting user with username %s: %v", user.Username, err)
	}
	return nil
}
