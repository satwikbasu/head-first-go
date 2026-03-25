package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var values [3]float64
	file, err := os.Open("chapter4/sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	index := 0
	var sum float64 = 0.0
	for scanner.Scan() {
		values[index], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += values[index]
		index++
	}

	average := sum / 3.0
	fmt.Printf("Sum of the array is %.3f and average is %.3f\n", sum, average)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}