/*
This program is to demonstrate functions and paramters
*/

package main

import (
	"fmt"
	"log"
)

// areaCalculate() accepts two parameters and return a single value
func areaCalculate (width float64, height float64) float64 {
	area := width * height
	return area
}

// paintNeeded() does the same but with error handling 
// paranthesis required for multiple return values
func paintNeeded (width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %.2f is invalid", height)
	}
	area := width * height
	return area / 10.00, nil
}

func main() {
	area := areaCalculate(4.2, -3.0)
	fmt.Printf("Area required to paint is %.2f\n", area)
	amount, err := paintNeeded(4.2,3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters of paint needed\n", amount)
}