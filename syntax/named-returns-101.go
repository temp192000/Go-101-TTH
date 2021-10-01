package main

import "fmt"

func main(){
	fmt.Println("Square		:", square(4))
	fmt.Println("Add		:", add(8, 8))
	fmt.Println("Subtract	:", subtract(32, 16))
}

func square(number int) int{
	return number * number
}

func add(augend float64, addend float64)(sum float64){
	return augend + addend
}

func subtract(minuend, subtrahend float64)(difference float64){
	difference = minuend - subtrahend
	return
}