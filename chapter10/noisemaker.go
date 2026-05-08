package main 

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type Robot string // new Robot type declared

// Robot type satisfies NoiseMaker interface with MakeSound()
func (r Robot) MakeSound() {
	fmt.Println("Beep Boop!")
}

// An additional method (concrete method)
func (r Robot) Walk() {
	fmt.Println("Powering legs...")
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound() // OK, part of NoiseMaker interface
	n.Walk()  // NOT OK, not part of NoiseMaker interface
}

func main() {
	play(Robot("Bot")) // will throw an error until n.Walk() is commented in play()
}

