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

