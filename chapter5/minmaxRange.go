/* This program takes a min number and a max number as argument
and takes a variable number of values as argument and returns a slice of only those values
where value <= max & value >= min */

package main

import (
	"fmt"
)

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	fmt.Println(inRange(1, 100, -21.5, 23, 69, -31, 99.999999))
	fmt.Println(inRange(-10, 10, 4.1, 12, -12, -5.2))
}