package main

import (
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

	utils.Part1(Day2Part1(input))
	utils.Part2(Day2Part2(input))
}

// Day2Part1 solver
func Day2Part1(input []string) int {
	count := 0

	for _, row := range input {
		policy, password := parseInput(row)
		if policy.isValidLegacy(password) {
			count++
		}
	}

	return count
}

// Day2Part2 solver
func Day2Part2(input []string) int {
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

func (policy passwordPolicy) isValidLegacy(password string) bool {
	regex := regexp.MustCompile(policy.character)
	matches := regex.FindAllString(password, -1)
	count := len(matches)
	return policy.min <= count && policy.max >= count
}

func (policy passwordPolicy) isValid(password string) bool {
	charByte := []byte(policy.character)[0]
	atMin := password[policy.min-1] == charByte
	atMax := password[policy.max-1] == charByte

	return atMin != atMax
}
