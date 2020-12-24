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
	utils.Part2(Day10Part2(input))
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

// Day10Part2 solver
func Day10Part2(numbers []int) int {
	numbers = append(numbers, 0)
	sort.Ints(numbers)

	subGraphs := breakGraph(numbers)
	prod := 1
	for _, v := range subGraphs {
		prod *= subSolver(v)
	}
	return prod
}

func breakGraph(numbers []int) [][]int {
	length := len(numbers)
	graphs := [][]int{}

	ix := 0
	graph := []int{}

	for {
		if ix == length {
			graphs = append(graphs, graph)
			break
		}
		val := numbers[ix]
		if len(graph) == 0 || val-3 < graph[len(graph)-1] {
			graph = append(graph, val)
		} else {
			graphs = append(graphs, graph)
			graph = []int{val}
		}
		ix++
	}

	return graphs
}

func subSolver(numbers []int) int {
	if len(numbers) < 3 {
		return 1
	}

	sort.Ints(numbers)

	start := []int{0}
	totals := map[int]int{}

	for {
		if len(start) == 0 {
			break
		}
		// Get first
		i := start[0]
		start = start[1:]

		val := numbers[i]
		// How many branches can we make?
		for dx := 1; dx < 4; dx++ {
			ix := i + dx
			if ix >= len(numbers) {
				break
			}
			val2 := numbers[ix]
			diff := val2 - val
			if diff > 3 {
				break
			}
			start = append(start, ix)
			totals[ix]++
		}
	}
	return totals[len(numbers)-1]
}
