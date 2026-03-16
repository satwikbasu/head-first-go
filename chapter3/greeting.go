/*
This code is actually present in /usr/share/go/src/greeting 
*/
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
