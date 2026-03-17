//This is a main.go file for the greeting package
//This file is in greeting/cmd/greet so that the executable is named "greet"
//The command to create the executable is `go install head-first-go/src/greeting/cmd/greet`
//The executable can be found in ~/go/bin/greet

package main

import (
	"head-first-go/src/greeting"
	"head-first-go/src/greeting/hindi"
	"head-first-go/src/greeting/bengali"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	hindi.Hello()
	bengali.Hello()
}