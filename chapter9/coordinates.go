package main

import (
	"fmt"
	"log"
	"head-first-go/src/geo"
)

func main() {
	coordinates := geo.Coordinates{}
	err := coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}

	err = coordinates.SetLongitude(-22.28)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())
}