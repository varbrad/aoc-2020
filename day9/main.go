package main

import (
	"log"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	input, err := utils.ReadInputToIntegerList("day9/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Part1(Day9Part1(input, 25, 25))
}

// Day9Part1 solver
func Day9Part1(numbers numberList, preamble int, length int) int {
	ix := preamble

	for {
		if ix >= len(numbers) {
			break
		}
		val := numbers[ix]
		nl := numbers[ix-length : ix]

		if !nl.hasAdditivePair(val) {
			return val
		}

		ix++
	}

	return -1
}

type numberList []int

func (nl *numberList) hasAdditivePair(value int) bool {
	for ix, v := range *nl {
		need := value - v
		if nl.contains(need, ix) {
			return true
		}
	}
	return false
}

func (nl *numberList) contains(n int, excludedIndex int) bool {
	for ix, v := range *nl {
		if ix == excludedIndex {
			continue
		}
		if v == n {
			return true
		}
	}
	return false
}
