/*
This program is to demonstrate robust error/result handling
*/

package main 

import (
	"fmt"
	"errors"
)

// divide() only returns result & errors if they occur while dividing
func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}
	return dividend / divisor, nil
}

// errorHandle() is tasked with printing both result and errors
func errorHandle(quotient float64, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.3f\n", quotient)
	}
}

// main() simply passes values to divide() and then errorHandle()
func main() {
	quotient, err := divide(5.6, 2.0)
	errorHandle(quotient, err)
	quotient, err = divide(5.6, 0.0)
	errorHandle(quotient, err)
}