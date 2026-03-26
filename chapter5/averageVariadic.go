/*
Calculates the average from the floats passed to a variadic function from command line
*/

package main 

import (
	"fmt"
	"os"
	"log"
	"strconv"
)

func average(numbers ...float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	arguments := os.Args[1:]
	var values []float64
	for _, argument := range arguments {
		value, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, value)
	}
	// The line below will throw an error since average() expects a variadic float64, not a slice
	// fmt.Printf("arguments passed = %v and average = %.3f\n", arguments, average(values))
	// The fix is to add ellipsis (...) after the argument passed to average()
	// This is also used to append slice to a slice
	fmt.Printf("arguments passed = %v and average = %.3f\n", arguments, average(values...))
}