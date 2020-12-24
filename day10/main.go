package main

import (
	"log"
	"sort"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	input, err := utils.ReadInputToIntegerList("day10/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Part1(Day10Part1(input))
}

// Day10Part1 solver
func Day10Part1(numbers []int) int {
	numbers = append(numbers, 0)
	sort.Ints(numbers)

	// 3 always starts at 1 for the final max jump
	joltJumps := map[int]int{1: 0, 2: 0, 3: 1}

	for i := 0; i < len(numbers)-1; i++ {
		a, b := numbers[i], numbers[i+1]
		jmp := b - a
		joltJumps[jmp]++
	}

	return joltJumps[1] * joltJumps[3]
}
