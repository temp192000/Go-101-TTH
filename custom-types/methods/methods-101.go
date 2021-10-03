package main

import (
	"fmt"
	"strings"
)

type Title string

// t is from T(first letter) of the Type Title. It's a convention for naming receiver type

func (t Title) FixCase() string {
	return strings.Title(string(t))
}

func main(){
	movieName := Title("the matrix")

	movieName = Title(movieName.FixCase()) // movieName.FixCase() returns a string, either use it as string or convert it to Title (Custom Type)
	fmt.Println(movieName)
}