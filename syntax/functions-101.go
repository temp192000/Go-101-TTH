package main

import "fmt"

func Exported(){
	fmt.Println("Exported Function")
}

func unexported(){
	fmt.Println("UnExported Function")
}

func ExportedMultipleWords(){
	fmt.Println("Exported Function With Multiple Words")
}

func unexportedMultipleWords(){
	fmt.Println("UnExported Function With Multiple Words")
}

func main(){
	customFunc()
	
	fmt.Print("Square		: ")
	square(4)
	
	fmt.Print("Addition	: ")
	add(8, 8)

	fmt.Print("Subtraction	: ")
	subtract(32, 16)
}

func customFunc(){
	fmt.Println("Running a user-defined function.")
}

func square(number int){
	fmt.Println(number * number)
}

func add(augend float64, addend float64){
	fmt.Println(augend + addend)
}

func subtract(minuend, subtrahend float64){
	fmt.Println(minuend - subtrahend)
}