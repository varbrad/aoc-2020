package main

import (
	"log"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	input, err := utils.ReadInputToList("day17/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Part1(Day17Part1(input, 6))
	utils.Part2(Day17Part2(input, 6))
}

// Day17Part1 solver
func Day17Part1(rawInput []string, cycles int) int {
	input := parseInput3d(rawInput)
	for i := 0; i < cycles; i++ {
		input = input.cycle()
	}
	return input.totalAlive()
}

// Day17Part2 solver
func Day17Part2(rawInput []string, cycles int) int {
	input := parseInput4d(rawInput)
	for i := 0; i < cycles; i++ {
		input = input.cycle()
	}
	return input.totalAlive()
}

func nextCellState(alive bool, neighbours int) bool {
	if alive && neighbours == 2 {
		return true
	}
	return neighbours == 3
}
