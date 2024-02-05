package seeder

import (
	"database/sql"
	"log"
	"math/rand"
)

func SeedProductsCategories(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO products_categories (product_id, category_id) VALUES ($1, $2)")
	if err != nil {
		return err
	}

	for i := 0; i < 20; i++ {
		productId := rand.Int63n(10) + 1
		categoryId := rand.Int63n(5) + 1
		_, err := stmt.Exec(productId, categoryId)
		if err != nil {
			return err
		}
	}

	log.Println("Seed products categories successfull")
	return nil
}
