package api

import (
	"log"
	"net/http"

	"github.com/albugowy15/synapsis-backend-test/internal/api/router"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/config"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/db"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/utils"
)

func setConfiguration(path string) {
	config.LoadConfig(path)
	db.SetupDB()
}

func Run(path string) {
	setConfiguration(path)
	conf := config.GetConfig()
	utils.SetupAuth(conf.Secret)

	log.Printf("Server running on port %s", conf.Port)
	web := router.Setup()
	http.ListenAndServe(":"+conf.Port, web)
}
