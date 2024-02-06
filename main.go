package main

import (
	"github.com/albugowy15/synapsis-backend-test/internal/api"
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
	api.Run(".")
}
