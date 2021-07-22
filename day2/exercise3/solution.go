package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	state     sync.Mutex
	balance int
)

func Deposit( wg *sync.WaitGroup) {

	state.Lock()
	defer state.Unlock()
	rand.Seed(time.Now().UnixNano())
	amount := rand.Intn(500)
	fmt.Printf("Current balance = %d, deposit amount = %d \n",amount,balance)
	balance += amount
	fmt.Println("New balance = ",balance)
	wg.Done()
}

func Withdraw(wg * sync.WaitGroup)  {

	state.Lock()
	defer state.Unlock()
	rand.Seed(time.Now().UnixNano())

	amount := rand.Intn(1000)

	fmt.Printf("Current balance = %d , Withdrawing %d  \n",balance, amount)
	if amount > balance {
		panic("Insufficient funds! Transaction aborted!")
	}

	balance -= amount
	fmt.Println("New balance = ",balance)
	wg.Done()
}

func main(){

	balance = 500
	var waitgroup sync.WaitGroup
	for transactions := 0 ; transactions <= 10 ; transactions ++{

		waitgroup.Add(2)
		go Deposit(&waitgroup)
		go Withdraw(&waitgroup)

	}
	waitgroup.Wait()
}