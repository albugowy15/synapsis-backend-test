package api

import (
	"github.com/albugowy15/synapsis-backend-test/internal/api/router"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/config"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/db"
)

func setConfiguration(path string) {
	config.LoadConfig(path)
	db.SetupDB()
}

func Run(path string) {
	setConfiguration(path)
	conf := config.GetConfig()

	router.Setup()
	router.Run(conf.Port)
}
