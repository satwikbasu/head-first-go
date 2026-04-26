package main

import "fmt"

type Date struct {
	Year int
	Month int
	Day int
}

func main() {
	date := Date{Day: 27, Month: 4, Year: 2026}
	fmt.Println(date)
}