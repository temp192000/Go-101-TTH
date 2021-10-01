package main

import (
	"fmt"
)

func main(){
	fmt.Println("------------- Incrementation -------------")
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("------------- Decrementation -------------")
	for i := 3; i > 0; i-- {
		fmt.Println(i)
	}

	fmt.Println("------------- Step through -------------")
	for i := 1; i <= 5; i += 2 {
		fmt.Println(i)
	}

	fmt.Println("------------- Scope of 'for' loop variables -------------")
	
	beforeLoop := "Before Loop"

	for i:= 1; i <= 5; i++ {
		insideLoop := "Inside Loop"
		fmt.Println(i)
		// fmt.Println(insideLoop)
	}

	fmt.Println(beforeLoop)
	// fmt.Println(i) // ERROR: Scope considered to inside for and Ended at so
	// fmt.Println(insideLoop) // ERROR: Scope Ended in for
}