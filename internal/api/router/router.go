package router

import (
	"github.com/albugowy15/synapsis-backend-test/internal/api/controllers"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/utils"

	_ "github.com/albugowy15/synapsis-backend-test/docs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Setup() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5, "text/html", "application/json"))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	r.Get("/v1/products", controllers.GetProducts)

	r.Post("/v1/auth/login", controllers.Login)
	r.Post("/v1/auth/register", controllers.Register)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(utils.GetAuth()))
		r.Use(jwtauth.Authenticator)
		r.Get("/v1/carts", controllers.GetCarts)
		r.Post("/v1/carts", controllers.AddCart)
		r.Delete("/v1/carts/{product_id}", controllers.DeleteCart)
		r.Post("/v1/carts/checkout", controllers.Checkout)

		r.Get("/v1/transactions", controllers.GetTransactions)
		r.Get("/v1/transactions/{transaction_id}", controllers.GetTransactionById)
	})
	return r
}
