package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay12Part1(t *testing.T) {
	t.Run(`given test case`, func(t *testing.T) {
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
	})

	t.Run(`Improve test coverage`, func(t *testing.T) {
		input := []string{
			"F10",
			"S5",
			"E22",
			"W13",
			"L90",
			"L270",
			"L90",
			"R270",
			"R180",
			"L180",
			"F3",
			"F17",
			"R90",
			"F11",
			"S20",
			"R90",
			"F5",
			"R90",
			"F5",
			"R90",
			"F5",
			"R90",
			"F5",
		}

		expected := 15
		actual := Day12Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay12Part2(t *testing.T) {
	t.Run(`given test case`, func(t *testing.T) {
		input := []string{
			"F10",
			"N3",
			"F7",
			"R90",
			"F11",
		}

		expected := 286
		actual := Day12Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run(`Improve test coverage`, func(t *testing.T) {
		input := []string{
			"F10",
			"S5",
			"E22",
			"W13",
			"L90",
			"L270",
			"L90",
			"R270",
			"R180",
			"L180",
			"F3",
			"F17",
			"R90",
			"F11",
			"S20",
			"R90",
			"F5",
			"R90",
			"F5",
			"R90",
			"F5",
			"R90",
			"F5",
		}

		expected := 535
		actual := Day12Part2(input)

		assert.Equal(t, expected, actual)
	})
}
