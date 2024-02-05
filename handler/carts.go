package handler

import (
	"fmt"
	"net/http"
)

func Carts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from carts")
}

func Cart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from cart")
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Checkout")
}
