package main

import "errors"

var (
	ErrZeroInput = errors.New("invalid input")
)

func Sum(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, ErrZeroInput
	}
	return a + b, nil
}
