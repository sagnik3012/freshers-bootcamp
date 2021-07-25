package Controllers

import (
	"fmt"
	"freshers-bootcamp/day3/restAPI/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetEntries ... Get all users
func GetEntries(c *gin.Context) {
	var user []Models.Student
	err := Models.GetAllEntries(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//CreateEntry ... Create Student
func CreateEntry(c *gin.Context) {
	var user Models.Student
	c.BindJSON(&user)
	err := Models.CreateEntry(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetEntryByID ... Get the student by id
func GetEntryByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.Student
	err := Models.GetEntryByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateEntry ... Update the student information
func UpdateEntry(c *gin.Context) {
	var user Models.Student
	id := c.Params.ByName("id")
	err := Models.GetEntryByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateEntry(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteEntry ... Delete the student
func DeleteEntry(c *gin.Context) {
	var user Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteEntry(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"ID" + id: "is deleted"})
	}
}

func DeleteAllEntries(c *gin.Context) {
	err := Models.DeleteAll()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, "table deleted!")
	}
}
