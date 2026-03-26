/*Calculates the average from the floats passed as arguments through command-line*/

package main 

import (
	"fmt"
	"os"
	"log"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	var sum float64
	for _, argument := range arguments {
		value, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += value
	}
	average := sum / float64(len(arguments))
	fmt.Printf("arguments passed = %v and average = %.3f\n", arguments, average)
}