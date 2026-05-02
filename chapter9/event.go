/*
This program is a main.go to implement event type from package calender.
Event type uses embedded type Date in itself as a field of a struct.
Even though Date itself has fields that are unexported, event type can be used
to call setter and getter methods (which are exported) through event type.
*/

package main 

import (
	"head-first-go/src/calender"
	"log"
	"fmt"
)

func main() {
	event := calender.Event{}
	event.Title = "title of the event" // exported field of event type
	err := event.SetDay(2) // can only call exported setter methods of Date through Event
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
	fmt.Println(event.Day()) // can only call exported getter methods of Date through Event
	fmt.Println(event.Month())
	fmt.Println(event.Year())
}