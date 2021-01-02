package main

import (
	"log"
	"math"
	"strings"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	input, err := utils.ReadInputToList("day13/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Part1(Day13Part1(input))
}

// PuzzleInput structured puzzle input struct
type PuzzleInput struct {
	target int
	buses  []int
}

// Day13Part1 solver
func Day13Part1(input []string) int {
	return parseInput(input).solve1()
}

func Day13Part2(input []string) int {
	return parseInput(input).solve2()
}

func parseInput(input []string) *PuzzleInput {
	pi := &PuzzleInput{}
	target, _ := utils.ToInteger(input[0])
	pi.target = target
	for _, v := range strings.Split(input[1], ",") {
		num, err := utils.ToInteger(v)
		if err == nil {
			pi.buses = append(pi.buses, num)
		}
	}
	return pi
}

func (pi *PuzzleInput) solve1() int {
	minVal, minBus := -1.0, -1
	target := float64(pi.target)
	for _, v := range pi.buses {
		floatVal := float64(v)
		ceil := math.Ceil(target/floatVal) * floatVal
		if minVal == -1 || ceil < minVal {
			minVal = ceil
			minBus = v
		}
	}
	return int(minVal-target) * minBus
}

func (pi *PuzzleInput) solve2() int {
	return 0
}
