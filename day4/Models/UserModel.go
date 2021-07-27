package Models

type Product struct {
	Id          uint   `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
}
type Order struct {
	Id         uint   `json:"order_id"`
	CustomerId uint   `json:"customer_id"`
	ProductId  uint `json:"product_id"`
	Quantity   uint   `json:"quantity"`
}
type Customer struct {
	Id   uint   `json:"customer_id"`
	Name string `json:"name"`
}
