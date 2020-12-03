package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInputCanParseAFile(t *testing.T) {
	expected := "1\n2\n3\n"
	actual, err := ReadInput("__tests__/list")

	assert.Equal(t, expected, string(actual))
	assert.Nil(t, err)
}

func TestReadInputHandlesNonExistantFile(t *testing.T) {
	_, err := ReadInput("__tests__/non-existant-file")

	assert.Error(t, err)
}

func TestReadInputToListCanParseAList(t *testing.T) {
	expected := []string{"1", "2", "3"}
	actual, err := ReadInputToList("__tests__/list")

	assert.Equal(t, expected, actual)
	assert.Nil(t, err)
}

func TestReadInputToIntegerListCanParseAList(t *testing.T) {
	expected := []int{1, 2, 3}
	actual, err := ReadInputToIntegerList("__tests__/list")

	assert.Equal(t, expected, actual)
	assert.Nil(t, err)
}
