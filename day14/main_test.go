package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay14Part1(t *testing.T) {
	input := []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}

	expected := 165
	actual := Day14Part1(input)

	assert.Equal(t, expected, actual)
}
