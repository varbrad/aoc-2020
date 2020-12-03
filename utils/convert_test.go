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
