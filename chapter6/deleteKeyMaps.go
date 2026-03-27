package main 

import (
	"fmt"
)

func main() {
	var ok bool
	ranks := make(map[string]int)
	var rank int

	rank, ok = ranks["bronze"] // ok=false since value is not assigned yet
	fmt.Printf("Before assigning, ranks[bronze] = %v, ok = %t\n", rank, ok)

	ranks["bronze"] = 3 // assigning a value to "bronze" key
	rank, ok = ranks["bronze"] // ok=true since value was assigned
	fmt.Printf("After assigning, ranks[bronze] = %v, ok = %t\n", rank, ok)
	
	delete(ranks, "bronze") // delete() removes the "bronze" key from the 
	rank, ok = ranks["bronze"]  // ok=false since value is deleted
	fmt.Printf("After deletion, ranks[bronze] = %v, ok = %t\n", rank, ok)
}