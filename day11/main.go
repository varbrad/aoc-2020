package main

import (
	"log"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	input, err := utils.ReadInputToList("day11/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Part1(Day11Part1(input))
}

// Day11Part1 solver
func Day11Part1(rawLayout []string) int {
	layout := parseLayout(rawLayout)

	for {
		changed, nextState := layout.generate()
		layout = nextState
		if !changed {
			break
		}
	}

	return layout.getOccupiedSeats()
}

type seatState int

const (
	floor seatState = iota
	empty
	occupied
)

type seatLayout struct {
	grid [][]seatState
	w    int
	h    int
}

func emptyLayout(w int, h int) *seatLayout {
	sl := &seatLayout{w: w, h: h}
	sl.grid = make([][]seatState, h)
	for y := 0; y < h; y++ {
		sl.grid[y] = make([]seatState, w)
	}
	return sl
}

func parseLayout(layout []string) *seatLayout {
	l := emptyLayout(len(layout[0]), len(layout))

	for y, row := range layout {
		for x, seat := range row {
			l.grid[y][x] = toSeatState(string(seat))
		}
	}

	return l
}

func toSeatState(char string) seatState {
	switch char {
	case "L":
		return empty
	case ".":
		return floor
	case "#":
		return occupied
	}
	return occupied
}

func (sl *seatLayout) generate() (changed bool, layout *seatLayout) {
	changed = false
	layout = emptyLayout(sl.w, sl.h)

	for y, row := range layout.grid {
		for x := range row {
			current := sl.getSeatState(x, y)
			next := sl.getNextSeatState(x, y)
			layout.setSeat(x, y, next)
			if current != next {
				changed = true
			}
		}
	}

	return
}

func (sl *seatLayout) getNextSeatState(x int, y int) seatState {
	currentState := sl.getSeatState(x, y)
	neighbours := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			state := sl.getSeatState(x+dx, y+dy)
			if state == occupied {
				neighbours++
			}
		}
	}
	if currentState == empty && neighbours == 0 {
		return occupied
	} else if currentState == occupied && neighbours >= 4 {
		return empty
	}
	return currentState
}

func (sl *seatLayout) getSeatState(x int, y int) seatState {
	if x < 0 || x > sl.w-1 || y < 0 || y > sl.h-1 {
		return empty
	}
	return sl.grid[y][x]
}

func (sl *seatLayout) getOccupiedSeats() int {
	count := 0
	for _, row := range sl.grid {
		for _, seat := range row {
			if seat == occupied {
				count++
			}
		}
	}
	return count
}

func (sl *seatLayout) setSeat(x int, y int, state seatState) {
	sl.grid[y][x] = state
}
