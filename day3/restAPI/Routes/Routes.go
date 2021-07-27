package Routes

import (
	"freshers-bootcamp/day3/restAPI/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/studentdb")
	{
		grp1.GET("student", Controllers.GetAllEntries)
		grp1.POST("student", Controllers.CreateEntry)
		grp1.GET("student/:id", Controllers.GetEntryByID)
		grp1.PUT("student/:id", Controllers.UpdateEntry)
		grp1.DELETE("student/:id", Controllers.DeleteEntry)
		grp1.DELETE("student", Controllers.DeleteAllEntries)
	}
	return r
}
