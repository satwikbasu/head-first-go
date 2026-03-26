/*Calculates the maximum of the floats passed into a variadic function*/

package main

import (
	"fmt"
	"math"
)

func max(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func main() {
	fmt.Println(max(12.5, 34.2, 54.1, -65, -121, 69))
	fmt.Println(max(90.5, 90.55, 90.555))
}