/*
This program and src/gadeget package demonstrates the need for interfaces.
Interfaces define what methods should a type satisfy.
*/

package main

import "head-first-go/src/gadget"

type Player interface {
	Play(string)
	Pause()
}

// ####### older implementation without interface ######
// function to play all songs in a slice through type TapePlayer and its method Play() and Pause()
// func playList(device gadget.TapePlayer, songs []string) {
// 	for _, song := range songs {
// 		device.Play(song)
// 	}
// 	device.Pause()
// }

// now playList() accepts any Player, not just a TapePlayer
func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Pause()
}

func main() {
	mixtape := []string{"Heera", "Let Down", "Bohemian Rhapsody"}

	// cannot use dynamic assignment since we need to mention the type explicitly as interface Player
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)

	player = gadget.TapeRecorder{}
	playList(player, mixtape)
}