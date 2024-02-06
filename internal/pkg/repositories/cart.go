package repositories

import (
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/db"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
)

type CartRepository struct{}

var cartRepository *CartRepository

func GetCartRepository() *CartRepository {
	if cartRepository == nil {
		cartRepository = &CartRepository{}
	}
	return cartRepository
}

func (r *CartRepository) GetByUserId(userId int64) ([]models.Cart, error) {
	var carts []models.Cart

	rows, err := db.GetDB().Query(`SELECT p.id as product_id, p.name as product_name,
      p.price as product_price, (p.price * i.quantity) as item_price, i.quantity
      FROM products p
      INNER JOIN items i ON p.id = i.product_id
      INNER JOIN shopping_carts sc ON sc.id = i.shopping_cart_id
      WHERE sc.user_id = $1`, userId)
	if err != nil {
		return carts, err
	}

	for rows.Next() {
		var cart models.Cart
		err := rows.Scan(&cart.ProductId, &cart.ProductName, &cart.ProductPrice, &cart.ItemPrice, &cart.Quantity)
		if err != nil {
			return carts, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}
