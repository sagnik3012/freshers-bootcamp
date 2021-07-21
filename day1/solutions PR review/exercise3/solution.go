package main

import "fmt"

type Employee interface {
	getSalary()
}

type FullTime struct {
	basicPay int
	daysWorked int
}
type Contractor struct {
	basicPay int
	daysWorked int
}
type Freelancer struct{
	basicPay int
	hoursWorked int
}

func (ft FullTime) getSalary() {
	fmt.Println( ft.daysWorked * ft.basicPay)

}

func ( ct Contractor) getSalary(){
	fmt.Println(ct.daysWorked*ct.basicPay)
}

func ( fr Freelancer) getSalary(){
	fmt.Println(fr.hoursWorked*fr.basicPay)
}
func main(){
	// create employees of three types
	fullTimeEmp := FullTime{500,28}
	contractorEmp := Contractor{100,28}
	freelancerEmp := Freelancer{basicPay: 10, hoursWorked: -1}
	fmt.Print("Enter number of hours worked by the freelancer : ")
	fmt.Scanf("%d",&freelancerEmp.hoursWorked)
	// create an array of employees
	employees := [] Employee{ fullTimeEmp,contractorEmp,freelancerEmp}
	// print the salary of each employee using getSalary function
	for _ , emp := range employees{
		emp.getSalary()
	}

}
