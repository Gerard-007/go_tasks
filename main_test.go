package main

import (
    "testing"
)


func TestFindFactors(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{100, 36, 28, 12}, 6},
		{[]int{1}, 1},
	}

	for _, test := range tests {
		result := findTheHighestCommonFactor(test.input)
		if result != test.expected {
			t.Errorf("findTheHighestCommonFactor(%v) = %d; expected %d", test.input, result, test.expected)
		}
	}
}