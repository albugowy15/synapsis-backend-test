package utils

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth"
)

func GetJwtClaim(r *http.Request) (int64, error) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userId, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid token")
	}
	return int64(userId), nil
}
