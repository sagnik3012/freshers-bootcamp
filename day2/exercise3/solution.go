package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	state   sync.Mutex
	balance int
)

func deposit(wg *sync.WaitGroup) {

	state.Lock()
	defer state.Unlock()
	rand.Seed(time.Now().UnixNano())
	amount := rand.Intn(500)

	fmt.Printf("Current balance = %d, deposit amount = %d \n", amount, balance)
	balance += amount
	fmt.Println("New balance = ", balance)
	wg.Done()
}

func withdraw(wg *sync.WaitGroup) {

	state.Lock()
	defer state.Unlock()
	rand.Seed(time.Now().UnixNano())

	amount := rand.Intn(1000)

	fmt.Printf("Current balance = %d , Withdrawing %d  \n", balance, amount)
	if amount > balance {
		fmt.Printf("Insufficient funds! Transaction aborted!\n")
	} else {
		balance -= amount
		fmt.Println("New balance = ", balance)
	}
	wg.Done()
}

func main() {

	balance = 500
	var wg sync.WaitGroup
	for trn := 0; trn <= 10; trn++ {
		wg.Add(2)
		go deposit(&wg)
		go withdraw(&wg)

	}
	wg.Wait()
}
