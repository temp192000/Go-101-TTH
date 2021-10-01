package main

import "fmt"

func main(){
	a := [5] int {0, 1, 2, 3, 4}

	s1 := a[0:3]
	s2 := a[2:5]

	s2[0] = 100

	fmt.Println()
	fmt.Println("-------- Before append() --------")
	fmt.Println()

	fmt.Println(a, s1, s2)

	fmt.Println()
	fmt.Println("-------- After append() --------")
	fmt.Println()

	s2 = append(s2, 5)
	s2[0] = 99

	fmt.Println(a, s1, s2)


	s := [] int {1, 2, 3}

	s = append(s, 4, 5)
	s = append(s, 6, 7, 8)

	fmt.Println()
	fmt.Println("-------- Best Practice --------")
	fmt.Println()
	fmt.Println(s)
	fmt.Println()

	fmt.Println("-------- for-range --------")
	fmt.Println()
	
	for index, value := range s {
		fmt.Println("Index	:", index, "Value	:", value)
	}
}