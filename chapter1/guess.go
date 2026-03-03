package main 

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"math/rand"
	"time"
)

func main() {
	/*
	seconds variable is used as a seed to generate random numbers
	time.Now() gives current date and time convert current date and time to integer
	time.Now().Unix() converts it into integer (seconds since Jan 1 1970)

	Uncomment next 4 lines to understand:

	timeDate := time.Now()
	fmt.Println("time.Now() output =",timeDate)
	secondsInt := time.Now().Unix()
	fmt.Println("time.Now.()Unix() output =",secondsInt)
	*/
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chose a random number between 1 & 100.")
	fmt.Print("Can you guess it? ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
	fmt.Println(target)
}