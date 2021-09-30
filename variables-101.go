package main

import (
	"fmt"
)

// Can only be accessed within the current package
// Syntax: <var/const/func/<custom type>> <lowercase_letter><variableName>
var unexported = "unexported Variable, Constant, Function or Custom Types"

// can be accessed outside the current package
// Syntax: <var/const/func/<custom type>> <Uppercase_letter><variableName>
var Exported = "Exported Variable, Constant, Function or Custom Types"

// Scope of Variables
var packageScope = "I'd be visible through the package. And, I won't be exported."


func main(){
	fmt.Println()
	fmt.Println("---------------------- Variable Declaration & Assignment ---------------------------------")
	fmt.Println()

	var a int // var <variableName> <datatype>
	a = 1
	
	var b, c int // Comma [,] is supported in Variable declaration and assignment operations through out Go
 	b, c = 2, 3
	
	var d = 4 // var d, e = 4, 5

	e := 5 // e, f := 5, 6

	var f1 = 6
	var _g = 7
	var h_ = 8

	fmt.Println(a, b, c, d, e, f1, _g, h_)

	fmt.Println()
	fmt.Println("---------------------- TypeCasting 101 ---------------------------------")
	fmt.Println()

	var wholeNumber int = 1
	var fractionalNumber float64 = 96.4
	
	var wholeNumber_fractionalNumber int = int(fractionalNumber)
	var fractionalNumber_wholeNumber float64 = float64(wholeNumber)

	fmt.Println("Whole Number <- Fractional Number: ", wholeNumber_fractionalNumber, " <- ", fractionalNumber)
	fmt.Println("Fractional Number <- Whole Number: ", fractionalNumber_wholeNumber, " <- ", wholeNumber)

	fmt.Println()
	fmt.Println("---------------------- Typecast before operating over different types ---------------------------------")
	fmt.Println()

	fmt.Println("Whole Number + Fractional Number: ", float64(wholeNumber) + fractionalNumber)
	fmt.Println("Whole Number < Fractional Number: ", wholeNumber < int(fractionalNumber))

	fmt.Println()
	fmt.Println("---------------------- Scope of Variables ---------------------------------")
	fmt.Println()
	
	var functionScope = "I'd be visible until this function is in stack."

	{
		var blockScope = "I'd be visible only in this block."

		{
			var nestedBlockScope = "I'd be visible only in this block."

			fmt.Println("		---- At Inner Most Block ----")
			fmt.Println()

			fmt.Println("var packageScope 		package main	:", packageScope)
			fmt.Println("var functionScope		func(){}	:", functionScope)
			fmt.Println("var blockScope    		{}		:", blockScope)
			fmt.Println("var nestedBlockScope    	{{}}		:", nestedBlockScope)
		}

		fmt.Println()
		fmt.Println("		---- At Inner Block of Function ----")
		fmt.Println()

		fmt.Println("var packageScope 		package main	:", packageScope)
		fmt.Println("var functionScope		func(){}	:", functionScope)
		fmt.Println("var blockScope    		{}		:", blockScope)
		// fmt.Println("var nestedBlockScope    	{{}}		:", nestedBlockScope) 	// {{}} scope ended
	}

	fmt.Println()
	fmt.Println("		---- At Function Block ----")
	fmt.Println()

	fmt.Println("var packageScope 		package main	:", packageScope)
	fmt.Println("var functionScope		func(){}	:", functionScope)
	// fmt.Println("var blockScope    		{}		:", blockScope) 				// {} scope ended
	// fmt.Println("var nestedBlockScope    	{{}}		:", nestedBlockScope) 	// {{}} scope ended

	fmt.Println()
}