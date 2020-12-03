package main

import (
	"testing"
)

func TestDay2Part1(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	result := Day2Part1(input)
	expected := 2

	if result != expected {
		t.Fatalf("The provided example case failed, expected %v, got %v", expected, result)
	}
}
