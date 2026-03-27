package main 

import(
	"fmt"
	"head-first-go/src/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("chapter6/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}