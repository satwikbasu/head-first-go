package main

import (
	"head-first-go/src/calender"
)

func main() {
	value := calender.MyType{}
	value.ExportedMethod()
	// value.unexportedMethod() // ignored since unexported method
}