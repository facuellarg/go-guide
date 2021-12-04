package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	//Test invalid input
	t.Run("Test invalid input",
		func(t *testing.T) {
			var a, b int
			want := ErrZeroInput
			val, err := Sum(a, b)
			if !errors.Is(err, want) || val != 0 {
				t.Fatalf(
					"Sum(%d,%d) = %d,%v, want match with %d, %v",
					a, b, val, err, 0, want,
				)
			}
		})
}

//Test invalid input
func TestSumInvalidInput(t *testing.T) {
	var a, b int
	want := ErrZeroInput
	val, err := Sum(a, b)
	if !errors.Is(err, want) || val != 0 {
		t.Fatalf(
			"Sum(%d,%d) = %d,%v, want match with %d, %v",
			a, b, val, err, 0, want,
		)
	}
}

//Testing With assert a invalid input
func TestSumInvalidInputAssert(t *testing.T) {
	assert := assert.New(t)
	var a, b int
	want := ErrZeroInput
	val, err := Sum(a, b)
	assert.Equal(val, 0)
	assert.ErrorIs(err, want)
}

//Answer that we want for normal input
func TestSumAssert(t *testing.T) {
	assert := assert.New(t)
	var a, b int
	a, b = 4, 3
	val, err := Sum(a, b)
	assert.Equal(val, 7)
	assert.Nil(err)

}
