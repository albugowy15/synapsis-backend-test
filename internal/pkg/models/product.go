package models

type Product struct {
	Name        string  `db:"name" json:"name"`
	Description string  `db:"description" json:"description"`
	ID          int64   `db:"id" json:"id"`
	Price       float64 `db:"price" json:"price"`
	Stock       int     `db:"stock" json:"stock"`
}
