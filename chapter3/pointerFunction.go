package main

import "fmt"

// returns a pointer for myFloat
func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

// returns the value of the pointer passed through arguments from main()
func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

// prints the value at the pointer returned from createPointer()
// calls printPointer() and passes a boolean value to it
func main() {
	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)
	var myBool bool = true
	printPointer(&myBool)
}