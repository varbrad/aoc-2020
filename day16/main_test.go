package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay15Part1(t *testing.T) {
	t.Run(`solves a minimal case`, func(t *testing.T) {
		input := `
			test: 1-5 or 7-8

			your ticket:
			4,2

			nearby tickets:
			5
			9,5,2
			2,3,4,5
			12,4
		`

		expected := 21
		actual := Day16Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run(`solves the part 1 example`, func(t *testing.T) {
		input := `
			class: 1-3 or 5-7
			row: 6-11 or 33-44
			seat: 13-40 or 45-50

			your ticket:
			7,1,14

			nearby tickets:
			7,3,47
			40,4,50
			55,2,20
			38,6,12
		`

		expected := 71
		actual := Day16Part1(input)

		assert.Equal(t, expected, actual)
	})
}
