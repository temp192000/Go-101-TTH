package main

import (
	"fmt"
	"reflect"
)

type myType float64

func (m myType) myMethod() myType {
	myValueResult := m
	fmt.Println("In myMethod: I'll take myType values, process them and return myType value back again.")
	return myValueResult
}

func main(){
	myValue := myType(99.9)
	// myValue.myMethod()
	fmt.Println("Return Type from myMethod:", reflect.TypeOf(myValue.myMethod()))
}