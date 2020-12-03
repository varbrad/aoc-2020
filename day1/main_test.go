package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Part1(t *testing.T) {
	values := []int{1721, 979, 366, 299, 675, 1456}
	expected := 514579
	actual := Day1Part1(values, 2020)

	assert.Equal(t, expected, actual)
}

func TestDay1Part2(t *testing.T) {
	values := []int{1721, 979, 366, 299, 675, 1456}
	expected := 241861950
	actual := Day1Part2(values, 2020)

	assert.Equal(t, expected, actual)
}
