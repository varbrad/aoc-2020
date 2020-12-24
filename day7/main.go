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

type bagType struct {
	name     string
	contains []container
}

type container struct {
	bag   *bagType
	count int
}

type bagMap map[string]*bagType

// Day7Part1 solver
func Day7Part1(input []string) int {
	bm := bagMap{}

	for _, str := range input {
		name := bagRegex.FindStringSubmatch(str)
		contains := containRegex.FindAllStringSubmatch(str, -1)
		bm.add(name[1])
		for _, match := range contains {
			count, _ := utils.ToInteger(match[1])
			bagName := match[2]
			bm.add(bagName)
			bm.addContainsBag(name[1], count, bagName)
		}
	}

	return bm.countHas("shiny gold")
}

func (b *bagMap) add(name string) *bagType {
	bag, exists := (*b)[name]

	if exists {
		return bag
	}

	newBag := &bagType{name: name, contains: []container{}}
	(*b)[name] = newBag

	return newBag
}

func (b *bagMap) addContainsBag(cont string, count int, bag string) {
	containerBag := (*b)[cont]
	innerBag := (*b)[bag]

	contained := container{bag: innerBag, count: count}

	containerBag.contains = append(containerBag.contains, contained)
}

func (b *bagMap) countHas(name string) int {
	count := 0
	for _, v := range *b {
		if v.name == name {
			continue
		}
		if v.has(name) {
			count++
		}
	}
	return count
}

func (b *bagType) has(name string) bool {
	if b.name == name {
		return true
	}
	if len(b.contains) == 0 {
		return false
	}
	for _, c := range b.contains {
		contained := c.bag.has(name)
		if contained {
			return true
		}
	}
	return false
}
