package main

import (
	"fmt"
)

func main(){
	a := [5] int {0, 1, 2, 3, 4 }

	s1 := a[0:3]
	s2 := a[2:5]

	fmt.Println()
	fmt.Println(a, s1, s2)

	fmt.Println()
	fmt.Println("---------- Modify Slices ----------")
	fmt.Println()
	
	a[2] = 100
	fmt.Println(a, s1, s2)

	fmt.Println()
	fmt.Println("---------- Re-Slicing ----------")
	fmt.Println()

	s1 = s1[0:4]
	// s2 = s2[0:4] // ERROR: slice bounds out of range. ToDo: Growing Slice at the beginning, shrinking at the rear.

	fmt.Println(a, s1, s2)
	
	fmt.Println()
	fmt.Println("---------- len(), cap() of Slices ----------")
	fmt.Println()

	fmt.Println("len(a)	:", len(a), "cap(a)	:", cap(a))
	fmt.Println("len(s1)	:", len(s1), "cap(s1)	:", cap(s1))
	fmt.Println("len(s2)	:", len(s2), "cap(s2)	:", cap(s2))

	s2 = append(s2, 5)

	fmt.Println()
	fmt.Println(a, s1, s2)

	fmt.Println()
	fmt.Println("len(a)	:", len(a), "cap(a)	:", cap(a))
	fmt.Println("len(s1)	:", len(s1), "cap(s1)	:", cap(s1))
	fmt.Println("len(s2)	:", len(s2), "cap(s2)	:", cap(s2))
	fmt.Println()

	// a[3] = 200
	s2[0] = 800
	fmt.Println(a, s1, s2)

	// fmt.Println()
	// fmt.Println("----------  ----------")
}