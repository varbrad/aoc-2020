package utils

import (
	"strconv"
)

// ToIntegers converts a slice of strings to a slice of integers
func ToIntegers(stringList []string) []int {
	ints := make([]int, len(stringList))
	for i, s := range stringList {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}
