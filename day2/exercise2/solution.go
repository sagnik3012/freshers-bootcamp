package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getRating(wg *sync.WaitGroup, ratings chan<- float32) {

	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	rating := rand.Intn(10)
	ratings <- float32(rating)
}

func main() {

	var wg sync.WaitGroup
	totStudents := 200
	Ratings := make(chan float32, totStudents)

	for i := 1; i <= totStudents; i++ {
		wg.Add(1)
		go getRating(&wg, Ratings)
	}

	wg.Wait()
	close(Ratings)

	totRating := float32(0)
	for rate := range Ratings {
		totRating += rate
	}
	avgRating := totRating / float32(totStudents)

	fmt.Println("Average rating of the teacher = ", avgRating)
}
