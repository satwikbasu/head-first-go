package main 

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"math/rand"
	"time"
	"strings"
	"strconv"
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
	// fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)
	success := false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		//string.TrimSpace() is used to remove preceding and suceeding spaces
		input = strings.TrimSpace(input) 
		/*
		strconv.Atoi() is used to string representation of a number into an integer value
		It is a convenience function that assumes a decimal (base-10) integer 
		and is equivalent to strconv.ParseInt(input, base, bitsize) 
		*/
		guess, err := strconv.Atoi(input) //equivalent to strconv.ParseInt(10, 0, 0)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			fmt.Println("You guessed it!")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess it correctly. It was", target)
	}
}