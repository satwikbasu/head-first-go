package keyboard

import (
	"os"
	"bufio"
	"strconv"
	"strings"
)

// capitalized function name means it is exported and can be used in other packages
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err:= reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	floatValue, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return floatValue, nil
}
