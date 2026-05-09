package main 

import "fmt"

type Appliance interface {
	TurnOn()
}

type Fan string
// TurnOn() cannot be removed since it satisfies interface
func (f Fan) TurnOn() {
	fmt.Println("Spinning")
}

type CoffeePot string

// TurnOn() cannot be removed since it satisfies interface
func (c CoffeePot) TurnOn() {
	fmt.Println("Powering up")
}

// Brew() can be removed since it does not satisfy interface
func (c CoffeePot) Brew() {
	fmt.Println("Heating up")
}

func main() {
	var device Appliance
	device = Fan("Havells")
	device.TurnOn()

	device = CoffeePot("Nescafe")
	device.TurnOn()
	// device.Brew() will not work since it calls a method 
	// from the concrete type that is not defined in the interface
}