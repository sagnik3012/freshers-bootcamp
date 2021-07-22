package main

import (
	"fmt"
	"sync"
)

var(
	mutex sync.Mutex
	balance int
)

func Deposit(  wg *sync.WaitGroup){
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Printf("Enter amount to deposit : ")
	var depositAmount int
	fmt.Scanf("%d",&depositAmount)
	fmt.Printf("Initial balance = %d , Depositing %d\n", balance , depositAmount)
	balance += depositAmount
	fmt.Printf("Deposit successful! New balance = %d\n",balance)
	wg.Done()
}

func Withdraw( wg *sync.WaitGroup){
	mutex.Lock()
	//defer mutex.Unlock()
	var withdrawAmount int
	fmt.Printf("Enter amount to withdraw : ")
	fmt.Scanf("%d",&withdrawAmount)
	if withdrawAmount > balance {
		fmt.Printf("Trancastion Aborted!!!!!\n Attempt to withdraw %d , but balance = %d\n", withdrawAmount, balance)
		//panic("Transation aborted!")
		mutex.Unlock()
		return

	}
	balance -= withdrawAmount
	fmt.Printf(" Withdraw successful. New balance = %d\n",balance)
	mutex.Unlock()
	wg.Done()

}
func main(){
	var wg sync.WaitGroup
	balance = 500
	for transactions := 0 ; transactions <= 10 ; transactions++{
		wg.Add(2)
		go Withdraw(&wg)
		go Deposit(&wg)
	}
	wg.Wait()

}