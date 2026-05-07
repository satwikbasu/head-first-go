/*
This program and chapter10/interface.go demonstrate how interfaces are used.
Interfaces define what methods should a type satisfy.
*/

package interpkg

import "fmt"

// interface type named "MyInterface" with 3 methods to satisfy
type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
	// NotAMethod() int - cannot have standalone functions, interface accepts only methods
}

type MyType int // type "MyType" to create methods (without types, methods are functions)

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called") // 1st method to satisfy for MyInterface
}

func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called with", f) // 2nd method to satisfy for MyInterface
}

func (m MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue" // 3rd method to satisfy for MyInterface
}

func (m MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called") // concrete method not required to satisfy MyInterface
}

// REMINDER: Go interfaces can ONLY be satisfied by METHODS (functions with receivers on a type)
// NOT by standalone functions. If an interface requires NotAMethod(), the type must implement
// it as a method: func (m MyType) NotAMethod() int { ... }
// func NotAMethod() int {
// 	return 69
// }