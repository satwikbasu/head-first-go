/*
This program is to demonstrate overriding "pass by value" using pointers inside functions
*/

package main 

import (
	"fmt"
)

func main() {
	amount := 6
	fmt.Println("Original number in main() before double() is called =",amount) // amount is unchanged
	double(&amount) // pointer of amount is sent to double()
	fmt.Println("Original number in main() after double() is called =",amount) // amount changes (would not happen when pass by value)
}

func double(number *int) {
	value := *number // value at pointer number
	fmt.Println("Value present at pointer before change =", *number)
	*number = value * 2 // value present at pointer is changed
	fmt.Println("Value present at pointer after change =", *number) 
}