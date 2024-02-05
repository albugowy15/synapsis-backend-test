package main

import (
	"log"
	"net/http"

	"github.com/albugowy15/synapsis-backend-test/handler"
)

/*
* Avaliable endpoints
* POST /v1/auth/register
* GET /v1/auth/login
* GET /v1/products?category
* POST /v1/shopping_cart
* GET /v1/shopping_carts
* DELETE /v1/shopping_cart
* POST /v1/shopping_carts/checkout
 */

func main() {
	http.HandleFunc("/v1/auth/register", handler.Register)
	http.HandleFunc("/v1/auth/login", handler.Login)

	http.HandleFunc("/v1/products", handler.Products)

	http.HandleFunc("/v1/carts", handler.Carts)
	http.HandleFunc("/v1/cart", handler.Cart)
	http.HandleFunc("/v1/carts/checkout", handler.Checkout)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
