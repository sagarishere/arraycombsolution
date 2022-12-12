package main

import "testing"

func TestSolution(t *testing.T) {
	// Define a table of test cases
	testCases := []struct {
		name     string
		A        []int
		expected int
	}{
		{"return3", []int{13, 7, 2, 8, 3}, 3},
		{"return1", []int{1, 2, 4, 8}, 1},
		{"return2", []int{16, 16}, 2},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Solution(testCase.A)
			if result != testCase.expected {
				t.Errorf("got %d, want %d", result, testCase.expected)
			}
		})
	}
}
