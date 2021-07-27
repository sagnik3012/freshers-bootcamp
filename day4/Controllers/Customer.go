package Controllers

import (
	"fmt"
	"freshers-bootcamp/day4/Models/Customer"
	"github.com/gin-gonic/gin"
	"net/http"
)

//CreateUser ... Create User
func AddCustomer( c*gin.Context){
	var customer Models.Customer
	c.BindJSON(&customer)
	err := Models.AddCustomer(&customer)
	if err != nil {
		fmt.Println("error : ", err)
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK,gin.H{"customer_id":customer.ID, "customer_name":customer.Name})
	}
}


func GetCustomerByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var customer Models.Customer
	err := Models.GetCustomerByID(&customer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"customer_id" : customer.ID,
			"customer_name" : customer.Name})
	}
}
