package main

import (
	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	values := []int{2, 0, 1, 7, 4, 14, 18}

	utils.Part1(Day15Part1(values, 2020))
}

// History of the spoken numbers
type History map[int]*[2]int

// Day15Part1 solver
func Day15Part1(values []int, nth int) int {
	history := createHistory(values)

	n := len(values)
	lastNumber := values[n-1]

	for n < nth {
		n++
		lastNumber = history.step(n, lastNumber)
	}

	return lastNumber
}

func createHistory(values []int) *History {
	h := &History{}

	for ix, v := range values {
		h.track(ix+1, v)
	}

	return h
}

func (h *History) track(index int, value int) int {
	if (*h)[value] == nil {
		(*h)[value] = &[2]int{-1, -1}
	}
	(*h)[value][0] = (*h)[value][1]
	(*h)[value][1] = index
	return value
}

func (h *History) saidMoreThanOnce(value int) bool {
	arr := (*h)[value]
	return arr != nil && arr[0] != -1
}

func (h *History) step(index int, lastNumber int) int {
	heardManyTimes := h.saidMoreThanOnce(lastNumber)
	if heardManyTimes {
		relative := h.twiceDiff(lastNumber)
		return h.track(index, relative)
	} else {
		return h.track(index, 0)
	}
}

func (h *History) twiceDiff(value int) int {
	arr := (*h)[value]
	return arr[1] - arr[0]
}
