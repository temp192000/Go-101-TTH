package main

import (
	"fmt"
	"math"
	// "time"
)

var myNumber = 1.23

func main() {
	// := short hand operator to declare and intialize
	roundUp := math.Ceil(myNumber)
	roundDown := math.Floor(myNumber)
	fmt.Println(roundUp, roundDown)
	// otherFunction()
}

// func otherFunction(){
// 	fmt.Println("The Other Function")
// }