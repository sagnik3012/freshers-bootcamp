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

	fullTimeEmp := FullTime{500,28}
	contractorEmp := Contractor{100,28}
	freelancerEmp := Freelancer{ 10,-1}

	fmt.Print("Enter number of hours worked by the freelancer : ")
	fmt.Scanf("%d",&freelancerEmp.hoursWorked)

	employees := [] Employee{ fullTimeEmp,contractorEmp,freelancerEmp}
	for _ , employee := range employees{
		employee.getSalary()
	}

}
