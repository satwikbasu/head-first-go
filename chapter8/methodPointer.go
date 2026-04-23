package main

import "fmt"

type MyType string

// a method that accepts a MyType value
func (m MyType) method() { 
	fmt.Println("Method with", m)
}

// a method that accepts a pointer to a MyType value
func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer", m)
}

func main() {
	value := MyType("a value") // a value of type MyType
	pointer := &value // a pointer to a value of type MyType
	value.method() // value is passed to method that accepts value - should work, does work
	value.pointerMethod() // value is passed to method that accepts pointers - still works
	pointer.method() // pointer is passed to method that accepts values - still works
	pointer.pointerMethod() // pointer is passed to method that accepts pointers - should work, does work

	// What won't work: 
	// &MyType("example").pointerMethod
}