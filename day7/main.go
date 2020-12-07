package main

import (
	"log"
	"regexp"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	input, err := utils.ReadInputToList("day7/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Part1(Day7Part1(input))
}

var bagRegex = regexp.MustCompile(`^([\w\s]+) bags contain`)
var containRegex = regexp.MustCompile(`(\d+) ([\w\s]+) bags?`)

type bagMap map[string]int

// Day7Part1 solver
func Day7Part1(input []string) int {
	m := map[string]bagMap{}

	for _, s := range input {
		bagColor := bagRegex.FindStringSubmatch(s)[1]
		contains := containRegex.FindAllStringSubmatch(s, -1)
		m[bagColor] = map[string]int{}

		for _, result := range contains {
			count, _ := utils.ToInteger(result[1])
			m[bagColor][result[2]] = count
		}
	}

	bagStack := []string{"shiny gold"}
	validBags := map[string]bool{}

	for len(bagStack) > 0 {
		bag := bagStack[0]
		bagStack = bagStack[1:]

		// Go thru the map and find any where bag is in the map
		for key, contained := range m {
			if key == bag || contains(bagStack, key) {
				continue
			}
			in := contained[bag] > 0
			if in {
				bagStack = append(bagStack, key)
				validBags[key] = true
			}
		}
	}

	return len(validBags)
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
