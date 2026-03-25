package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"log"
)

func main() {
	var values []float64
	file, err := os.Open("chapter5/sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var sum float64
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, value)
		sum += value
	}

	average := sum / float64(len(values))
	fmt.Printf("Values extracted from sample.txt = %v\n", values)
	fmt.Printf("Sum = %v, Average = %v\n", sum, average)
}