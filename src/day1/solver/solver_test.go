package solver

import (
	"testing"
)

func TestExampleCasePart1(t *testing.T) {
	values := []int{1721, 979, 366, 299, 675, 1456}
	result := Part1(values, 2020)
	expected := 514579

	if result != expected {
		t.Fatalf("The provided example case failed, expected %v, got %v", expected, result)
	}
}

func TestExampleCasePart2(t *testing.T) {
	values := []int{1721, 979, 366, 299, 675, 1456}
	result := Part2(values, 2020)
	expected := 241861950

	if result != expected {
		t.Fatalf("The provided example case failed, expected %v, got %v", expected, result)
	}
}