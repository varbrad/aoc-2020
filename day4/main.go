package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	contents, err := utils.ReadInput("day4/input")

	if err != nil {
		log.Fatal(err)
	}

	input := string(contents)

	fmt.Printf("Part 1 = %v", Day4Part1(input))
	fmt.Println("")
}

// Day4Part1 solver
func Day4Part1(input string) int {
	trimmedInput := strings.TrimSpace(input)
	groups := strings.Split(trimmedInput, "\n\n")

	valid := 0
	for _, group := range groups {
		if isValidGroup(group) {
			valid++
		}
	}

	return valid
}

func isValidGroup(group string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	propertyRegex := regexp.MustCompile(`(\w+):\S+`)
	results := propertyRegex.FindAllStringSubmatch(group, -1)

	for _, field := range requiredFields {
		found := false
		for _, data := range results {
			key := data[1]
			if field == key {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
