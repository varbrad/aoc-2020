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

// Part2 solver
func Part2(numbers []int, target int) int {
	for i := 0; i < len(numbers)-2; i++ {
		for j := i + 1; j < len(numbers)-1; j++ {
			for k := j + 1; k < len(numbers); k++ {
				a, b, c := numbers[i], numbers[j], numbers[k]
				if sum := a + b + c; sum == target {
					return a * b * c
				}
			}
		}
	}
	return -1
}
