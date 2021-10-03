package main

import (
	"fmt"
)

type Minutes int

// Receiver is passed by value
// func (m Minutes) Increment() {
// 	m = (m + 1) % 60
// }

// Receiver is passed by address : Pointer Receiver
func (m *Minutes) Increment() {
	*m = (*m + 1) % 60
}

func main(){
	minutes := Minutes(58)

	for i := 1; i <= 3; i++ {
		minutes.Increment()
		// (&minutes).Increment() // As Increment() is a Pointer receiver Method, it's Implicit.
		fmt.Println(minutes)
	}
}