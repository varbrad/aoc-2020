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
	utils.Part2(Day12Part2(input))
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

// Pos vec2 type
type Pos struct {
	x int
	y int
}

// State2 format for part 2
type State2 struct {
	ship     *Pos
	waypoint *Pos
}

// Day12Part1 solver
func Day12Part1(input []string) int {
	state := &State{dir: east}

	for _, ins := range input {
		state.process(ins)
	}

	return state.manhattan()
}

// Day12Part2 solver
func Day12Part2(input []string) int {
	state := &State2{
		ship:     &Pos{},
		waypoint: &Pos{x: 10, y: 1},
	}

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

func (s *State2) process(ins string) {
	res := puzzleRegex.FindStringSubmatch(ins)
	action := res[1]
	value, _ := utils.ToInteger(res[2])

	switch action {
	case "N":
		s.waypoint.y += value
	case "S":
		s.waypoint.y -= value
	case "E":
		s.waypoint.x += value
	case "W":
		s.waypoint.x -= value
	case "L":
		s.rotateRight(360 - value)
	case "R":
		s.rotateRight(value)
	case "F":
		s.ship.x += s.waypoint.x * value
		s.ship.y += s.waypoint.y * value
	}
}

func (s *State2) manhattan() int {
	x, y := s.ship.x, s.ship.y
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func (s *State2) rotateRight(deg int) {
	if deg == 0 {
		return
	}
	s.waypoint.x, s.waypoint.y = s.waypoint.y, -s.waypoint.x
	s.rotateRight(deg - 90)
}
