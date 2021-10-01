package main

import "go-101/userFunctions"

func main(){
	userFunctions.Exported()
	userFunctions.ExportedMultipleWords()

	// userFunctions.unexported()
	// userFunctions.unexportedMultipleWords()
	// userFunctions.customFunc()
}