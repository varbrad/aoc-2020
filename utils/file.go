package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// ReadInput Reads a file relative to the repository root
func ReadInput(path string) ([]byte, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fullPath := filepath.Join(cwd, path)
	return ioutil.ReadFile(fullPath)
}

// ReadInputToList Reads a file and formats it as a slice of strings
func ReadInputToList(path string) ([]string, error) {
	contents, err := ReadInput(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(contents)), "\n"), nil
}

// ReadInputToIntegerList Reads a file and formats as a slice of integers
func ReadInputToIntegerList(path string) ([]int, error) {
	rows, err := ReadInputToList(path)
	if err != nil {
		return nil, err
	}
	return ToIntegers(rows), nil
}
