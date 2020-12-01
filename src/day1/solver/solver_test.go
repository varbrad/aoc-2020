package solver

import (
	"testing"
)

func TestExampleCase(t *testing.T) {
	values := []int{1721, 979, 366, 299, 675, 1456}
	result := Part1(values, 2020)
	expected := 514578

	if Part1(values, 2020) != 514578 {
		t.Fatalf("The provided example case failed, expected %v, got %v", expected, result)
	}
}
