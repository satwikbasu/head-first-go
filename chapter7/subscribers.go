/*
This program demonstrates how passing pointers of types (or structs) is better than 
passing the whole struct to a function to modify any of its fields

printInfo is passed a copy of the struct
defaultSubscriber and applyDiscount use pointers of the struct instead
*/

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

// prints all the fields of the type subscriber
func printInfo(s subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?:", s.active)
}

// initializes default values to variables of type subscriber using a pointer of the struct
func defaultSubscriber (name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

// applyDiscount function demonstrates how passing a pointer 
// instead of a whole copy of the struct is better for memory usage
// The function applies discounter rate to a subscriber type using a pointer passed
func applyDiscount(s *subscriber) {
	s.rate = 4.99 // or (*s).rate = 4.99 works too 
}

func applyFakeDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("John Doe")
	subscriber1.rate = 6.99 // values of fields can be changed in a type using direct assignment
	// applyFakeDiscount(subscriber1) // but they cannot be changed by passing the type directly
	printInfo(*subscriber1) // struct is passed intead of pointer for printInfo function

	subscriber2 := defaultSubscriber("Maria Po")
	applyDiscount(subscriber2) // however, they can be changed using a pointer
	printInfo(*subscriber2)
}