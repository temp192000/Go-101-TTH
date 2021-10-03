package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTask() int {
	delay := rand.Intn(5)
	
	fmt.Println("Task Started")
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("Task Completed")
	
	return delay
}

func main(){
	rand.Seed(time.Now().Unix())
	// time := longTask()
	// time := go longTask() // Not possible to start a go routine
	fmt.Println("Took", time, "Seconds")
}