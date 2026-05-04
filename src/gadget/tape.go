/*
This program and chapter10/tape.go demonstrates the need for interfaces.
Interfaces define what methods should a type satisfy.
*/

package gadget

import "fmt"

// type TapePlayer needs Play() and Pause() methods
type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Pause() {
	fmt.Println("Paused!")
}

// type TapeRecorder needs Play(), Pause() and Record() methods
type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording!")
}

func (t TapeRecorder) Pause() {
	fmt.Println("Paused!")
}