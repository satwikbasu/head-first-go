package main 

import(
	"fmt"
	"head-first-go/src/datafile"
	"log"
	"sort"
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
	
	var names []string
	for name := range(votes) {
		names = append(names, name)
	}
	sort.Strings(names) // sorting slice to maintain order since for...range handles maps in random order

	for _, name := range names {
		fmt.Printf("Votes for %s : %d\n", name, votes[name])
	}
}