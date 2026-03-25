package datafile 

import (
	"os"
	"bufio"
	"strconv"
)

func GetFloats(filename string) ([]float64, error) {
	var values []float64
	file, err := os.Open(filename)
	if err != nil {
		return values, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		extractedValue, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return values, err
		}
		values = append(values, extractedValue)
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