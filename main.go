package main

import (
	_ "github.com/albugowy15/synapsis-backend-test/docs"
	"github.com/albugowy15/synapsis-backend-test/internal/api"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
func main() {
	api.Run(".")
}
