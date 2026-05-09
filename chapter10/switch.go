package main

import "fmt"

type Switch string

// The toggle method on the Switch type has to use a pointer receiver
// so it can modify the receiver
func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type Toggleable interface {
	toggle()
}

func main() {
	s := Switch("off")
	// `var t Toggleable = s` returns an error Switch does not implement Toggleable
	var t Toggleable = &s
	t.toggle()
	t.toggle()
}