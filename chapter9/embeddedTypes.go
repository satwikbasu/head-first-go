/*
This program and src/embeddedTypes.go from calender package demonstrate how embeddedTypes
can call methods of the underlying outer type as long as they are exported
*/

package main

import (
	"head-first-go/src/calender"
)

func main() {
	value := calender.MyType{}
	value.ExportedMethod()
	// value.unexportedMethod() // unexported method cannot be called
}