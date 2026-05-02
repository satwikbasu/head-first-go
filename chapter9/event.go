package main 

import (
	"head-first-go/src/calender"
	"log"
	"fmt"
)

func main() {
	event := calender.Event{}
	event.Title = "title of the event"
	err := event.SetDay(2)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2026)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event)
}