package main

import (
	"fmt"
)

func main() {

	//Most used methods for formatted string is Printf() and Sprintf()
	fmt.Println("About one-third:", 1.0/3.0)
	fmt.Printf("About one-third: %0.2f\n", 1.0/3.0) //upto 2 decimal places
	resultString := fmt.Sprintf("About one-third: %0.3f\n", 1.0/3.0) //same but results in formatted string 
	fmt.Println(resultString)

	//All formatting verbs
	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 15)
	fmt.Printf("A string: %s\n", "hello")
	fmt.Printf("A boolean: %t\n", false)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n\n")

	//Formatting value widths
	fmt.Printf("%2s | %10s\n", "Verb", "Output")
	fmt.Printf("---------------------------\n")
	fmt.Printf("%4s | %s", "%f", "Floating-point number\n")
	fmt.Printf("%4s | %s", "%d", "Decimal integer\n")
	fmt.Printf("%4s | %s", "%s", "String\n")
	fmt.Printf("%4s | %s", "%t", "Boolean\n")
	fmt.Printf("%4s | %s", "%v", "Any value(based on the supplied value's type)\n")
	fmt.Printf("%4s | %s", "%#v", "Any value(as it would appear in code)\n")
	fmt.Printf("%4s | %s", "%T", "Type of the supplied value (int, string etc.)\n")
	fmt.Printf("%4s | %s", "%", "A percent sign\n\n")

	//Formatting fractional number widths - numbers are rounded (value is lost)
	var n float64 = 12.3456 
	fmt.Printf("%%7.3f: %7.3f\n", n)
	fmt.Printf("%%7.2f: %7.2f\n", n)
	fmt.Printf("%%7.1f: %7.1f\n", n)
	fmt.Printf("%%.1f: %.1f\n", n)
	fmt.Printf("%%.2f: %.2f\n", n)

}