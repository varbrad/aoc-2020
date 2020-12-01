package solver

// Part1 solver
func Part1(numbers []int, target int) int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			a, b := numbers[i], numbers[j]
			if sum := a + b; sum == target {
				return a * b
			}
		}
	}
	return -1
}
