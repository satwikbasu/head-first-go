package main

import (
	"fmt"
	"log"
	"head-first-go/src/calender"
)

func main() {
	date := calender.Date{}
	// date := calender.Date.Year = 2016 ❌
	// cannot directly use the unexported struct fields from the package
	// can only use setter methods like date.SetYear(2026), not date.Year = 2026
	err := date.SetYear(2026)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(26)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}