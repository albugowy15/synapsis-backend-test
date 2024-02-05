package seeder

import (
	"database/sql"
	"log"
)

type PaymentMethodSeed struct {
	Type     string
	Merchant string
	Fee      float64
}

var paymentMethods = []PaymentMethodSeed{
	{
		Type:     "TRANSFER",
		Merchant: "BNI",
		Fee:      3000,
	},
	{
		Type:     "TRANSFER",
		Merchant: "Mandiri",
		Fee:      3000,
	},
	{
		Type:     "TRANSFER",
		Merchant: "BRI",
		Fee:      3000,
	},
	{
		Type:     "EWALLET",
		Merchant: "Gopay",
		Fee:      2000,
	},
	{
		Type:     "EWALLET",
		Merchant: "Dana",
		Fee:      2000,
	},
	{
		Type:     "EWALLET",
		Merchant: "Ovo",
		Fee:      2000,
	},
	{
		Type:     "QRIS",
		Merchant: "qris",
		Fee:      500,
	},
}

func SeedPaymentMethods(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO payment_methods (type, merchant, fee) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, method := range paymentMethods {
		_, err := stmt.Exec(method.Type, method.Merchant, method.Fee)
		if err != nil {
			return err
		}
	}

	log.Println("Seed payment methods successfull")
	return nil
}
