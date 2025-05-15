package main

import (
    "testing"
)


func TestFindFactors(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{100, 9},
		{36, 9},
		{28, 6},
		{12, 6},
		{1, 1},
		{0, 0},
		{-1, 0},
	}

	for _, test := range tests {
		result := findFactors(test.input)
		if result != test.expected {
			t.Errorf("findFactors(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}