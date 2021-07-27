package Controllers

import (
	_ "encoding/json"
	"fmt"

	"freshers-bootcamp/day4/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)
func GetAllOrders( c *gin.Context){
	var order []Models.Order
	err := Models.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func GetAllProducts(c *gin.Context) {
	var prod []Models.Product
	err := Models.GetAllProducts(&prod)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		resp := gin.H{ "products" : prod}
		c.JSON(http.StatusOK, resp)
	}
}

func CreateProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Models.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusConflict)
	} else {

		c.JSON(http.StatusOK,gin.H{"id":product.Id , "product_name": product.ProductName,
			"price":product.Price,"quantity":product.Quantity,
			"response":"product successfully added!"})

	}
}
func UpdateProduct(c *gin.Context) {
	var prod Models.Product
	id := c.Params.ByName("id")

	err := Models.GetProductByID(&prod, id)
	if err != nil {
		c.JSON(http.StatusNotFound, prod)
	}
	c.BindJSON(&prod)
	err = Models.UpdateProduct(&prod, id)
	if err != nil {
		c.JSON(http.StatusNotFound, prod)
	} else {
		c.JSON(http.StatusOK, prod)
	}
}

func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var prod Models.Product
	err := Models.GetProductByID(&prod, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, prod)
	}
}

func PlaceOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	err := Models.PlaceOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else{

		c.JSON(http.StatusOK, order)
	}
}
func GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Order
	err := Models.GetOrderByID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func DeleteAllProducts(c *gin.Context) {
	err := Models.DeleteAllProducts()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, "all product details deleted!")
	}

}

/*

 */