package main

import "fmt"

type Liters float64
type Gallons float64

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	// Before adding fuel
	fmt.Printf("Car fuel: %.1f gallons and type = %T\n", carFuel, carFuel)
	fmt.Printf("Bus fuel: %.1f liters and type = %T\n", busFuel, busFuel)
	carFuel += ToGallons(Liters(40.0))
	busFuel += ToLiters(Gallons(30.0))
	// After adding fuel
	fmt.Printf("Car fuel: %.1f gallons and type = %T\n", carFuel, carFuel)
	fmt.Printf("Bus fuel: %.1f liters and type = %T\n", busFuel, busFuel)
}