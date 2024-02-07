package connection

import (
	"database/sql"
	"log"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/config"
	_ "github.com/lib/pq"
)

func NewConnection() *sql.DB {
	config.LoadConfig(".")
	conf := config.GetConfig()
	db, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatalf("Error create db connection: %v", err)
	}
	return db
}
