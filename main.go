package main

import (
	"github.com/albugowy15/synapsis-backend-test/docs"
	"github.com/albugowy15/synapsis-backend-test/internal/api"
)

// @title Synapsis API Documentation
// @version 1.0
// @description This is a swagger documentation for Synapsis API Backend Test.

// @contact.name Mohamad Kholid Bughowi
// @contact.url https://bughowi.com
// @contact.email kholidbughowi@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
func main() {
	docs.SwaggerInfo.Host = "synapsis-backend-test.fly.dev"
	api.Run(".")
}
