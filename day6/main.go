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
	utils.Part2(Day6Part2(input))
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

// Day6Part2 solver
func Day6Part2(input string) int {
	lineRegex := regexp.MustCompile(`\n`)
	whitespaceRegex := regexp.MustCompile(`\s+`)
	trimmedInput := strings.TrimSpace(input)
	groups := strings.Split(trimmedInput, "\n\n")

	yesAnswers := 0

	for _, s := range groups {
		people := len(lineRegex.FindAllStringIndex(s, -1)) + 1
		answers := whitespaceRegex.ReplaceAllString(s, "")
		yesAnswers += getAllYes(answers, people)
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

func getAllYes(input string, count int) int {
	m := make(map[string]int)
	for _, s := range input {
		m[string(s)]++
	}

	maxes := 0

	for _, v := range m {
		if v == count {
			maxes++
		}
	}

	return maxes
}
