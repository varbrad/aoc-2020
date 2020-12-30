package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay13Part1(t *testing.T) {
	input := []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}

	expected := 295
	actual := Day13Part1(input)

	assert.Equal(t, expected, actual)
}
