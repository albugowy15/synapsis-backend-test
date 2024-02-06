package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/albugowy15/synapsis-backend-test/internal/api/controllers"
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
func Setup() {
	http.HandleFunc("/v1/auth/register", controllers.Register)
	http.HandleFunc("/v1/auth/login", controllers.Login)

	http.HandleFunc("/v1/products", controllers.Products)

	http.HandleFunc("/v1/carts", controllers.Carts)
	http.HandleFunc("/v1/cart", controllers.Cart)
	http.HandleFunc("/v1/carts/checkout", controllers.Checkout)
}

func Run(port string) {
	log.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
