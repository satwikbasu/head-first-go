/*
This program demonstrates how type assertions allow you to extract the concrete type
from an interface value, enabling access to type-specific methods not defined in the interface.
*/

package main

import (
	"head-first-go/src/gadget"
	"fmt"
)

type Player interface {
	Play(string)
	Pause()
}

func TryOut(player Player) {
	player.Play("Test track")
	player.Pause()
	/*
	// player.Record() 
	will throw a player.Record undefined error since the interface does not have it

	// recorder := gadget.TapeRecorder(player)
	// recorder.Record() 
	will also throw an error saying it needs type assertion
	*/

	// Below is how type assertion solves this issue:
	// Type assertion syntax: value.(ConcreteType)
	// player = the interface value we're asserting from
	// .(gadget.TapeRecorder) = the .() syntax performs type assertion, converting to the concrete type
	// recorder := ... = assigns the extracted TapeRecorder value to recorder variable
	recorder := player.(gadget.TapeRecorder)
	// Now we can call Record() which is NOT part of the Player interface, only TapeRecorder
	recorder.Record()
	fmt.Printf("%T\n", recorder) // to verify the underlying type of the asserted concrete type
}

func main() {
	TryOut(gadget.TapeRecorder{})
}