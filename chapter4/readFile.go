package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("sample.txt")
	// the path to the file depends on the directory from where the program is ran
	// if the program is run from ~/head-first-go, the path should be chapter4/sample.txt
	if err != nil {
		// if an error was found in opening the file
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	// loops until the end of the file is reached and scanner.Scan returns false
	for scanner.Scan() {
		fmt.Printf("%#v ",scanner.Text())
		fmt.Printf("%T\n",scanner.Text())
	}

	err = file.Close()
	if err != nil {
		// if an error was found in closing the file
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		// if there was any error scanning the file
		log.Fatal(scanner.Err())
	}
}