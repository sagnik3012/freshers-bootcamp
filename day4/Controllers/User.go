package Controllers

import (
	_ "encoding/json"
	"fmt"
	"freshers-bootcamp/day4/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Get all orders issued till current date ( for retailer )

func GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	//prodID := c.Params.ByName("product_id")
	var ord Models.Order
	err := Models.GetOrderByID(&ord, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"order_id":ord.ID , "product_id" : ord.ProductId,
			"quantity":ord.Quantity, "status":"order placed"})
	}
}

func GetAllOrders( c *gin.Context){
	var order []Models.Order
	err := Models.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var prod *Models.Product
		for i := range order{
			order[i].Customer.ID = order[i].CustomerId
		}
		c.JSON(http.StatusOK, order)
	}
}

// get all products in the retail shop ( for retailer)

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

// add a new product in the retail shop ( for retailer)

func CreateProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Models.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusConflict)
	} else {

		c.JSON(http.StatusOK,gin.H{"id":product.ID , "product_name": product.ProductName,
			"price":product.Price,"quantity":product.Quantity,
			"response":"product successfully added!"})
	}
}

// edit details of an existing product

func PatchProduct(c *gin.Context) {
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

// fetch product by its ID ( for customer)

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

func PlaceOrder (c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	if err := Models.PlaceOrder(&order); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
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

