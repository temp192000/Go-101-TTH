package main

import (
	"fmt"
)

func main() {
	fmt.Print("You won: ")

	doorNumber := 2

	switch doorNumber {
	case 1:
		fmt.Println("a cat!")
	case 2:
		fmt.Print("a dog")
		fmt.Print(" and ")
		fallthrough
	case 3:
		fmt.Println("a hamster!")
	default:
		fmt.Println("nothing!")
	}
}
