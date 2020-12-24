package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay10Part1(t *testing.T) {
	t.Run(`works on a small example`, func(t *testing.T) {
		numbers := []int{
			16,
			10,
			15,
			5,
			1,
			11,
			7,
			19,
			6,
			12,
			4,
		}

		expected := 35 // 7 x 5
		actual := Day10Part1(numbers)

		assert.Equal(t, expected, actual)
	})

	t.Run(`works on a larger example`, func(t *testing.T) {
		numbers := []int{
			28,
			33,
			18,
			42,
			31,
			14,
			46,
			20,
			48,
			47,
			24,
			23,
			49,
			45,
			19,
			38,
			39,
			11,
			1,
			32,
			25,
			35,
			8,
			17,
			7,
			9,
			4,
			2,
			34,
			10,
			3,
		}

		expected := 220 // 22 x 10
		actual := Day10Part1(numbers)

		assert.Equal(t, expected, actual)
	})
}

func TestDay10Part2(t *testing.T) {
	t.Run(`works on a small example`, func(t *testing.T) {
		numbers := []int{
			16,
			10,
			15,
			5,
			1,
			11,
			7,
			19,
			6,
			12,
			4,
		}

		expected := 8
		actual := Day10Part2(numbers)

		assert.Equal(t, expected, actual)
	})

	t.Run(`works on a larger example`, func(t *testing.T) {
		numbers := []int{
			28,
			33,
			18,
			42,
			31,
			14,
			46,
			20,
			48,
			47,
			24,
			23,
			49,
			45,
			19,
			38,
			39,
			11,
			1,
			32,
			25,
			35,
			8,
			17,
			7,
			9,
			4,
			2,
			34,
			10,
			3,
		}

		expected := 19208
		actual := Day10Part2(numbers)

		assert.Equal(t, expected, actual)
	})
}
