package main 

import (
	"fmt"
)

func main() {
	var mySlice []string //declares an empty slice
	fmt.Println(mySlice)

	var notes []string = make([]string, 7) //creates a slice with 7 empty strings initialized
	fmt.Println(notes)

	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes) //slice is populated only till 2nd index

	var myArray = [4]string{"fa","so","la","ti"} //array to populate slice note till the last index

	for index, _ := range notes {
		if index > 2 {
			notes[index] = myArray[index-3]
		}
	}
	fmt.Println(notes) //slice is now populated fully by adding array elements

	primes := []int{2, 3, 5} //slice literal with pre-determined elements
	fmt.Println(primes)

	underlyingArray := [5]string{"a","b","c","d","e"}
	slice1 := underlyingArray[1:3] //slice from 1st index to 3-1=2nd index
	slice2 := underlyingArray[:4] //slice from beginning index to 4-1=3rd index
	slice3 := underlyingArray[2:] //slice from 2nd index to last index
	slice4 := underlyingArray[:] //slice = entire underlying array
	fmt.Println("underlyingArray=", underlyingArray)
	fmt.Println("slice1=", slice1)
	fmt.Println("slice2=", slice2)
	fmt.Println("slice3=", slice3)
	fmt.Println("slice4=", slice4)
	
	// if underlying array values are changed, slices change too since their values depend on the underlying array
	underlyingArray[0] = "A"
	underlyingArray[2] = "C"
	fmt.Println("underlyingArray after change=", underlyingArray)
	fmt.Println("slice1 after change in underlyingArray=", slice1)
	fmt.Println("slice2 after change in underlyingArray=", slice2)
	fmt.Println("slice3 after change in underlyingArray=", slice3)
	fmt.Println("slice4 after change in underlyingArray=", slice4)
}