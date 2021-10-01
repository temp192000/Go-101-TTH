package main

import (
	"fmt"
)

func main(){

	fmt.Println("--------- Pointers 101 ---------")
	fmt.Println()
	
	var aValue float64 = 1.23
	var aPointer *float64 = &aValue

	fmt.Println(" aPointer :", aPointer)
	fmt.Println("*aPointer :", *aPointer)

	fmt.Println()
	fmt.Println("--------- Pass by Value ---------")
	fmt.Println()

	myNumber := 12.6
	
	fmt.Println("Before buggyHalve()	:", myNumber)
	buggyHalve(myNumber)
	fmt.Println("After buggyHalve()	:", myNumber)

	fmt.Println()
	fmt.Println("--------- Pass by Address ---------")
	fmt.Println()
	fmt.Println("Before halve()	:", myNumber)
	halve(&myNumber)
	fmt.Println("After halve()	:", myNumber)

	// fmt.Println()
	// fmt.Println("--------- Inefficient Usage ---------")
	// fmt.Println()

	// car := Car{
	// 	Doors:	4,
	// 	Transmission:	"automatic",
	// 	Odometer:	36000,
	// }

	// fatDescribe(car)

	// fmt.Println("--------- Efficient Usage ---------")
	// fmt.Println()

	// describe(&car)

}

func buggyHalve(number float64) {
	number = number / 2
	fmt.Println("buggyHalve()		:", number)
}

func halve(number *float64) {
	*number = *number / 2
	fmt.Println("halve()		:", *number)
}

// func fatDescribe(c Car){
// 	fmt.Print("A", c.Doors, "door")
// 	fmt.Print(c.Transmission, "car")
// 	fmt.Print("with", c.Odometer, "miles")
// }

// func describe(c *Car){
// 	fmt.Print("A", c.Doors, "door")
// 	fmt.Print(c.Transmission, "car")
// 	fmt.Print("with", c.Odometer, "miles")
// }