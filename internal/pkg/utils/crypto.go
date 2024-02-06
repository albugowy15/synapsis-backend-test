package utils

import (
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

var tokenAuth *jwtauth.JWTAuth

func HashString(source string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(source), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func SetupAuth(secret string) {
	tokenAuth = jwtauth.New("HS256", []byte(secret), nil)
}

func GetAuth() *jwtauth.JWTAuth {
	return tokenAuth
}
