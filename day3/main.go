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
	fmt.Printf("Part 2 = %v", Day3Part2(input))
	fmt.Println("")
}

// Day3Part1 solver
func Day3Part1(input []string) int {
	treemap := treeMap{
		grid: input,
		dx:   3,
		dy:   1,
	}
	return treemap.solve()
}

// Day3Part2 solver
func Day3Part2(input []string) int {
	pairs := [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	mul := 1
	for _, dxdy := range pairs {
		treemap := treeMap{input, dxdy[0], dxdy[1]}
		mul *= treemap.solve()
	}
	return mul
}

type treeMap struct {
	grid []string
	dx   int
	dy   int
}

func (treemap treeMap) solve() int {
	x, y := treemap.dx, treemap.dy
	trees := 0
	height := len(treemap.grid)
	width := len(treemap.grid[0])

	for {
		if y >= height {
			break
		}

		isTree := string(treemap.grid[y][x]) == "#"

		if isTree {
			trees++
		}

		x = (x + treemap.dx) % width
		y += treemap.dy
	}

	return trees
}
