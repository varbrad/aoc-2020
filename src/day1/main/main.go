package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"varbrad/aoc-2020/src/day1/solver"
)

func main() {
	contents, err := ioutil.ReadFile("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	splitStrings := strings.Split(strings.TrimSpace(string(contents)), "\n")
	values := toInts(splitStrings)

	fmt.Printf("Part 1 = %v", solver.Part1(values, 2020))
	fmt.Println("")
	fmt.Printf("Part 2 = %v", solver.Part2(values, 2020))
	fmt.Println("")
}

func toInts(stringList []string) []int {
	ints := make([]int, len(stringList))
	for i, s := range stringList {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}
