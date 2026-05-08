/*
This program demonstrates implicit interface satisfaction and duck typing in Go.

Key Concepts:
- Interfaces define a CONTRACT: a set of methods that a type must implement to satisfy that interface.
- Satisfaction is IMPLICIT: there's no explicit "implements" keyword in Go. Any type with matching method signatures automatically satisfies the interface.
- Interfaces are BEHAVIORAL contracts: they specify WHAT a type can DO (its methods), not WHAT it IS (its underlying type or storage).
- A type satisfies an interface if and only if it implements ALL methods defined by that interface.
- Multiple distinct types can independently satisfy the same interface (polymorphism).
- A variable of interface type can hold any concrete type that satisfies that interface.
- Interface values store BOTH the concrete value AND its type information.
*/

package main

import "fmt"

// Whistle is a defined type with string as its underlying type.
// It will implicitly satisfy the NoiseMaker interface because it has a MakeSound() method.
type Whistle string

// MakeSound() is a method receiver on the Whistle type.
// This method satisfies the definition of the NoiseMaker interface (which requires MakeSound()).
func (w Whistle) MakeSound() {
	fmt.Println("Tweet")
}

// Horn is a defined type with string as its underlying type.
// It will implicitly satisfy the NoiseMaker interface because it has a MakeSound() method.
type Horn string

// MakeSound() is a method receiver on the Horn type.
// This method also satisfies the NoiseMaker interface. Different types can satisfy the same interface.
func (h Horn) MakeSound() {
	fmt.Println("Honk")
}

// Water is a defined type with string as its underlying type.
// It deliberately DOES NOT satisfy the NoiseMaker interface because it lacks the MakeSound() method.
// This demonstrates that interface satisfaction depends on having ALL required methods.
type Water string 

// MakeSplash() is a method on Water type, but it does NOT match the NoiseMaker interface.
// NoiseMaker requires MakeSound(), so Water cannot be assigned to a NoiseMaker variable.
// However, Water DOES satisfy the SplashMaker interface.
func (w Water) MakeSplash() {
	fmt.Println("Splash!")
}

// NoiseMaker is an interface that defines a contract requiring one method: MakeSound().
// Any type that implements a MakeSound() method (with matching signature) automatically satisfies NoiseMaker.
// The interface does not specify HOW to MakeSound, only that the capability must exist.
type NoiseMaker interface {
	MakeSound()
}

// SplashMaker is an interface that defines a contract requiring one method: MakeSplash().
// Any type that implements a MakeSplash() method automatically satisfies SplashMaker.
type SplashMaker interface {
	MakeSplash()
}

// example function to prove interface types can be declared in function parameters as well
func play(n NoiseMaker){
	n.MakeSound()
}

func main() {
	var toy NoiseMaker
	// toy is a variable of interface type NoiseMaker. It can only hold values of types that satisfy NoiseMaker.
	toy = Whistle("Canary") // This assignment is valid because Whistle has a MakeSound() method
	fmt.Printf("Value of toy = %v\n", toy) // prints the concrete value: Canary
	fmt.Printf("Type of toy = %T\n", toy) // prints the concrete type (Whistle), not the interface type
	toy.MakeSound()
	toy = Horn("Truck")
	fmt.Printf("Value of toy = %v\n", toy)
	fmt.Printf("Type of toy = %T\n", toy)
	toy.MakeSound()
	play(toy) // same as toy.MakeSound()
	
	var liquid SplashMaker
	liquid = Water("H20")
	fmt.Printf("Value of liquid = %v\n", liquid)
	fmt.Printf("Type of liquid = %T\n", liquid)
	liquid.MakeSplash()
}