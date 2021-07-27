package Models

type Product struct {
	ID          int64  `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
}

func (b *Product) TableName() string {
	return "Products"
}


