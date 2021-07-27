package main

import(
	"fmt"

	"github.com/jinzhu/gorm"

	"freshers-bootcamp/day4/Config"
	"freshers-bootcamp/day4/Models"
	"freshers-bootcamp/day4/Routes"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Product{}, &Models.Order{}, &Models.Customer{})
	r := Routes.SetupRouter()
	r.Run()
}
