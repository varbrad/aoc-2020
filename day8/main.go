package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/varbrad/aoc-2020/utils"
)

func main() {
	input, err := utils.ReadInputToList("day8/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Part1(Day8Part1(input))
	utils.Part2(Day8Part2(input))
}

// Day8Part1 solver
func Day8Part1(input []string) int {
	program := parseProgram(input)
	program.run()
	return program.acc
}

// Day8Part2 solver
func Day8Part2(input []string) int {
	for i, v := range input {
		isNop, isJmp := strings.HasPrefix(v, "nop"), strings.HasPrefix(v, "jmp")
		if !isNop && !isJmp {
			continue
		}
		inputCopy := append([]string(nil), input...)
		if isNop {
			inputCopy[i] = strings.Replace(v, "nop", "jmp", 1)
		} else {
			inputCopy[i] = strings.Replace(v, "jmp", "nop", 1)
		}
		program := parseProgram(inputCopy)
		ok := program.run()
		if ok {
			return program.acc
		}
	}

	return -1
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

func (p *program) run() bool {
	visited := map[int]bool{}
	totalIns := len(p.instructions)

	for {
		if visited[p.index] {
			return false
		}
		if totalIns == p.index {
			break
		}
		visited[p.index] = true
		p.step()
	}
	return true
}
