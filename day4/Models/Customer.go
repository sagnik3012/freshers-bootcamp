package Models

import (
	"freshers-bootcamp/day4/Config"
	_ "github.com/go-sql-driver/mysql"
)

func AddCustomer(customer *Customer) (err error) {
	if err = Config.DB.Model(&Customer{}).Create(customer).Error; err != nil {
		return err
	}
	return nil
}

// helper functions for  CUSTOMER table
func GetCustomerByID(custom *Customer, id string) (err error) {
	if err = Config.DB.Model(&Customer{}).Where("id = ?", id).First(custom).Error; err != nil {
		return err
	}
	return nil
}
