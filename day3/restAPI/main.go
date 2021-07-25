package main

import (
	"fmt"
	"freshers-bootcamp/day3/restAPI/Config"
	"freshers-bootcamp/day3/restAPI/Models"
	"freshers-bootcamp/day3/restAPI/Routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Student{})
	r := Routes.SetupRouter()
	r.Run()
}
