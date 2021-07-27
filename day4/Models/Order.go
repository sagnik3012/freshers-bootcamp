package Models

import (
	"errors"
	"fmt"
	"freshers-bootcamp/day4/Config"
)

func PlaceOrder(ord *Order) (err error) {
	var product Product
	Config.DB.Model(&Product{}).Where("id = ?", ord.ProductId).First(&product)
	if ord.Quantity > product.Quantity {
		fmt.Println("cancelled!!!!")
		ord.Status = "order cancelled!!"
		Config.DB.Model(&Order{}).Create(ord)
		return errors.New("not sufficient products in store ")
	} else {
		fmt.Println("placed!!!!!")
		ord.Status = "order placed!!"
		Config.DB.Model(&Product{}).Where("id = ?", ord.ProductId).Update("quantity", product.Quantity-ord.Quantity)
		Config.DB.Model(&Order{}).Create(ord)
		return nil
	}
}

// GET / order / :id
func GetOrderByID(ord *Order, id string) (err error) {
	if err = Config.DB.Model(&Order{}).Where("id = ?", id).First(ord).Error; err != nil {
		return err
	}
	return nil

}

func GetAllOrders(ord *[]Order) (err error) {
	if err = Config.DB.Model(&Order{}).Find(ord).Error; err != nil {
		return err
	}
	return nil
}

// change the post order function
