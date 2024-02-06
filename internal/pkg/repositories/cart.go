package repositories

import (
	"database/sql"
	"log"

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

func (r *CartRepository) CheckProductExistInCart(productId int64, userId int64) bool {
	var ID int64
	err := db.
		GetDB().
		QueryRow(
			`SELECT p.id 
			FROM products p 
			INNER JOIN items i ON i.product_id = p.id 
			INNER JOIN shopping_carts sc ON sc.id = i.shopping_cart_id
			WHERE p.id = $1 AND sc.user_id = $2`, productId, userId).
		Scan(&ID)
	if err == nil {
		return true
	}
	return err != sql.ErrNoRows
}

func (r *CartRepository) CheckStockAvailable(productId int64, quantity int64) (bool, error) {
	var stock int64
	err := db.
		GetDB().
		QueryRow(
			`SELECT stock FROM products WHERE id = $1`,
			productId,
		).Scan(&stock)
	if err != nil {
		return false, err
	}
	return stock > quantity, nil
}

func (r *CartRepository) Add(productId int64, quantity int64, userId int64) error {
	tx, err := db.GetDB().Begin()
	if err != nil {
		log.Println("errro creating tx: ", err)
		return err
	}
	var cartId int64
	err = tx.QueryRow(`SELECT id as cartId from shopping_carts WHERE user_id = $1`, userId).Scan(&cartId)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err := tx.Exec(`INSERT INTO shopping_carts (user_id) VALUES ($1)`, userId)
			if err != nil {
				log.Println("error inserting to shopping_cart: ", err)
				tx.Rollback()
				return err
			}

			err = tx.QueryRow(`SELECT id as cartId from shopping_carts WHERE user_id = $1`, userId).Scan(&cartId)
			if err != nil {
				log.Println("error get cartid ", err)
				tx.Rollback()
				return err
			}
		} else {
			log.Print("hello")
			tx.Rollback()
			return err
		}
	}

	_, err = tx.Exec(`INSERT INTO items (quantity, shopping_cart_id, product_id) VALUES ($1, $2, $3)`, quantity, cartId, productId)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
