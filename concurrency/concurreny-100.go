package main

import (
	"fmt"
	"time"
)

func longTask() {
	fmt.Println("Task Started")
	time.Sleep(3 * time.Second)
	fmt.Println("Task Completed")
}

func main(){
	// takes 9 seconds to complete. 3 func calls, 3 times = 3 * 3 = 9 seconds
	// longTask()
	// longTask()
	// longTask()

	go longTask()
	go longTask()
	go longTask()

	time.Sleep(4 * time.Second) // Main goroutine completes execution and child go routines get killed right after.
}