package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	input, err := utils.ReadInput("day16/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Part1(Day16Part1(string(input)))
}

// Day16Part1 solver
func Day16Part1(rawInput string) int {
	input := parseInput(rawInput)
	return input.getTotalErrorRate()
}

// Range type
type Range = [2]int

// Field type
type Field = [2]Range

// FieldMap type
type FieldMap = map[string]Field

// Ticket numbers
type Ticket = []int

// Input parsed as a struct
type Input struct {
	fields        FieldMap
	yourTicket    Ticket
	nearbyTickets []Ticket
}

var fieldsRegex = regexp.MustCompile(`(\w+): (\d+)-(\d+) or (\d+)-(\d+)`)
var yourTicketRegex = regexp.MustCompile(`your ticket:\s*([\d,]+)\s`)
var nearbyTicketsRegex = regexp.MustCompile(`nearby tickets:\s+((?:[\d,]+\s*)*)`)

func parseInput(input string) Input {
	i := Input{
		fields: FieldMap{},
	}

	fieldsMatch := fieldsRegex.FindAllStringSubmatch(input, -1)
	yourTicketMatch := yourTicketRegex.FindStringSubmatch(input)
	nearbyTicketsMatch := nearbyTicketsRegex.FindStringSubmatch(input)

	for _, res := range fieldsMatch {
		values, _ := utils.ToIntegers(res[2:])
		range1 := Range{values[0], values[1]}
		range2 := Range{values[2], values[3]}
		i.fields[res[1]] = Field{range1, range2}
	}

	i.yourTicket = parseTicket(yourTicketMatch[1])
	i.nearbyTickets = parseTickets(nearbyTicketsMatch[1])

	return i
}

func parseTicket(input string) Ticket {
	list, _ := utils.ToIntegers(strings.Split(strings.TrimSpace(input), ","))
	return list
}

func parseTickets(input string) []Ticket {
	ts := []Ticket{}
	for _, row := range strings.Split(strings.TrimSpace(input), "\n") {
		ts = append(ts, parseTicket(row))
	}
	return ts
}

func (i *Input) getTotalErrorRate() int {
	sum := 0
	for _, t := range i.nearbyTickets {
		sum += getErrorRate(t, i.fields)
	}
	return sum
}

// Helper methods
func getErrorRate(t Ticket, fm FieldMap) int {
	sum := 0
	for _, v := range t {
		if !isInFieldMap(fm, v) {
			sum += v
		}
	}
	return sum
}

func isInFieldMap(fm FieldMap, v int) bool {
	for _, f := range fm {
		if isInField(f, v) {
			return true
		}
	}
	return false
}

func isInField(f Field, v int) bool {
	return isInRange(f[0], v) || isInRange(f[1], v)
}

func isInRange(r Range, v int) bool {
	return v >= r[0] && v <= r[1]
}
