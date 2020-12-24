package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay11Part1(t *testing.T) {
	layout := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}

	expected := 37
	actual := Day11Part1(layout)

	assert.Equal(t, expected, actual)
}
