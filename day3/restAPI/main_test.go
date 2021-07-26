package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"freshers-bootcamp/day3/restAPI/Config"
	"freshers-bootcamp/day3/restAPI/Controllers"
	"freshers-bootcamp/day3/restAPI/Models"
	"freshers-bootcamp/day3/restAPI/Routes"
	"github.com/jinzhu/gorm"
)

func TestGetEntries(t *testing.T) {
	//open sql connection by gorm
	Config.DB, _ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	fmt.Println("Starting testing GetEntries...")
	//setup router via gin
	router := Routes.SetupRouter()
	router.GET("/student", Controllers.GetAllEntries)

	//send get request
	req, _ := http.NewRequest("GET", "/studentdb/student", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	expected := `[{"id":25,"stu_id":1001,"firstname":"sagnik","surname":"saha","dob":"30.12.1997","address":"e-28, srinagar west","subject":"Maths","marks":99},{"id":28,"stu_id":1002,"firstname":"sagnik","surname":"ghosal","dob":"31.12.1997","address":"a-7, Newtown","subject":"Maths","marks":90}]`
	if resp.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Body.String(), expected)
	}
}

func TestCreateEntry(t *testing.T) {
	//Setup and open sql db
	Config.DB, _ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	//setup router
	router := Routes.SetupRouter()
	router.POST("/student", Controllers.CreateEntry)
	newStudent := Models.Student{
		Stu_id:    2222,
		FirstName: "Local",
		LastName:  "Test",
		DOB:       "1.1.1990",
		Address:   "Pluto",
		Subject:   "History",
		Marks:     75,
	}

	responseBody, _ := json.Marshal(newStudent)
	req, _ := http.NewRequest("POST", "/studentdb/student", bytes.NewBuffer(responseBody))
	//send request
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	//read body of response
	//fmt.Println(resp.Body.String(), "some text posted")
	expected := `{"id":42,"stu_id":2222,"firstname":"Local","surname":"Test","dob":"1.1.1990","address":"Pluto","subject":"History","marks":75}`
	//check for test case by status code
	if resp.Body.String()[20:] != expected[20:] {
		t.Errorf("handler returned unexpected body: got \n %v want \n %v",
			resp.Body.String(), expected)
	}
}

func TestUpdateEntry(t *testing.T) {
	//Setup and open sql db
	Config.DB, _ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	//setup router
	router := Routes.SetupRouter()
	router.PUT("/student/25/", Controllers.UpdateEntry)

	newStudent := Models.Student{LastName: "Updated", Address: "Pluto", Marks: 175}
	responseBody, _ := json.Marshal(newStudent)

	//send request
	req, _ := http.NewRequest("PUT", "/studentdb/student/25/", bytes.NewBuffer(responseBody))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	//read body of response

	expected := `{"id":25,"stu_id":1001,"firstname":"sagnik","surname":"Updated","dob":"31.12.1997","address":"Pluto","subject":"History","marks":175}`

	//check for test case by status code
	if resp.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v \n want %v \n", resp.Body.String(), expected)
	}
}
