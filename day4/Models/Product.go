package Models

import (
	"fmt"
	"freshers-bootcamp/day4/Config"
)

// create a new product
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
