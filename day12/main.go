package main

import (
	"log"
	"math"
	"regexp"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	input, err := utils.ReadInputToList("day12/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Part1(Day12Part1(input))
}

const (
	north = 0
	east  = 90
	south = 180
	west  = 270
)

var puzzleRegex = regexp.MustCompile(`^([A-Z])(\d+)$`)

// State of the puzzle
type State struct {
	dir int // 0 = north, 90 = east, 180 = south, 270 = west
	x   int
	y   int
}

// Day12Part1 solver
func Day12Part1(input []string) int {
	state := &State{dir: east}

	for _, ins := range input {
		state.process(ins)
	}

	return state.manhattan()
}

func (s *State) process(ins string) {
	res := puzzleRegex.FindStringSubmatch(ins)
	action := res[1]
	value, _ := utils.ToInteger(res[2])

	switch action {
	case "N":
		s.y += value
	case "S":
		s.y -= value
	case "E":
		s.x += value
	case "W":
		s.x -= value
	case "L":
		s.dir -= value
	case "R":
		s.dir += value
	case "F":
		dx, dy := s.dxdy()
		s.x += dx * value
		s.y += dy * value
	}
	for s.dir < 0 {
		s.dir += 360
	}
	s.dir = s.dir % 360
}

func (s *State) manhattan() int {
	return int(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}

func (s *State) dxdy() (dx int, dy int) {
	switch s.dir {
	case north:
		return 0, 1
	case east:
		return 1, 0
	case south:
		return 0, -1
	default:
		return -1, 0
	}
}
