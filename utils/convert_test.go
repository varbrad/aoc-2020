package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToIntegerHandlesABasicCase(t *testing.T) {
	expected := 2491
	actual, err := ToInteger("2491")

	assert.Equal(t, expected, actual)
	assert.Nil(t, err)
}

func TestToIntegerReturnsAnError(t *testing.T) {
	_, err := ToInteger("2491a")

	assert.Error(t, err)
}

func TestToIntegersHandlesABasicCase(t *testing.T) {
	expected := []int{1, 2, 3}
	actual, err := ToIntegers([]string{"1", "2", "3"})

	assert.Equal(t, expected, actual)
	assert.Nil(t, err)
}

func TestToIntegersReturnsAnyErrors(t *testing.T) {
	_, err := ToIntegers([]string{"1", "b", "3"})

	assert.Error(t, err)
}
