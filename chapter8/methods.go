/*
Methods are functions associated with values of a particular given type.

E.g. time package provides time.Now that returns a time.Time as a defined type
with a value representing current date and time.

Similary, we can define methods of our own to accept/return a particular type defined by us.
*/
package main

import "fmt"

type MyType string // define a type my method is going to accept as receiving parameter

func (m MyType) sayHi() { // sayHi only accepts values of type MyType to call the method with
	fmt.Println("Hi from", m) // usually first letter of type in lowercase is used as receiving parameter i.e. m for MyType
}

func (m MyType) methodWithParameters(number int, flag bool) { // unexported method
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

func (m MyType) WithReturn() int { // exported method
	return len(m)
}

func main() {
	value := MyType("John") // value of MyType type is created
	value.sayHi() // sayHi method is called with a MyType type, which it accepts
	anotherValue := MyType("Satwik")
	anotherValue.sayHi()

	value.methodWithParameters(69, false)
	fmt.Println(anotherValue.WithReturn())
}