package main

import (
	"database/sql"
	"log"

	"github.com/albugowy15/synapsis-backend-test/database/connection"
	"github.com/albugowy15/synapsis-backend-test/database/seeder"
)

func main() {
	db := connection.NewConnection()
	defer db.Close()

	seedingFunctions := []func(*sql.DB) error{
		seeder.SeedUser,
		seeder.SeedPaymentMethods,
		seeder.SeedProducts,
		seeder.SeedCategories,
		seeder.SeedProductsCategories,
	}

	// Execute each seeding function
	for _, seedFunc := range seedingFunctions {
		if err := seedFunc(db); err != nil {
			log.Fatal(err)
		}
	}
}
