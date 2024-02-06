package db

import (
	"database/sql"
	"log"
	"time"

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
		log.Fatalf("Error create db connection: %v", err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(time.Duration(100) * time.Second)
	DB = db
}

func GetDB() *sql.DB {
	return DB
}
