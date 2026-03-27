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

	votes := make(map[string]int)
	for _, line := range(lines) {
		votes[line]++
	}

	for key, value := range(votes) {
		fmt.Printf("Votes for %s : %d\n", key, value)
	}
}