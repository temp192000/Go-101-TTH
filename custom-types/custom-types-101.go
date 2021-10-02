package main

import "fmt"

type Minutes int
type Hours int
type Weight float64
type Title string
type Answer bool

func main(){
	minutes := Minutes(35)
	hours := Hours(2)
	weight := Weight(62.5)
	title := Title("The Matrix")
	answer := Answer(true)

	fmt.Println()
	fmt.Println("---------- Custom Types 101 ----------")
	fmt.Println()
	
	fmt.Println("The Time is", hours, "hours and", minutes, "minutes.")
	fmt.Println("Minutes :", minutes, "Hours :", hours, "Weight :", weight, "Title :", title, "Answer :", answer)

	fmt.Println()
	fmt.Println("---------- Operations over Custom Types ----------")
	fmt.Println()

	minutes += 10
	fmt.Println("Added 10 minutes :", minutes)

	fmt.Println()
	fmt.Println("---------- Comparisions over Custom Types ----------")
	fmt.Println()
	
	weight += 40

	if weight > Weight(79.9) {
		fmt.Println("Overweight")
	}

	if weight > 99.9 {
		fmt.Println("Super Overweight")
	}

	// if minutes > hours {
	// 	fmt.Println("Yes, this is not invalid!")
	// }

	// fmt.Println("----------  ----------")
	// fmt.Println()
}