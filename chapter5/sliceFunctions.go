package main

import (
	"fmt"
)

func main() {
	
	var testSlice []int = make([]int, 3, 5) // creates a slice of length 3 capacity 5
	fmt.Println("\ntestSlice=", testSlice, "length=", len(testSlice), "capacity=", cap(testSlice))
	// if elements more than its capacity are added to slice,
	// a larger underlying array is assigned 
	testSlice = append(testSlice, 1, 2, 3, 4) // appended 4 elements which exceeds previous capacity
	// length becomes 7, capacity is doubled to 10
	fmt.Println("testSlice=", testSlice, "length=", len(testSlice), "capacity=", cap(testSlice))


	var mySlice []int // declares a null or empty slice of length=capacity=0
	fmt.Println("\nmySlice=", mySlice, "length=", len(mySlice), "capacity=", cap(mySlice))
	mySlice = append(mySlice, 1, 2, 3) // length and capacity increase but both are of the same value
	fmt.Println("mySlice=", mySlice, "length=", len(mySlice), "capacity=", cap(mySlice))
	mySlice = append(mySlice, 4, 5, 6, 7) // length increases 3 -> 7, capacity increases 3 -> 8 
	fmt.Println("mySlice=", mySlice, "length=", len(mySlice), "capacity=", cap(mySlice))

	s1 := []int {1, 2, 3}
	s2 := []int {4, 5, 6}
	s3 := append(s1, 4, 5, 6) // appends individual elements
	s4 := append(s1, s2...) // appends slices
	fmt.Println("\ns3 = append(s1, <elements>) =", s3, "length=", len(s3), "capacity=", cap(s3))
	fmt.Println("s4 = append(s1, s2...) =", s4, "length=", len(s4), "capacity=", cap(s4))
}