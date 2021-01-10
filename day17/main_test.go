package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay17Part1(t *testing.T) {
	t.Run(`should calculate the correct value after 1 cycle`, func(t *testing.T) {

		input := []string{
			".#.",
			"..#",
			"###",
		}

		expected := 11
		actual := Day17Part1(input, 1)

		// expected := 112
		// actual := Day17Part1(input, 6)

		assert.Equal(t, expected, actual)
	})

	t.Run(`should calculate the correct value after 6 cycles`, func(t *testing.T) {

		input := []string{
			".#.",
			"..#",
			"###",
		}

		expected := 112
		actual := Day17Part1(input, 6)

		assert.Equal(t, expected, actual)
	})
}
