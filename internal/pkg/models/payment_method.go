package models

type PaymentMethod struct {
	Type     string  `db:"type" json:"type"`
	Merchant string  `db:"merchant" json:"metchant"`
	ID       int     `db:"id" json:"id"`
	Fee      float64 `db:"fee" json:"fee"`
}
