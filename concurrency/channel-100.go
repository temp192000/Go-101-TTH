package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTask(channel chan int) {
	delay := rand.Intn(5)
	
	fmt.Println("Task Started")
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("Task Completed")
	
	channel <- delay
}

func main(){
	rand.Seed(time.Now().Unix())
	
	channel := make(chan int)

	for i := 1; i <= 3; i++ {
		go longTask(channel)
	}

	for i := 1; i <= 3; i++ {
		fmt.Println("Took", <- channel, "seconds")
	}
}