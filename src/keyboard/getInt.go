// Package keyboard reads user input from the keyboard
package keyboard

import (
	"os"
	"bufio"
	"strconv"
	"strings"
)

// GetInt reads an integer number from the keyboard
// It returns the number read and any error encountered
func GetInt() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err:= reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	intValue, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0, err
	}

	return int(intValue), nil
}
