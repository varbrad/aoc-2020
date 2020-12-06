package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestDay6Part1(t *testing.T) {
	input := `
		abc

		a
		b
		c

		ab
		ac

		a
		a
		a
		a

		b
	`

	expected := 11
	actual := Day6Part1(input)

	assert.Equal(t, expected, actual)
}

func TestDay6Part2(t *testing.T) {
	input := `
		abc

		a
		b
		c

		ab
		ac

		a
		a
		a
		a

		b
	`

	expected := 6
	actual := Day6Part2(input)

	assert.Equal(t, expected, actual)
}
