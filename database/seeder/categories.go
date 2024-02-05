package seeder

import (
	"database/sql"
	"log"
)

type CateogorySeed struct {
	Name        string
	Description string
}

var categories = []CateogorySeed{
	{
		Name:        "fashion",
		Description: "Koleksi barang fashion terbaik",
	},
	{
		Name:        "kesehatan",
		Description: "Koleksi barang dan alat kesehatan penunjang sehari-hari",
	},
	{
		Name:        "elektronik",
		Description: "Koleksi elektronik murah kualitas oke",
	},
	{
		Name:        "mainan",
		Description: "Koleksi mainan untuk dimainkan",
	},
	{
		Name:        "makanan",
		Description: "Koleksi bahan makanan mentah dan siap saji",
	},
	{
		Name:        "material",
		Description: "Koleksi material untuk bangunan dan ruang",
	},
}

func SeedCategories(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO categories (name, description) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, category := range categories {
		_, err := stmt.Exec(category.Name, category.Description)
		if err != nil {
			return err
		}
	}

	log.Println("Seed categories successfull")
	return nil
}
