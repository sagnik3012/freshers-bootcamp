package Models
import(
	"freshers-bootcamp/day4/Config"
)
func AddCustomer(customer *Customer) (err error) {
	if err = Config.DB.Model(&Customer{}).Create(customer).Error; err != nil {
		return err
	}
	return nil
}

// helper functions for handling CUSTOMER table
func GetCustomerByID(custom *Customer, id string) (err error) {
	if err = Config.DB.Model(&Customer{}).Where("id = ?", id).First(custom).Error; err != nil {
		return err
	}
	return nil
}
