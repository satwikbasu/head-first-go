/*
This program demonstrates creating a defined type with an underlying type as struct
*/

package main

import(
	"fmt"
)

// type defined is "part" while underlying type is a struct
type part struct {
	description string
	count int
}

// type defined is "car" while underlying type is a struct
type car struct {
	name string
	topSpeed float64
}

// function that takes user-defined type as argument
func showInfo(p car) {
	fmt.Println("Name:", p.name)
	fmt.Println("Top speed:", p.topSpeed)
}

// function that returns user-defined type after initializing the values passed
func minimumOrder(description string, count int) part {
	var p part
	p.description = description
	p.count = count
	return p
}

func main() {
	var porsche car // porsche is a variable of type car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	showInfo(porsche) // prints the name and top speed of porsche

	var bolts part // bolts is a variable of type part
	bolts.description = "Hex Bolts"
	bolts.count = 69
	p := minimumOrder(bolts.description, bolts.count) // prints the description and count of bolts
	fmt.Printf("Bolts description: %s\nMinimum order count: %d\n", p.description, p.count)
}