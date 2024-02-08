package main

import (
	"log"
	"net/http"

	"github.com/albugowy15/synapsis-backend-test/docs"
	"github.com/albugowy15/synapsis-backend-test/internal/api/router"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/config"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/db"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/utils"
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
	config.LoadConfig(".")
	conf := config.GetConfig()

	db.SetupDB()
	utils.SetupAuth(conf.Secret)
	web := router.Setup()
	docs.SwaggerInfo.Host = conf.ApiUrl

	log.Printf("Server running on port %s", conf.Port)
	http.ListenAndServe(":"+conf.Port, web)
}
