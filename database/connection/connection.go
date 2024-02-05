package connection

import (
	"database/sql"
	"log"

	"github.com/albugowy15/synapsis-backend-test/utils"
	_ "github.com/lib/pq"
)

func NewConnection() *sql.DB {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("Error create db connection: %v", err)
	}
	return db
}
