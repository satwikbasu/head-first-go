package main

import "fmt"

type MyType string

func (m MyType) method() {
	fmt.Println("Method with", m)
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer", m)
}

func main() {
	value := MyType("a value")
	pointer := &value
	value.method()
	value.pointerMethod()
	pointer.method()
	pointer.pointerMethod()
}