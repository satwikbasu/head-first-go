package main 

import (
	"fmt"
	"head-first-go/src/keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f degrees Fahrenheit is %.2f degrees Celsius.\n", fahrenheit, celsius)
}