package datafile 

import (
	"os"
	"bufio"
	"strconv"
)

func GetFloats(filename string) ([3]float64, error) {
	var values [3]float64
	file, err := os.Open(filename)
	if err != nil {
		return values, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return values, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return values, err
	}
	if scanner.Err() != nil {
		return values, err
	}
	return values, nil
}