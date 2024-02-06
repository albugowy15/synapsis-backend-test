package utils

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

func SetupAuth(secret string) {
	tokenAuth = jwtauth.New("HS256", []byte(secret), nil)
}

func GetAuth() *jwtauth.JWTAuth {
	return tokenAuth
}

func GetJwtClaim(r *http.Request) (int64, error) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userId, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid token")
	}
	return int64(userId), nil
}
