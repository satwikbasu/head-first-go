package main

import (
	"head-first-go/src/greeting/hindi" //sub-package 1
	"head-first-go/src/greeting/bengali" //sub-package 2
	"head-first-go/src/greeting" //parent package
)

func main() {
	hindi.Hello() 
	bengali.Hello()
	greeting.Hi()
	greeting.Hello()
}