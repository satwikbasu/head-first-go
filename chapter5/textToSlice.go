package main 

import (
	"head-first-go/src/datafile"
	"fmt"
	"log"
)

func main() {
	var values []float64 
	values, err := datafile.GetFloats("chapter5/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(values)
}