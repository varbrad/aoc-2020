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

func TestDay13Part2(t *testing.T) {
	t.Run("example test case", func(t *testing.T) {
		input := []string{
			"939",
			"7,13,x,x,59,x,31,19",
		}

		expected := 1068781
		actual := Day13Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("simple test case", func(t *testing.T) {
		input := []string{
			"0",
			"17,x,13,19",
		}

		expected := 3417
		actual := Day13Part2(input)

		assert.Equal(t, expected, actual)
	})
}
