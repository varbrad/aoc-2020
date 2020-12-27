package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay12Part1(t *testing.T) {
	input := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}

	expected := 25
	actual := Day12Part1(input)

	assert.Equal(t, expected, actual)
}
