package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       uint32 `json:"price"`
	Stock       int16  `json:"stock"`
}

var products = []Product{
	{
		Name:        "Sepatu Vintage",
		Price:       200000,
		Description: "Sepatu vintage premium edisi terbaru 2024",
		Stock:       50,
		Category:    "shoes",
	},
	{
		Name:        "Charger Laptop Lenovo",
		Price:       50000,
		Description: "Charger ori laptop lenovo v14 ada",
		Stock:       10,
		Category:    "gadget",
	},
	{
		Name:        "Gatsby parfum baju",
		Price:       24000,
		Description: "Gatsby urban cologne long lasting parfume",
		Stock:       45,
		Category:    "clothes",
	},
	{
		Name:        "Baseus Bowie WM01",
		Price:       300000,
		Description: "Earbud baseus bowie keluaran terbaru seri WM01",
		Stock:       18,
		Category:    "accesories",
	},
	{
		Name:        "Xiaomi Redmi 9",
		Price:       1900000,
		Description: "HP Xiaomi seri Redmi edisi 9 dengan RAM 4 dan storage 64",
		Stock:       67,
		Category:    "gadget",
	},
}

func Products(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		query := r.URL.Query()
		categoryParam := query.Get("category")
		var productsByCategory []Product

		for _, product := range products {
			if product.Category == categoryParam {
				productsByCategory = append(productsByCategory, product)
			}
		}

		jsonResponse, err := json.Marshal(productsByCategory)
		if err != nil {
			log.Printf("Error marshal response json: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)

	case http.MethodPost:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	case http.MethodPut:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	case http.MethodDelete:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
