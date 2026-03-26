package main

import(
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	fmt.Println(os.Args) // prints the whole slice with arguments and filename
	fmt.Printf("slice=%v type=%T\n", arguments[1:], arguments) // 0th index element is omitted
}