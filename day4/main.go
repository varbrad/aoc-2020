package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/varbrad/aoc-2020/utils"
)

var validate *validator.Validate

func main() {
	contents, err := utils.ReadInput("day4/input")

	if err != nil {
		log.Fatal(err)
	}

	input := string(contents)

	fmt.Printf("Part 1 = %v", Day4Part1(input))
	fmt.Println("")
	fmt.Printf("Part 2 = %v", Day4Part2(input))
	fmt.Println("")
}

// Day4Part1 solver
func Day4Part1(input string) int {
	trimmedInput := strings.TrimSpace(input)
	groups := strings.Split(trimmedInput, "\n\n")

	valid := 0
	for _, group := range groups {
		if hasValidFields(group) {
			valid++
		}
	}

	return valid
}

// Day4Part2 solver
func Day4Part2(input string) int {
	trimmedInput := strings.TrimSpace(input)
	groups := strings.Split(trimmedInput, "\n\n")

	validate = validator.New()

	valid := 0
	for _, group := range groups {
		if isValidGroup(group) {
			valid++
		}
	}

	return valid
}

func hasValidFields(group string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	propertyRegex := regexp.MustCompile(`(\w+):\S+`)
	results := propertyRegex.FindAllStringSubmatch(group, -1)

	for _, field := range requiredFields {
		found := false
		for _, data := range results {
			key := data[1]
			if field == key {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func isValidGroup(group string) bool {
	propertyRegex := regexp.MustCompile(`(\w+):(\S+)`)
	results := propertyRegex.FindAllStringSubmatch(group, -1)
	propertyMap := make(map[string]string)

	for _, result := range results {
		key := strings.ToUpper(result[1])
		value := result[2]
		propertyMap[key] = value
	}

	return makePassport(propertyMap).isValid()
}

func makePassport(data map[string]string) *passport {
	birthYear, _ := utils.ToInteger(data["BYR"])
	expiryYear, _ := utils.ToInteger(data["EYR"])
	issueYear, _ := utils.ToInteger(data["IYR"])

	return &passport{
		BYR: birthYear,
		CID: data["CID"],
		ECL: data["ECL"],
		EYR: expiryYear,
		HCL: data["HCL"],
		HGT: data["HGT"],
		IYR: issueYear,
		PID: data["PID"],
	}
}

type passport struct {
	BYR int `validate:"required,min=1920,max=2002"`
	CID interface{}
	ECL string `validate:"required,oneof=amb blu brn gry grn hzl oth"`
	EYR int    `validate:"required,min=2020,max=2030"`
	HCL string `validate:"required,hexcolor"`
	HGT string `validate:"required"`
	IYR int    `validate:"required,min=2010,max=2020"`
	PID string `validate:"required,len=9"`
}

func (data *passport) isValid() bool {
	err := validate.Struct(data)
	if err != nil {
		return false
	}

	// Custom height validator
	heightRegex := regexp.MustCompile(`^(\d+)(cm|in)$`)
	result := heightRegex.FindStringSubmatch(data.HGT)

	if result == nil {
		return false
	}

	value, _ := utils.ToInteger(result[1])

	switch result[2] {
	case "in":
		return validate.Var(value, "min=59,max=76") == nil
	case "cm":
		return validate.Var(value, "min=150,max=193") == nil
	}

	return true
}
