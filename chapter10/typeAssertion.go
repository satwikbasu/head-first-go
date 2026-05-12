/*
This program demonstrates how type assertions allow you to extract the concrete type
from an interface value, enabling access to type-specific methods not defined in the interface.
*/

package main

import (
	"head-first-go/src/gadget"
	"fmt"
	"log"
)

type Player interface {
	Play(string)
	Pause()
}

func TryOut(player Player) {
	player.Play("Test track")
	player.Pause()
	/*
	player.Record() will throw a player.Record undefined error since the interface does not have it

	// recorder := gadget.TapeRecorder(player)
	// recorder.Record() 
	will also throw an error saying it needs type assertion

	Below is how type assertion solves this issue:
	Type assertion syntax: value.(ConcreteType)
	player = the interface value we're asserting from
	.(gadget.TapeRecorder) = the .() syntax performs type assertion, converting to the concrete type 
	*/
	recorder, ok := player.(gadget.TapeRecorder) // Now we can call Record() which is NOT part of the Player interface, only TapeRecorder
	// However we also need to assign ok as a second return value from type assertion whether it was successful
	// This is done becaue if type assertion is not successful, it panics at runtime
	if ok { 
		recorder.Record()
		fmt.Printf("Played through %T\n", recorder) // to verify the underlying type of the asserted concrete type
	} else {
		log.Fatal(ok)
	}
}

func main() {
	TryOut(gadget.TapeRecorder{}) // does not panic since TapeRecorder type matches the type assertion
	TryOut(gadget.TapePlayer{}) // fails since TapePlayer type does not match gadget.TapeRecorder assertion
}