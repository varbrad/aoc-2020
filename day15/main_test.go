package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay15Part1(t *testing.T) {
	t.Run(`solves the given example case`, func(t *testing.T) {
		expected := 436
		actual := Day15Part1([]int{0, 3, 6}, 2020)

		assert.Equal(t, expected, actual)
	})

	t.Run(`solves a small case`, func(t *testing.T) {
		input := []int{0, 3, 6}
		expected := 3
		actual := Day15Part1(input, 5)

		assert.Equal(t, expected, actual)
	})
}

func TestDay15Part2(t *testing.T) {
	expected := 175594
	actual := Day15Part1([]int{0, 3, 6}, 30000000)

	assert.Equal(t, expected, actual)
}
