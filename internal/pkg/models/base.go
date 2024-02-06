package models

type ArrayResponse struct {
	Data       any `json:"data"`
	TotalItems int `json:"total_items"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
