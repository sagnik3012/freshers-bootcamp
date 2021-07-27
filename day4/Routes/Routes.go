package Routes

import (
	"freshers-bootcamp/day4/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/")
	{
		grp1.GET("product/", Controllers.GetAllProducts)
		grp1.POST("product/", Controllers.CreateProduct)
		grp1.GET("product/:id", Controllers.GetProductByID)
		grp1.PATCH("product/:id", Controllers.PatchProduct)
		grp1.POST("order/", Controllers.PlaceOrder)
		grp1.GET("order/",Controllers.GetAllOrders)
		grp1.GET("order/:id", Controllers.GetOrderByID)

	}
	return r
}
