/*This program demonstrates importing constants from packages*/
package main 

import (
	"fmt"
	"head-first-go/src/dates"
)

func main() {
	days := 3
	fmt.Println("Your appointment is in", days, "days")
	fmt.Println("with a follow-up in", days + dates.DaysInWeek, "days")
}