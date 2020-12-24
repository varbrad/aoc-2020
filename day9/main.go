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
	utils.Part2(Day9Part2(input, 25, 25))
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

// Day9Part2 solver
func Day9Part2(numbers numberList, preamble int, length int) int {
	target := Day9Part1(numbers, preamble, length)
	set := numbers.contiguousSet(target)
	min, max := set.minMax()
	return min + max
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

func (nl *numberList) sum() int {
	sum := 0
	for _, v := range *nl {
		sum += v
	}
	return sum
}

func (nl *numberList) contiguousSet(n int) numberList {
	length := len(*nl)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			list := (*nl)[i : j+1]
			sum := list.sum()
			if sum == n {
				return list
			} else if sum > n {
				break
			}
		}
	}
	return nil
}

func (nl *numberList) minMax() (int, int) {
	min, max := (*nl)[0], (*nl)[0]
	for _, v := range *nl {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}
