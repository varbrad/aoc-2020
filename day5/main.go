package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	input, err := utils.ReadInputToList("day5/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Part1(Day5Part1(input))
	utils.Part2(Day5Part2(input))
}

// Day5Part1 solver
func Day5Part1(input []string) int {
	max := -1

	for _, s := range input {
		id := GetSeatID(s)
		if id > max {
			max = id
		}
	}

	return max
}

// Day5Part2 solver
func Day5Part2(input []string) int {
	max := Day5Part1(input)
	m := make(map[int]bool)

	for _, s := range input {
		m[GetSeatID(s)] = true
	}

	for i := 0; i < max; i++ {
		if m[i] {
			continue
		}
		if m[i-1] && m[i+1] {
			return i
		}
	}

	return -1
}

// GetSeatID returns the seat id of a given input string
func GetSeatID(input string) int {
	inputRegex := regexp.MustCompile(`([FB]{7})([RL]{3})`)
	result := inputRegex.FindStringSubmatch(input)

	column := getValue(result[1], "F", "B")
	row := getValue(result[2], "L", "R")

	return column*8 + row
}

func getValue(input string, low string, high string) int {
	binary := strings.ReplaceAll(strings.ReplaceAll(input, low, "0"), high, "1")
	val, _ := strconv.ParseInt(binary, 2, 0)
	return int(val)
}
