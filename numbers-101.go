package main

import (
	"fmt"
	"reflect"
	"net"
	"time"
)

func main(){
	// Built-in Data types
	// fmt.Println("Hewwo" + 123) // ERROR: mismatched type
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(1.2))
	fmt.Println(reflect.TypeOf("Go"))
	fmt.Println(reflect.TypeOf(true))

	// Custom Data types in in-built Packages
	fmt.Println(reflect.TypeOf(net.IPv4(127, 0, 0, 1)))
	fmt.Println(reflect.TypeOf(time.Now()))
}