package main

import (
	"fmt"
)

// user defined type is subscriber and underlying type is struct
type subscriber struct {
	name string
	rate float64
	active bool
}

func main() {
	var subscriber1 subscriber
	subscriber1.name = "Satwik Basu"
	fmt.Println("Name:", subscriber1.name)
	var subscriber2 subscriber
	subscriber2.name = "John Doe"
	fmt.Println("Name:", subscriber2.name)
}