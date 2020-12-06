package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	contents, err := utils.ReadInput("day6/input")

	if err != nil {
		log.Fatal(err)
	}

	input := string(contents)

	utils.Part1(Day6Part1(input))
}

// Day6Part1 solver
func Day6Part1(input string) int {
	whitespaceRegex := regexp.MustCompile(`\s+`)
	trimmedInput := strings.TrimSpace(input)
	groups := strings.Split(trimmedInput, "\n\n")

	yesAnswers := 0

	for i, s := range groups {
		groups[i] = whitespaceRegex.ReplaceAllString(s, "")
		yesAnswers += getUniqueChars(groups[i])
	}

	return yesAnswers
}

func getUniqueChars(input string) int {
	m := make(map[string]bool)
	for _, s := range input {
		m[string(s)] = true
	}
	return len(m)
}
