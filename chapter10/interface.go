/*
This program and package interpkg demonstrate how interfaces are used.
Interfaces define what methods should a type satisfy.
*/
package main

import (
	"fmt"
	"head-first-go/src/interpkg"
)

func main() {
	// value is of type MyInterface, which has an underlying type interface
	var value interpkg.MyInterface
	value = interpkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(3.69)
	fmt.Println(value.MethodWithReturnValue())

	// value.MethodNotInInterface() will not work since value is defined as type MyInterface
	var example interpkg.MyType // value of type MyType can be used to call methods not defined under interface
	example.MethodNotInInterface()
	// REMINDER: Go interfaces require METHODS on types, not standalone functions.
	// Commented out: fmt.Println(interpkg.NotAMethod())
}