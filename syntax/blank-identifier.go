package main

import (
	"fmt"
	"os"
)

func main(){
	fileInfo, err := os.Stat("existent.txt")
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(fileInfo.Size())
	}

	// fileInfo, _ := os.Stat("nonexistent.txt") //  ERROR: no new variables on left side of :=
	fileInfo, err = os.Stat("nonexistent.txt")
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(fileInfo.Size())
	}
}