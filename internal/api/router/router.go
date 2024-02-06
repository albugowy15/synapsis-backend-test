package router

import (
	"github.com/albugowy15/synapsis-backend-test/internal/api/controllers"

	_ "github.com/albugowy15/synapsis-backend-test/docs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
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
func Setup() *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5, "text/html", "application/json"))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	r.Get("/v1/products", controllers.Products)

	r.Post("/v1/auth/register", controllers.Register)
	r.Post("/v1/auth/login", controllers.Login)

	r.Get("/v1/carts", controllers.Carts)
	r.Post("/v1/cart", controllers.Cart)
	r.Post("/v1/carts/checkout", controllers.Checkout)

	return r
}
