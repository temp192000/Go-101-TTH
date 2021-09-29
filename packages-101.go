package main
	
import (
	"fmt"
	"go-101/welcome"
)

func main(){
	fmt.Println(welcome.English)
	// fmt.Println(welcome.tamil) // ERROR UNEXPORTED
	// fmt.Println(welcome.tulu) // ERROR UNEXPORTED + UNDECLARED
}