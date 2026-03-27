package main 

import (
	"fmt"
)

func main() {
	var ranks map[string]int
	// var myMapVariable map[key-type]value-type declares a map variable
	fmt.Printf("ranks=%v type=%T\n", ranks, ranks)
	// ranks["gold"] = 1 will cause panic since map is not created, only declared
	ranks = make(map[string]int) // actually creates a map called ranks
	ranks["gold"] = 1 // works fine since map is actually created before key-value is assigned
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Printf("ranks=%v type=%T\n", ranks, ranks)
	
	
	myMap := map[string]int {"A": 1, "B": 2} // short-hand map literal declaration
	fmt.Printf("myMap=%v type=%T\n", myMap, myMap)
	var elements = map[string]string { // multi-line map literal declaration
		"H" : "Hydrogen",
		"He" : "Helium",
	}
	fmt.Printf("elements=%v type=%T\n", elements, elements)

	numbers := make(map[string]int)
	numbers["I've been assigned"] = 69
	fmt.Printf("%#v\n", numbers["I've been assigned"])
	fmt.Printf("%#v\n", numbers["I haven't been assigned"]) // default value of type is assigned if key does not exist
	
	fmt.Println("myMap = ",myMap)
	myMap["A"]++ // increasing from 0 to 1
	myMap["A"]++ // increasing from 1 to 2
	myMap["C"]++ // key does not exist but still incremented from 0 to 1
	fmt.Println("myMap = ",myMap)
}