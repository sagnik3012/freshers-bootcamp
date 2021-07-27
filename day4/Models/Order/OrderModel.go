package Models



type Order struct {
	ID         int64    `json:"order_id"`
	Customer   Customer `gorm:"foreign_key:CustomerId"`
	CustomerId int64    `json:"customer_id" `
	ProductId  int64    `json:"product_id"`
	Product    Product  `json:"product" gorm:"foreign_key:ProductId"`
	Quantity   int64    `json:"quantity"`
	Status     string   `json:"status"`
}

func (b *Order) TableName() string {
	return "Orders"
}
package Models
type Customer struct {
	ID       int64  `json:"customer_id"`
	Name     string `json:"name"`
	Activity bool   `json:"status"`
	// add other fields
}

func (b *Customer) TableName() string {
	return "Customers"
}

type Product struct {
	ID          int64  `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
}

func (b *Product) TableName() string {
	return "Products"
}




