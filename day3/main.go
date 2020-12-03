package main

import (
	"fmt"
	"log"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	input, err := utils.ReadInputToList("day3/input")

	if err != nil {
		log.Panic(err)
		return
	}

	fmt.Printf("Part 1 = %v", Day3Part1(input))
	fmt.Println("")
}

// Day3Part1 solver
func Day3Part1(input []string) int {
	x, y := 3, 1
	trees := 0
	height := len(input)
	width := len(input[0])

	for {
		if y >= height {
			break
		}

		isTree := string(input[y][x]) == "#"

		if isTree {
			trees++
		}

		x = (x + 3) % width
		y++
	}

	return trees
}
