/*
This program demonstrates how variadic functions allow undefined number of parameters as a slice
*/

package main 

import (
	"fmt"
)

func serveralStrings(strings ...string) {
	fmt.Println(strings)
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func test(value ...int) {
	fmt.Printf("value = %v type = %T\n", value, value) // the type of variadic parameter is a slice
}

func main() {
	serveralStrings("a", "b")
	serveralStrings("a", "b", "c", "d")
	serveralStrings()

	mix(1, true, "a", "b")
	mix(2, false)

	test(1,2,3)
	test()
}