package utils

import "testing"

func TestEvery(t *testing.T) {
	testCases := []struct {
		name      string
		input     []int
		condition func(int) bool
		expected  bool
	}{
		{name: "Empty slice", input: []int{}, condition: func(i int) bool { return true }, expected: true},
		{name: "All even", input: []int{2, 4, 6}, condition: func(i int) bool { return i%2 == 0 }, expected: true},
		{name: "One odd", input: []int{2, 5, 6}, condition: func(i int) bool { return i%2 == 0 }, expected: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Every(tc.input, tc.condition)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
