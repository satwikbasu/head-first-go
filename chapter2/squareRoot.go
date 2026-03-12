/*
This program is to demonstrate input and error handling
*/

package main

import (
	"fmt"
	"math"
	"bufio"
	"os"
	"strings"
	"strconv"
	"log"
)

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("negative number cannot have square root")
	}
	return math.Sqrt(number), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	input, err := reader.ReadString('\n')
	if (err != nil){
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	value, err := strconv.ParseFloat(input, 64)
	if (err != nil){
		log.Fatal(err)
	}

	root, err := squareRoot(value)
	if (err != nil){
		log.Fatal(err)
	} else {
		fmt.Printf("Square root of %.3f is %.3f\n", value, root)
	}
}