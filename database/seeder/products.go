package seeder

import (
	"database/sql"
	"log"
	"math/rand"
)

type ProductSeed struct {
	Name        string
	Description string
}

func randomStock() int {
	stock := rand.Intn(201)
	return stock
}

func randomPrice() float64 {
	price := rand.Int31n(50000)
	return float64(price)
}

var products = []ProductSeed{
	{
		Name:        "Kacamata Hitam Stainless Steel",
		Description: "Kacamata hitam premium berbahan stainles steel",
	},
	{
		Name:        "Masker Kf94 4Ply 4 ply korea",
		Description: "Masker Kf94 4Ply 4 ply korea Evo Plusmed Convex 4D Hitam Putih 10pcs - Hitam",
	},
	{
		Name:        "Ellipsesinc - Kaos Oversize Pria Wanita NY - Putih",
		Description: "Ellipsesinc - Kaos Oversize Pria Wanita NY - Putih",
	},
	{
		Name:        "PINZY Headset Bluetooth",
		Description: "PINZY Headset Bluetooth Sport Magnetic design Original Pz03",
	},
	{
		Name:        "SATOO FANES Headset",
		Description: "SATOO FANES Headset Earphone Handsfree High Bass With Microphone",
	},
	{
		Name:        "Kartu Pokemon",
		Description: "Kartu Pokemon Gendut Pelangi Bahan Metal Emas Silver 10pcs Inggris - 10pcs Blackmatt",
	},
	{
		Name:        "Stevor | Beras Pin pin 5 Kg Murah",
		Description: "Beras Pin Pin 5 Kg - Pilihan Utama untuk Gizi Keluarga ",
	},
	{
		Name:        "Rubber paint carlas",
		Description: "Rubber paint carlas 400ml - C190 GloCoating",
	},
	{
		Name:        "Coach League Hybrid Crossbody Men Sling Messenger Bag and Pouch Navy",
		Description: "Coach League Hybrid Crossbody Men Sling Messenger Bag and Pouch Navy",
	},
	{
		Name:        "Mr Mads",
		Description: "Harold Tas Tablet 12 Inch - Tas Selempang Sling Bag Anti Air",
	},
}

func SeedProducts(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO products (name, description, price, stock) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, product := range products {
		_, err := stmt.Exec(product.Name, product.Description, randomPrice(), randomStock())
		if err != nil {
			return err
		}
	}

	log.Println("Seed products successfull")
	return nil
}
