package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	input, err := utils.ReadInputToList("day2/input")

	if err != nil {
		log.Panic(err)
		return
	}

	fmt.Printf("Part 1 = %v", Day2Part1(input))
	fmt.Println("")
}

// Day2Part1 solver
func Day2Part1(input []string) int {
	count := 0

	for _, row := range input {
		policy, password := parseInput(row)
		if policy.isValid(password) {
			count++
		}
	}

	return count
}

func parseInput(row string) (passwordPolicy, string) {
	regex := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w+)$`)
	result := regex.FindStringSubmatch(row)[1:]

	min, _ := utils.ToInteger(result[0])
	max, _ := utils.ToInteger(result[1])
	character := result[2]
	password := result[3]

	return passwordPolicy{
		min:       min,
		max:       max,
		character: character,
	}, password
}

type passwordPolicy struct {
	min       int
	max       int
	character string
}

func (policy passwordPolicy) isValid(password string) bool {
	regex := regexp.MustCompile(policy.character)
	matches := regex.FindAllString(password, -1)
	count := len(matches)
	return policy.min <= count && policy.max >= count
}
