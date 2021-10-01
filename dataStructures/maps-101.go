package main

import "fmt"

func main(){
	ages := map [string] float64 {}
	
	ages["Alice"] = 21
	ages["Bob"] = 35
	ages["Cia"] = 20
	
	fmt.Println()
	fmt.Println("----------- Maps 101 -----------")
	fmt.Println()

	fmt.Println(ages)
	fmt.Println(ages["Alice"], ages["Bob"], ages["Cia"])

	fmt.Println()
	fmt.Println("----------- Prepopulate Map -----------")
	fmt.Println()

	// students := map [string] float64 {
	// 	"Ranjith Sr. ": 34,
	// 	"Ranjith Jr. ": 21,
	// 	"Ranjith": 12
	// }

	// students := map [string] float64 {
	// 	"Ranjith Sr. ": 34,
	// 	"Ranjith Jr. ": 21,
	// 	"Ranjith": 12,
	// }

	students := map [string] float64 {
		"Ranjith Sr. ": 34,
		"Ranjith Jr. ": 21,
		"Ranjith": 12}

	fmt.Println()
	fmt.Println("----------- for-range Map-----------")
	fmt.Println()

	for name, age := range students {
		fmt.Println(name, age)
	}

	fmt.Println()
	fmt.Println("----------- Blank Identifier -----------")
	fmt.Println()

	for name, _ := range students {
		fmt.Println(name)
	}

	fmt.Println()
}