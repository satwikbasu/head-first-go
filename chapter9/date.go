package main

import (
	"fmt"
	"log"
	"head-first-go/src/calender"
)

func main() {
	date := calender.Date{}
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