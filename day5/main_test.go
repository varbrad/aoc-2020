package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestDay5Part1(t *testing.T) {
	expected := 820
	actual := Day5Part1([]string{
		"FBFBBFFRLR",
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	})
	assert.Equal(t, expected, actual)
}

func TestGetSeatID(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		expected := 357
		actual := GetSeatID("FBFBBFFRLR")
		assert.Equal(t, expected, actual)
	})

	t.Run("test case 2", func(t *testing.T) {
		expected := 567
		actual := GetSeatID("BFFFBBFRRR")
		assert.Equal(t, expected, actual)
	})

	t.Run("test case 3", func(t *testing.T) {
		expected := 119
		actual := GetSeatID("FFFBBBFRRR")
		assert.Equal(t, expected, actual)
	})

	t.Run("test case 4", func(t *testing.T) {
		expected := 820
		actual := GetSeatID("BBFFBBFRLL")
		assert.Equal(t, expected, actual)
	})
}