/*
This program demonstrates how to tell zero values apart from assigned values in a map
using the "ok" boolean returned when a value is fetched using a key
*/

package main

import (
	"fmt"
)

func status(name string) {
	grades := map[string]float64{"John": 0, "Doe": 69.42}
	grade, ok := grades [name]
	if !ok {
		fmt.Printf("%s has not enrolled yet!\n", name)
		fmt.Printf("Bofean value ok for %s = %t\n", name, ok)
	} else if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
		fmt.Printf("Boolean value ok for %s = %t\n", name, ok)
	} else {
		fmt.Printf("%s has passed!\n", name)
		fmt.Printf("Boolean value ok for %s = %t\n", name, ok)
	}
}

func main() {
	status("John") // exists and fails
	status("Doe") // exists and passes
	status("Maria") // does not exist and hence boolean ok is false
}