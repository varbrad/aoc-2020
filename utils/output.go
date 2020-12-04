package utils

import "fmt"

// Part1 output part 1 text
func Part1(a ...interface{}) {
	log("Part 1 =", a...)
}

// Part2 output part 2 text
func Part2(a ...interface{}) {
	log("Part 2 =", a...)
}

func log(prefix string, a ...interface{}) {
	args := append([]interface{}{prefix}, a...)
	fmt.Println(args...)
}
