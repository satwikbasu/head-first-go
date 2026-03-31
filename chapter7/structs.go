/*
This program demonstrates basics of creating structs
and assinging values to its fields using the dot operator
*/

package main

import (
	"fmt"
)

func main() {
	// struct named myStruct is created with fields number, word and toggle
	var myStruct struct {
		number float64
		word string
		toggle bool
	}
	
	// myStruct fields are assigned values
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true

	// assigned values of myStruct is printed
	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)
}