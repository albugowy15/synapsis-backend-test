package validator

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/repositories"
)

func ValidateUsername(username string) error {
	usernameLength := len(username)
	if usernameLength < 8 || usernameLength > 16 {
		return fmt.Errorf("username must be between 8 and 16 character")
	}

	pattern := "^[a-zA-Z0-9]+$"
	re := regexp.MustCompile(pattern)
	if !re.MatchString(username) {
		return fmt.Errorf("username must be only contains uppercase or lowercase letter and number with no space")
	}

	return nil
}

func ValidateUniqueUsername(username string) error {
	s := repositories.GetUserRepository()
	_, err := s.GetByUsername(username)
	if err == nil {
		return fmt.Errorf("username already exist")
	}
	return nil
}

func ValidateEmail(email string) error {
	emailLength := len(email)
	if emailLength > 100 {
		return fmt.Errorf("email must be less than 100 character")
	}

	pattern := "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	re := regexp.MustCompile(pattern)
	if !re.MatchString(email) {
		return fmt.Errorf("email is not valid")
	}

	s := repositories.GetUserRepository()
	_, err := s.GetByEmail(email)
	if err == nil {
		return fmt.Errorf("email already exist")
	}

	return nil
}

func ValidateFullname(fullname string) error {
	fullnameLength := len(fullname)
	if fullnameLength > 200 {
		return fmt.Errorf("fullname maximum 200 character length")
	}
	return nil
}

func ValidatePassword(password string) error {
	passwordLength := len(password)
	if passwordLength < 8 || passwordLength > 16 {
		return fmt.Errorf("password must be between 8 and 16 character")
	}
	if strings.Contains(password, " ") {
		return fmt.Errorf("password cannot contain space")
	}
	upperRegex := `[A-Z]`
	if !regexp.MustCompile(upperRegex).MatchString(password) {
		return fmt.Errorf("password must have uppercase letter")
	}
	lowerRegex := `[a-z]`
	if !regexp.MustCompile(lowerRegex).MatchString(password) {
		return fmt.Errorf("password must have lowercase letter")
	}
	digitRegex := `[0-9]`
	if !regexp.MustCompile(digitRegex).MatchString(password) {
		return fmt.Errorf("password must have number")
	}
	specialRegex := `[^\w\s]`
	if !regexp.MustCompile(specialRegex).MatchString(password) {
		return fmt.Errorf("password must special character")
	}
	return nil
}

func ValidateRegisterBody(body models.UserRegisterRequest) error {
	if err := ValidateUsername(body.Username); err != nil {
		return err
	}

	if err := ValidateUniqueUsername(body.Username); err != nil {
		return err
	}

	if err := ValidateEmail(body.Email); err != nil {
		return err
	}

	if err := ValidateFullname(body.Fullname); err != nil {
		return err
	}

	if err := ValidatePassword(body.Password); err != nil {
		return err
	}
	return nil
}

func ValidateLoginRequest(body models.UserLoginRequest) error {
	if err := ValidateUsername(body.Username); err != nil {
		return err
	}

	if err := ValidatePassword(body.Password); err != nil {
		return err
	}
	return nil
}
