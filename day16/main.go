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
	utils.Part2(Day16Part2(string(input), "departure"))
}

// Day16Part1 solver
func Day16Part1(rawInput string) int {
	input := parseInput(rawInput)
	return input.getTotalErrorRate()
}

// Day16Part2 solver
func Day16Part2(rawInput string, startsWith string) int {
	input := parseInput(rawInput)
	input.removeInvalidNearbyTickets()

	fieldOrder := input.calculateFieldOrder()

	ticketData := map[string]int{}
	for ix, value := range input.yourTicket {
		fieldName := fieldOrder[ix]
		ticketData[fieldName] = value
	}

	prod := 1
	for key, value := range ticketData {
		if strings.HasPrefix(key, startsWith) {
			prod *= value
		}
	}
	return prod
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

var fieldsRegex = regexp.MustCompile(`(\w{1}[\w\s]+): (\d+)-(\d+) or (\d+)-(\d+)`)
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

func (i *Input) calculateFieldOrder() []string {
	possibles := getPossibleFields(i.yourTicket, i.fields)

	for _, ticket := range i.nearbyTickets {
		thisPossibles := getPossibleFields(ticket, i.fields)
		possibles = combinePossibleFields(possibles, thisPossibles)
	}

	index := -1
	for flatLength(possibles) > len(possibles) {
		possibles, index = flattenPossibles(possibles, index)
	}

	final := []string{}
	for _, v := range possibles {
		final = append(final, v[0])
	}

	return final
}

func (i *Input) removeInvalidNearbyTickets() {
	newNearbyTickets := []Ticket{}
	for _, v := range i.nearbyTickets {
		if isValidTicket(v, i.fields) {
			newNearbyTickets = append(newNearbyTickets, v)
		}
	}
	i.nearbyTickets = newNearbyTickets
}

func (i *Input) getTotalErrorRate() int {
	sum := 0
	for _, t := range i.nearbyTickets {
		sum += getErrorRate(t, i.fields)
	}
	return sum
}

// Helper methods
func flattenPossibles(possibles [][]string, ix int) ([][]string, int) {
	lenOneIx := findLengthOne(possibles, ix)
	if lenOneIx == -1 {
		return possibles, -1
	}
	val := possibles[lenOneIx][0]

	newPossibles := [][]string{}

	for ix := range possibles {
		newPossibles = append(newPossibles, []string{})
		if ix == lenOneIx {
			newPossibles[ix] = append(newPossibles[ix], val)
			continue
		} else {
			newPossibles[ix] = removeFromArr(possibles[ix], val)
		}
	}

	return newPossibles, lenOneIx
}

func findLengthOne(arr [][]string, ix int) int {
	for i, v := range arr {
		if i <= ix {
			continue
		}
		if len(v) == 1 {
			return i
		}
	}
	return -1
}

func flatLength(arr [][]string) int {
	length := 0
	for _, v := range arr {
		length += len(v)
	}
	return length
}

func combinePossibleFields(a [][]string, b [][]string) [][]string {
	final := [][]string{}

	for ix, groups := range a {
		final = append(final, []string{})
		for _, key := range groups {
			if isInArray(b[ix], key) {
				final[ix] = append(final[ix], key)
			}
		}
	}

	return final
}

func removeFromArr(arr []string, value string) []string {
	vs := []string{}
	for _, v := range arr {
		if v == value {
			continue
		}
		vs = append(vs, v)
	}
	return vs
}

func isInArray(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func getPossibleFields(t Ticket, fm FieldMap) [][]string {
	possibles := [][]string{}
	for ix, v := range t {
		possibles = append(possibles, []string{})
		for key, field := range fm {
			if isInField(field, v) {
				possibles[ix] = append(possibles[ix], key)
			}
		}
	}
	return possibles
}

func getErrorRate(t Ticket, fm FieldMap) int {
	sum := 0
	for _, v := range t {
		if !isInFieldMap(fm, v) {
			sum += v
		}
	}
	return sum
}

func isValidTicket(t Ticket, fm FieldMap) bool {
	return getErrorRate(t, fm) == 0
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
