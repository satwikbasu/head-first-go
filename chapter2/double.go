/*
This program is to demonstrate "pass by value" parameters in functions
*/

package main 

import (
	"fmt"
)

func main() {
	amount := 6
	double(amount) // copy of number is sent to double()
	fmt.Println("Original number in main() =",amount) // amount remains unchanged
}

func double(number int) {
	number *= 2 // number holds a copy value of amount from main()
	fmt.Println("Altered number in double() =",number) 
}