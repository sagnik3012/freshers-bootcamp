package Models

import (
	"fmt"
	"freshers-bootcamp/day3/restAPI/Config"
	_ "github.com/go-sql-driver/mysql"
)

//GetAllEntries Fetch all user data 1
func GetAllEntries(user *[]Student) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateEntry ... Insert New data //2
func CreateEntry(user *Student) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetEntryByID ... Fetch only one user by Id //3
func GetEntryByID(user *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateEntry ... Update user//4
func UpdateEntry(user *Student, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

//DeleteEntry ... Delete user//5
func DeleteEntry(user *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}

// 6
func DeleteAll() (err error) {
	Config.DB.Where("1=?", 1).Delete(&Student{})
	return nil
}
