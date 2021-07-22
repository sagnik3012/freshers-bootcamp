package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, ratings chan <- float32) {

	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	individualRating := rand.Intn(10)
	ratings <- float32(individualRating)
}

func main() {

	var wg sync.WaitGroup
	totalStudents := 200
	Ratings := make(chan float32,totalStudents)

	for i := 1; i <= totalStudents; i++ {
		wg.Add(1)
		go worker( &wg , Ratings)
	}

	wg.Wait()
	close(Ratings)

	totalRating := float32(0)
	for rate := range Ratings{
		totalRating += rate
	}
	averageRating := totalRating / float32(totalStudents)
	fmt.Println("Average rating of the teacher = ",averageRating)
}