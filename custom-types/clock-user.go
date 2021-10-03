package main

import (
	// "Go-101/clock" // Custom Type: Minutes
	"go-101/clock-struct" // Encapsulation by struct, getter() & setter() for accessing fields
)

func main(){

	// Custom Type: Minutes

	// minutes := clock.Minutes(58) 	// Custom Type: Minutes
	// minutes += 5 					// Minute is an invalid value as 58 + 5 = 63 // Custom Type: Minutes
	// minutes.Display() 				// Custom Type: Minutes


	// Encapsulation by struct, getter() & setter() for accessing fields

	minutes := clock.Minutes{}
	minutes.Set(63)

	for i := 1; i <= 3; i++ {
		minutes.Increment()
		minutes.Display()
	}
}