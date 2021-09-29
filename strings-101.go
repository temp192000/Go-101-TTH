package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
	// "github.com/golang/example/stringutil" // actually demo
	// "go-101/example/stringutil" // cheap fix
)


func main(){
	// fmt.Println("Hello, Gophers!")
	fmt.Println(stringutil.Reverse("Was it a car or a cat I saw?"))
	fmt.Println(stringutil.Reverse("?was I tac a ro rac a ti saW"))
}