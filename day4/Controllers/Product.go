package Controllers

import (
	_ "encoding/json"
	"fmt"
	"freshers-bootcamp/day4/Models"

	"github.com/gin-gonic/gin"
	"net/http"
)

// get all products in the retail shop ( for retailer)

func GetAllProducts(c *gin.Context) {
	var prod []Models.Product
	err := Models.GetAllProducts(&prod)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		resp := gin.H{"products": prod}
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

		c.JSON(http.StatusOK, gin.H{"id": product.ID, "product_name": product.ProductName,
			"price": product.Price, "quantity": product.Quantity,
			"response": "product successfully added!"})
	}
}

// edit details of a product

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

// fetch product by its ID
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

