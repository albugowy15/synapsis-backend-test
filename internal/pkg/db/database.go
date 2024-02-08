package db

import (
	"database/sql"
	"log"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/config"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

type Database struct {
	*sql.DB
}

func SetupDB() {
	configuration := config.GetConfig()

	db, err := sql.Open(configuration.DBDriver, configuration.DBSource)
	if err != nil {
		log.Printf("Error create db connection: %v", err)
	}
	DB = db
}

func GetDB() *sql.DB {
	return DB
}
