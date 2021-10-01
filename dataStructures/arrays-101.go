package main

import (
	"fmt"
)

func main(){
	
	fmt.Println()
	fmt.Println("---------- Arrays 101 ----------")
	fmt.Println()

	// var months [3] string
	
	// months[0] = "Jan"
	// months[1] = "Feb"
	// months[2] = "Mar"

	// var salesByMonth [3] float64

	// salesByMonth[0] = 987.65
	// salesByMonth[1] = 5432.10
	// salesByMonth[2] = 0123.45

	// fmt.Println(months[0], salesByMonth[0])
	// fmt.Println(months[1], salesByMonth[1])
	// fmt.Println(months[2], salesByMonth[2])

	fmt.Println()
	fmt.Println("---------- Pre-Populating Arrays ----------")
	fmt.Println()

	months := [3] string {"Jan", "Feb", "Mar"}
	salesByMonth := [3] float64 {987.65, 5432.10, 0123.45}

	fmt.Println("---------- Traversing Arrays ----------")
	fmt.Println()

	for i := 0; i < len(months); i++ {
		fmt.Println(months[i], salesByMonth[i])
	}

	fmt.Println("---------- for - range loop ----------")
	fmt.Println()

	for _, month := range months {
		fmt.Println(month)
	}

	fmt.Println()
	fmt.Println("Arrays neither can't grow nor shrink. Thus, we've Slices")
}