package models

type Item struct {
	Product  Product `db:"product" json:"product"`
	ID       int64   `db:"id" json:"id"`
	Quantity int     `db:"quantity" json:"quantit"`
}

type Cart struct {
	ProductName  string  `db:"product_name" json:"product_name"`
	ProductPrice float64 `db:"product_price" json:"product_price"`
	ProductId    int64   `db:"product_id" json:"product_id"`
	ItemPrice    float64 `db:"item_price" json:"item_price"`
	Quantity     int64   `db:"quantity" json:"quantity"`
}

type AddCartRequest struct {
	ProductId int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}
