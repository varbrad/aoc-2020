package main

import (
	"fmt"
	"log"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	values, err := utils.ReadInputToIntegerList("day1/input")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Part 1 = %v", Day1Part1(values, 2020))
	fmt.Println("")
	fmt.Printf("Part 2 = %v", Day1Part2(values, 2020))
	fmt.Println("")
}

// Day1Part1 solver
func Day1Part1(numbers []int, target int) int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			a, b := numbers[i], numbers[j]
			if sum := a + b; sum == target {
				return a * b
			}
		}
	}
	return -1
}

// Day1Part2 solver
func Day1Part2(numbers []int, target int) int {
	for i := 0; i < len(numbers)-2; i++ {
		for j := i + 1; j < len(numbers)-1; j++ {
			for k := j + 1; k < len(numbers); k++ {
				a, b, c := numbers[i], numbers[j], numbers[k]
				if sum := a + b + c; sum == target {
					return a * b * c
				}
			}
		}
	}
	return -1
}
