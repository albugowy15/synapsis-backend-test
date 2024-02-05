package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewConnection() *sql.DB {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "postgres"
	dbname := "synapsis_db"

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error create db connection: %v", err)
	}
	return db
}
