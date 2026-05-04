/*
This program and src/gadeget package demonstrates the need for interfaces.
Interfaces define what methods should a type satisfy.
*/

package main

import "head-first-go/src/gadget"

// function to play all songs in a slice through type TapePlayer and its method Play() and Pause()
func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Pause()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Heera", "Let Down", "Bohemian Rhapsody"}
	playList(player, mixtape)
}