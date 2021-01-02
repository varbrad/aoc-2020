package main

import (
	"log"
	"regexp"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	input, err := utils.ReadInputToList("day14/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Part1(Day14Part1(input))
}

// Instruction a memory set instruction
type Instruction struct {
	op   string
	line string
}

// Mask mask type
type Mask = map[uint]bool

// Memory 36-bit address space
type Memory struct {
	data map[uint]int
	mask Mask
	ins  []Instruction
}

// Day14Part1 solver
func Day14Part1(input []string) int {
	mem := parseInput(input)
	mem.process()
	return mem.sum()
}

var typeRegex = regexp.MustCompile(`(\w+).* = `)
var maskRegex = regexp.MustCompile(`mask = ([X10]{36})`)
var memRegex = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

func parseInput(input []string) *Memory {
	mem := &Memory{
		data: map[uint]int{},
		mask: nil,
		ins:  make([]Instruction, len(input)),
	}

	for ix, v := range input {
		mem.ins[ix] = Instruction{
			line: v,
			op:   typeRegex.FindStringSubmatch(v)[1],
		}
	}

	return mem
}

func (m *Memory) process() {
	for ix := range m.ins {
		m.step(ix)
	}
}

func (m *Memory) step(index int) {
	ins := &m.ins[index]

	switch ins.op {
	case "mask":
		m.setMask(ins)
	case "mem":
		m.setMem(ins)
	}
}

func (m *Memory) setMask(ins *Instruction) {
	mask := Mask{}
	result := maskRegex.FindStringSubmatch(ins.line)
	for ix, v := range result[1] {
		char := string(v)
		if char == "X" {
			continue
		}
		mask[uint(35-ix)] = char == "1"
	}
	m.mask = mask
}

func (m *Memory) setMem(ins *Instruction) {
	result := memRegex.FindStringSubmatch(ins.line)
	index, _ := utils.ToInteger(result[1])
	value, _ := utils.ToInteger(result[2])
	m.data[uint(index)] = maskValue(value, m.mask)
}

func (m *Memory) sum() int {
	sum := 0
	for _, v := range m.data {
		sum += v
	}
	return sum
}

func setBit(n int, pos uint) int {
	return n | (1 << pos)
}

func clearBit(n int, pos uint) int {
	return n &^ (1 << pos)
}

func maskValue(n int, mask Mask) int {
	value := n
	for ix, v := range mask {
		if v == true {
			value = setBit(value, ix)
		} else {
			value = clearBit(value, ix)
		}
	}
	return value
}
