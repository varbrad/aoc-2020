package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Part1(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	expected := 2
	actual := Day2Part1(input)

	assert.Equal(t, expected, actual)
}

func TestDay2Part2(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	expected := 1
	actual := Day2Part2(input)

	assert.Equal(t, expected, actual)
}
