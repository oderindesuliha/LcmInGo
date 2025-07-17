package main

import "testing"

func TestToFindLcmOfNumbers(t *testing.T) {
	testCases := []struct {
		numbers  []int
		expected int
	}{
		{[]int{14, 12, 7, 8}, 168},
		{[]int{8, 24}, 24},
		{[]int{36, 24}, 72},
		{[]int{15, 8}, 120},
		{[]int{}, 0},
	}
	for _, testCase := range testCases {
		actual := LCM(testCase.numbers)
		if actual != testCase.expected {
			t.Errorf("Expected %d, got %d", testCase.expected, actual)
		}
	}
}
