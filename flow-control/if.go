package main

import (
	"fmt"
)

func main() {

	fmt.Println("---------- if ----------")

	if true {
		fmt.Println("true")
	}

	if false {
		fmt.Println("false")
	}

	if 1 < 2 {
		fmt.Println("1 < 2")
	}

	if 1 > 2 {
		fmt.Println("1 > 2")
	}

	fmt.Println("---------- Boolean Operator ----------")

	if !true {
		fmt.Println("!true")
	}

	if !false {
		fmt.Println("!false")
	}

	if true && false {
		fmt.Println("true && false")
	}

	if true && true {
		fmt.Println("true && true")
	}

	if true || false {
		fmt.Println("true || false")
	}

	if true || true {
		fmt.Println("true || true")
	}

	fmt.Println("---------- if-else ----------")

	if true {
		fmt.Println("if")
	} else {
		fmt.Println("else")
	}

	if false {
		fmt.Println("if")
	} else {
		fmt.Println("else")
	}

	fmt.Println("---------- if - else if - else ----------")

	if true {
		fmt.Println("if")
	} else if true {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}

	if false {
		fmt.Println("if")
	} else if true {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}

	if false {
		fmt.Println("if")
	} else if false {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}

	fmt.Println("---------- Scope of Variables ----------")

	beforeIf := "Before if {}"

	if true {
		insideIf := "Inside if {}"

		fmt.Println("	Scope: Inside if {}")
		fmt.Println()

		fmt.Println(insideIf)
		fmt.Println(beforeIf)
	}

	fmt.Println()
	fmt.Println("	Scope: Outside if {}")
	fmt.Println()

	fmt.Println(beforeIf)
	// fmt.Println(insideIf) // Scope Ended
}
