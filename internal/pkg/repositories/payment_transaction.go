package repositories

import (
	"fmt"
	"log"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/db"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
)

type PaymentTransactionRepository struct{}

var paymentTransactionRepository *PaymentTransactionRepository

func GetPaymentTransactionRepository() *PaymentTransactionRepository {
	if paymentTransactionRepository == nil {
		paymentTransactionRepository = &PaymentTransactionRepository{}
	}
	return paymentTransactionRepository
}

func calculateTax(totalPrice float64) float64 {
	return 5.00 / 100.00 * totalPrice
}

func calculateTotalTransactionPrice(paymentFee float64, checkoutPrice float64, tax float64) float64 {
	return checkoutPrice + tax + paymentFee
}

var (
	VANumber      = "87562552325532"
	QRCode        = "some qr string"
	AccountNumber = "5800625266628828343"
)

func (r *PaymentTransactionRepository) Add(userId int64, paymentMethodId int) error {
	pm := GetPaymentMethodRepository()
	paymentMethod, err := pm.GetById(paymentMethodId)
	if err != nil {
		return err
	}
	// calculate total product price
	productsInCart, err := GetCartRepository().GetByUserId(userId)
	if err != nil {
		return err
	}
	totalPriceInCart := 0.0
	for _, item := range productsInCart {
		totalPriceInCart += item.ItemPrice
	}
	tax := calculateTax(totalPriceInCart)
	totalProductPrice := calculateTotalTransactionPrice(paymentMethod.Fee, totalPriceInCart, tax)
	// create qr code if payment using qris
	// create virtual account if payment useing VA
	// have account number
	//
	tx, err := db.GetDB().Begin()
	if err != nil {
		log.Fatal("Error starting transaction", err)
	}

	var qrCode *string = nil
	if paymentMethod.Type == "QRIS" {
		qrCode = &QRCode
	}
	var vaNumber *string = nil
	if paymentMethod.Type == "VA_TRANSFER" {
		vaNumber = &VANumber
	}
	var accountNumber *string = nil
	if paymentMethod.Type == "MANUAL_TRANSFER" {
		accountNumber = &AccountNumber
	}

	_, err = tx.Exec(`
    INSERT INTO payment_transactions (
      total_product_price, 
      tax, 
      total_price, 
      payment_method_id, 
      status, 
      qr_code, 
      virtual_account, 
      account_number, 
      user_id) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9
    )`, totalProductPrice, tax, totalPriceInCart, paymentMethod.ID, "WAIT", qrCode, vaNumber, accountNumber, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range productsInCart {
		_, err := tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (r *PaymentTransactionRepository) All(userId int64) ([]models.PaymentTransaction, error) {
	var transactions []models.PaymentTransaction
	rows, err := db.GetDB().
		Query(`
    SELECT pt.id, pt.total_product_price, pt.tax, pt.total_price, pt.payment_method_id, pt.status, pt.qr_code, pt.virtual_account,
      pt.account_number, pt.user_id 
      FROM payment_transactions pt 
      WHERE pt.user_id = $1
    `, userId)
	if err != nil {
		return transactions, err
	}

	log.Println(transactions)
	for rows.Next() {
		var transaction models.PaymentTransaction
		err := rows.Scan(&transaction.ID,
			&transaction.TotalProductPrice,
			&transaction.Tax,
			&transaction.TotalPrice,
			&transaction.PaymentMethodID,
			&transaction.Status,
			&transaction.QRCode,
			&transaction.VirtualAccount,
			&transaction.AccountNumber,
			&transaction.UserId)
		if err != nil {
			return transactions, fmt.Errorf("error scanning transaction row: %v", err)
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *PaymentTransactionRepository) GetById(transactionId int64, userId int64) (models.PaymentTransactionDetail, error) {
	var transaction models.PaymentTransactionDetail
	err := db.GetDB().QueryRow(`
    SELECT pt.id, pt.total_product_price, pt.tax, pt.total_price, pt.status, pt.qr_code, pt.virtual_account,
      pt.account_number, pm.id, pm.type, pm.merchant, pm.fee
      FROM payment_transactions pt 
      INNER JOIN payment_methods pm ON pm.id = pt.payment_method_id
      WHERE pt.id = $1 AND pt.user_id = $2
    `, transactionId, userId).
		Scan(
			&transaction.ID,
			&transaction.TotalProductPrice,
			&transaction.Tax,
			&transaction.TotalPrice,
			&transaction.Status,
			&transaction.QRCode,
			&transaction.VirtualAccount,
			&transaction.AccountNumber,
			&transaction.PaymentMethod.ID,
			&transaction.PaymentMethod.Type,
			&transaction.PaymentMethod.Merchant,
			&transaction.PaymentMethod.Fee,
		)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
