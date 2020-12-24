package main

import (
	"log"
	"regexp"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	input, err := utils.ReadInputToList("day8/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Part1(Day8Part1(input))
}

// Day8Part1 solver
func Day8Part1(input []string) int {
	program := parseProgram(input)

	visited := map[int]bool{}

	for {
		if visited[program.index] {
			break
		}
		visited[program.index] = true
		program.step()
	}

	return program.acc
}

var instructionRegex = regexp.MustCompile(`^(\w+) \+?([\d-]+)$`)

func parseProgram(input []string) *program {
	p := &program{}
	for _, v := range input {
		match := instructionRegex.FindStringSubmatch(v)
		op := match[1]
		value, _ := utils.ToInteger(match[2])
		ins := &instruction{op: op, value: value}
		p.addInstruction(ins)
	}
	return p
}

type instruction struct {
	op    string
	value int
}

type program struct {
	index        int
	acc          int
	instructions []*instruction
}

func (p *program) addInstruction(ins *instruction) {
	p.instructions = append(p.instructions, ins)
}

func (p *program) step() {
	ins := p.instructions[p.index]
	switch ins.op {
	case "acc":
		p.acc += ins.value
		p.index++
	case "jmp":
		p.index += ins.value
	default:
		p.index++
	}
}
