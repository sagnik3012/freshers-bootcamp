package Controllers

import (
	"fmt"
	"freshers-bootcamp/day4/Models"
	"freshers-bootcamp/day4/Models/Product"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

func ConcurrentUpdate(c *gin.Context) {
	cc := c.Copy()
	//GetError := false
	//userReturned := Models.User{}
	ChannelErrorUp := make(chan bool, 100)
	ChannelProductUp := make(chan Models.Product, 100)
	go UpdateProduct(cc, ChannelErrorUp, ChannelProductUp)
	ans := <-ChannelErrorUp
	p := <-ChannelProductUp
	if ans {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// UpdateProduct ... Update the product information
func UpdateProduct(c *gin.Context, ErrorInUpdate chan bool, ProductUpdate chan Models.Product) {
	var product Models.Product
	mutex := &sync.Mutex{}
	mutex.Lock()
	id := c.Params.ByName("id")
	err := Product.GetProductByID(&product, id)
	if err != nil {
		ErrorInUpdate <- true
		ProductUpdate <- product
	}
	err = c.BindJSON(&product)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = Product.UpdateProduct(&product, id)
	if err != nil {
		//c.AbortWithStatus(http.StatusNotFound)
		ErrorInUpdate <- true
		ProductUpdate <- product
	} else {
		//c.JSON(http.StatusOK, product)
		ErrorInUpdate <- false
		ProductUpdate <- product
	}
	mutex.Unlock()
}
