package main 

import (
	"fmt"
	"head-first-go/src/calculator" // the package name to import depends on the directory under /src
)

func main() {
	fmt.Println(calc.Add(1, 2))
	fmt.Println(calc.Subtract(7, 3))
	fmt.Println(calc.Multiply(2, 3))
	fmt.Println(calc.Divide(6, 3))
}