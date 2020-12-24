package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay9Part1(t *testing.T) {
	numbers := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}

	expected := 127
	actual := Day9Part1(numbers, 5, 5)

	assert.Equal(t, expected, actual)
}
