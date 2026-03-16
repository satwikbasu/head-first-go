package main

import "greeting"

func main() {
	greeting.Hello()
	greeting.Hi()
}

/*
This code below is actually present in /usr/share/go/src/greeting
which is being imported into this program packageTest.go

package greeting

import "fmt"

// method name "Hello" is capitalized so it can be exported
func Hello() {
	fmt.Println("Hello")
}

// method name "Hi" is capitalized so it can be exported
func Hi() {
	fmt.Println("Hi")
}

*/