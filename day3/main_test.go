package main

import (
	"testing"
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

	result := Day3Part1(input)
	expected := 7

	if result != expected {
		t.Fatalf("The provided example case failed, expected %v, got %v", expected, result)
	}
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

	result := Day3Part2(input)
	expected := 336

	if result != expected {
		t.Fatalf("The provided example case failed, expected %v, got %v", expected, result)
	}
}
