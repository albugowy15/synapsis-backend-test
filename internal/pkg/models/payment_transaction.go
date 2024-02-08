package models

import "database/sql"

type PaymentTransaction struct {
	Status            string         `db:"status" json:"status"`
	QRCode            sql.NullString `db:"qr_code" json:"qr_code"`
	VirtualAccount    sql.NullString `db:"virtual_account" json:"virtual_account"`
	AccountNumber     sql.NullString `db:"account_number" json:"account_number"`
	ID                int64          `db:"id" json:"id"`
	TotalProductPrice float64        `db:"total_product_price" json:"total_product_price"`
	Tax               float64        `db:"tax" json:"tax"`
	TotalPrice        float64        `db:"total_price" json:"total_price"`
	PaymentMethodID   int            `db:"payment_method_id" json:"payment_method_id"`
	UserId            int64          `db:"user_id" json:"user_id"`
}

type PaymentTransactionDetail struct {
	Status            string         `db:"status" json:"status"`
	QRCode            sql.NullString `db:"qr_code" json:"qr_code"`
	VirtualAccount    sql.NullString `db:"virtual_account" json:"virtual_account"`
	AccountNumber     sql.NullString `db:"account_number" json:"account_number"`
	PaymentMethod     PaymentMethod  `db:"payment_method" json:"payment_method"`
	ID                int64          `db:"id" json:"id"`
	TotalProductPrice float64        `db:"total_product_price" json:"total_product_price"`
	Tax               float64        `db:"tax" json:"tax"`
	TotalPrice        float64        `db:"total_price" json:"total_price"`
}

type PaymentTransactionDetailSwagger struct {
	Status            string        `db:"status" json:"status"`
	QRCode            string        `db:"qr_code" json:"qr_code"`
	VirtualAccount    string        `db:"virtual_account" json:"virtual_account"`
	AccountNumber     string        `db:"account_number" json:"account_number"`
	PaymentMethod     PaymentMethod `db:"payment_method" json:"payment_method"`
	ID                int64         `db:"id" json:"id"`
	TotalProductPrice float64       `db:"total_product_price" json:"total_product_price"`
	Tax               float64       `db:"tax" json:"tax"`
	TotalPrice        float64       `db:"total_price" json:"total_price"`
}
