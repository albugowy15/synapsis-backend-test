package controllers

import (
	"net/http"
)

func Carts(w http.ResponseWriter, r *http.Request) {
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

func Cart(w http.ResponseWriter, r *http.Request) {
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
