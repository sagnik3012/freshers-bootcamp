package Controllers

import (
	"fmt"
	"freshers-bootcamp/day4/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// get orders by id
func GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	//prodID := c.Params.ByName("product_id")
	var ord Models.Order
	err := Models.GetOrderByID(&ord, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"order_id": ord.ID, "product_id": ord.ProductId,
			"quantity": ord.Quantity, "status": "order placed"})
	}
}

//

func PlaceOrder(c *gin.Context) {
	var order Models.Order
	err := c.BindJSON(&order)
	if err != nil {
		fmt.Println("status : ", err)
	}
	if err = Models.PlaceOrder(&order); err != nil {
		order.Status = "order cancelled!"
		fmt.Println(order.Status)
		c.JSON(http.StatusBadRequest, gin.H{"order_id": order.ID, "product_id": order.ProductId,
			"quantity": order.Quantity, "status": order.Status})
	} else {
		order.Status = "order placed"
		fmt.Println(order.Status)
		c.JSON(http.StatusOK, gin.H{"order_id": order.ID, "product_id": order.ProductId,
			"quantity": order.Quantity, "status": order.Status})
	}
}

func GetAllOrders(c *gin.Context) {
	var order []Models.Order
	err := Models.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		var resp []gin.H
		for _, ord := range order {
			resp = append(resp, gin.H{"order_id": ord.ID, "customer_id": ord.CustomerId, "quantity": ord.Quantity, "status": ord.Status})

		}
		c.JSON(http.StatusOK, resp)
	}
}
