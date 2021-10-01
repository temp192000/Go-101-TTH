package main

import (
	"fmt"
	"math"
	"log"
)

func main(){
	squareRoot, err := squareRoot(16)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(squareRoot)
}

func squareRoot(number float64) (float64, error){
	if number < 0 {
		return 0, fmt.Errorf("Invalid Argument: cannot take square root of a negative number")
	}

	return math.Sqrt(number), nil
}