package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3Part1(t *testing.T) {
	input := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	expected := 7
	actual := Day3Part1(input)

	assert.Equal(t, expected, actual)
}

func TestDay3Part2(t *testing.T) {
	input := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	expected := 336
	actual := Day3Part2(input)

	assert.Equal(t, expected, actual)
}
