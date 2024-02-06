package models

type ArrayResponse struct {
	Data       any `json:"data"`
	TotalItems int `json:"total_items"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type JwtClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Id       int64  `json:"id"`
}
