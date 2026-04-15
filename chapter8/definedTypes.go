package main 

import "fmt"

type Gallons float64
type Liters float64

func main() {
	var carFuel Gallons // will not work with float64 even if underlying type matches
	carFuel = Gallons(3.16) // carFuel in gallons
	busFuel := Liters(13.3) // busFuel in liters
	fmt.Printf("Value of carFuel = %v gallons and type = %T\n", carFuel, carFuel)
	fmt.Printf("Value of busFuel = %v liters and type = %T\n", busFuel, busFuel)

	var totalFuelGallons Gallons
	var totalFuelLiters Liters
	totalFuelLiters= busFuel + Liters(carFuel * 0.264) // Gallons type carFuel is converted to Liters
	totalFuelGallons = Gallons(busFuel * 3.785) + carFuel // Liters type busFuel is converted to Gallons

	// Arithmetic operations between two similar types is possible
	// Gallons + Gallons is possible but Liters * Gallons is not possible
	// even if underlying type is same, operations between two separately defines types is not possible 
	fmt.Printf("Total fuel in gallons = %.5f and type = %T\n", totalFuelGallons, totalFuelGallons)
	fmt.Printf("Total fuel in liters = %.5f and type = %T\n", totalFuelLiters, totalFuelLiters)
}