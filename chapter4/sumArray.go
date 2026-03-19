package main 

import (
	"fmt"
	"head-first-go/src/keyboard"
	"log"
)

func main() {
	fmt.Printf("Enter length of array (default is 5): ")
	length, err := keyboard.GetInt()
	if err != nil || length <= 0 {
		length = 5 // since length of array require compile-time constants
	}

	var numbers [100]int // larger array since we cannot set the length of the array at compile time
	var value, sum int
	for index := 0; index < int(length); index++ {
		fmt.Printf("Enter value for element %d = ", index+1)
		value, err = keyboard.GetInt()
		if err != nil {
			log.Fatal(err)
		}
		numbers[index] = value 
		sum += value
	}
	fmt.Println("The sum of the array elements =", sum)
}