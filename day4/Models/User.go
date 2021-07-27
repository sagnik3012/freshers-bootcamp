package Models

import (
	"fmt"
	"freshers-bootcamp/day4/Config"
	_ "github.com/go-sql-driver/mysql"
)

// POST / product
func CreateProduct(prod *Product) (err error) {
	if err = Config.DB.Model(&Product{}).Create(prod).Error; err != nil {
		return err
	}
	return nil
}

// PATCH / product / id
func UpdateProduct(prod *Product, id string) (err error) {
	fmt.Println(prod)
	Config.DB.Save(prod)
	return nil
}

// GET / product / :id
func GetProductByID(prod *Product, id string) (err error) {
	if err = Config.DB.Model(&Product{}).Where("id = ?", id).First(prod).Error; err != nil {
		return err
	}
	return nil
}

// GET / products
func GetAllProducts(prods *[]Product) (err error) {
	if err = Config.DB.Model(&Product{}).Find(prods).Error; err != nil {
		return err
	}
	return nil
}

// POST / order
func PlaceOrder(ord *Order) (err error) {

	if err = Config.DB.Model(&Order{}).Create(ord).Error; err != nil {
		return err
	}
	return nil
}

// GET / order / :id
func GetOrderByID(ord *Order, id string) (err error) {
	if err = Config.DB.Model(&Order{}).Where("id = ?", id).First(ord).Error; err != nil {
		return err
	}
	return nil

}
func GetAllOrders( ord *[]Order )( err error){
	if err = Config.DB.Model(&Order{}).Find(ord).Error ; err != nil {
		return err

	}
	return nil

}
func DeleteAllProducts() (err error) {
	Config.DB.Where("1=?", 1).Delete(&Product{})
	return nil
}

/*
func GetAllOrders(ord *[]Order) (err error) {
	if err = Config.DB.Find(ord).Error; err != nil {
		return err
	}
	return nil
}
*/
