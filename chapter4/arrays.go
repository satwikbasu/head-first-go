package main

import "fmt"

func main() {

	var Numbers [5]int
	fmt.Println(Numbers) // empty array is populated with default value of the type

	var notes [7]string = [7]string{"do","re","mi","fa","so","la","ti"} // array declaration with elements
	primes := [5]int{2, 3, 5, 7, 11} // short-hand array declaration where type is derived
	fmt.Println(notes[0])
	fmt.Println(len(primes)) // len(array-variable) results in length of the array 

	text := [3]string{
		"This is a series of long strings",
		"which would be awkward to place",
		"together on a single line", // comma at the end is required
	}
	fmt.Println(text)
	
	// for index, element := range array is used for traversing an array
	for index, note := range notes {
		fmt.Println(index, note)
	}

	// The array Numbers has default values stored inside it
	// the loop below populates it with index values e.g. array[0] = 0
	// the _ variable allows omitting an unused variable like element/index 
	for index, _ := range Numbers {
		Numbers[index] = index + 1
		fmt.Println(Numbers[index])
	}
}