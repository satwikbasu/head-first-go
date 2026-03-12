package main 

import "fmt"
 
func main() {
	amount := 6
	fmt.Println("Value stored in amount =",amount)
	fmt.Println("Address of amount =",&amount)

	fmt.Printf("%T is type of amount\n",amount)
	fmt.Printf("%T is type of pointer for amount\n\n",&amount)
	
	var myInt int = 69
	var myIntPointer *int
	myIntPointer = &myInt
	myValueFromPointer := *myIntPointer
	fmt.Printf("Value of myInt = %d\n",myInt)
	fmt.Printf("Value of myIntPointer = &myInt = %v\n",myIntPointer)
	fmt.Printf("Value of myValueFromPointer = *myIntPointer = %v\n\n",myValueFromPointer)
	fmt.Printf("Type of myInt = %T\n",myInt)
	fmt.Printf("Type of myIntPointer = %T\n",myIntPointer)
	fmt.Printf("Type of myValueFromPointer = %T\n",myValueFromPointer)

}