package utils

import (
	"strconv"
)

// ToInteger converts a string to an integer
func ToInteger(text string) (int, error) {
	return strconv.Atoi(text)
}

// ToIntegers converts a slice of strings to a slice of integers
func ToIntegers(stringList []string) ([]int, error) {
	ints := make([]int, len(stringList))
	for i, s := range stringList {
		value, err := ToInteger(s)
		if err != nil {
			return nil, err
		}
		ints[i] = value
	}
	return ints, nil
}
