package main

import "head-first-go/src/gadget"

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