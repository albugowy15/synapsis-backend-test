package repositories

import (
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/db"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
)

type ProductRepository struct{}

var productRepository *ProductRepository

func GetProductRepository() *ProductRepository {
	if productRepository == nil {
		productRepository = &ProductRepository{}
	}
	return productRepository
}

func (r *ProductRepository) GetByCategory(category string) ([]models.Product, error) {
	var products []models.Product
	rows, err := db.
		GetDB().
		Query(
			`SELECT p.id, p.name, p.description, p.price, p.stock 
        FROM products p 
        INNER JOIN products_categories pc ON p.id = pc.product_id 
        INNER JOIN categories c ON c.id = pc.category_id 
        WHERE c.name = $1`,
			category,
		)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock); err != nil {
			return products, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return products, err
	}
	return products, nil
}

func (r *ProductRepository) All() ([]models.Product, error) {
	var products []models.Product
	rows, err := db.
		GetDB().
		Query(
			"SELECT id, name, description, price, stock FROM products",
		)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock); err != nil {
			return products, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return products, err
	}
	return products, nil
}

func (r *ProductRepository) GetById(id int64) (models.Product, error) {
	var product models.Product
	if err := db.
		GetDB().
		QueryRow(
			"SELECT id, name, description, price, stock FROM products WHERE id = $1",
			id,
		).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock); err != nil {
		return product, err
	}
	return product, nil
}
