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

func printInfo(s subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?:", s.active)
}

func defaultSubscriber (name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}

func main() {
	subscriber1 := defaultSubscriber("John Doe")
	subscriber1.rate = 4.99
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Maria Po")
	printInfo(subscriber2)
}