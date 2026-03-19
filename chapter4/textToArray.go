package main 

import (
	"fmt"
	"head-first-go/src/datafile"
	"log"
)

func main() {
	var data [3]float64
	data, err := datafile.GetFloats("chapter4/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n",data)
	fmt.Printf("%T\n",data)
}