/*
This program and chapter9/embeddedTypes.go demonstrate how embeddedTypes
can call methods of the underlying outer type as long as they are exported
*/

package calender 

import "fmt"

type MyType struct {
	EmbeddedType
}

type EmbeddedType string

func (e EmbeddedType) ExportedMethod() {
	fmt.Println("Hi from ExportedMethod on EmbeddedType")
}

func (e EmbeddedType) unexportedMethod() {

}
