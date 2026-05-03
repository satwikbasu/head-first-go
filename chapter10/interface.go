package main

import (
	"fmt"
	"head-first-go/src/interpkg"
)

func main() {
	var value interpkg.MyInterface
	value = interpkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(3.69)
	fmt.Println(value.MethodWithReturnValue())

	// value.MethodNotInInterface() will not work since value is defined as type MyInterface
	var example interpkg.MyType // value of type MyType can be used to call methods not in interface
	example.MethodNotInInterface()
}