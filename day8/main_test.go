package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay8Part1(t *testing.T) {
	input := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	expected := 5
	actual := Day8Part1(input)

	assert.Equal(t, expected, actual)
}

func TestDay8Part2(t *testing.T) {
	input := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	expected := 8
	actual := Day8Part2(input)

	assert.Equal(t, expected, actual)
}
