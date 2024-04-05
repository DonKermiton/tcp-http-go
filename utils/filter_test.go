package utils

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	testCases := []struct {
		name      string
		input     []int
		condition func(int) bool
		expected  []int
	}{
		{
			name:      "even numbers",
			input:     []int{1, 2, 3, 4, 5},
			condition: func(x int) bool { return x%2 == 0 },
			expected:  []int{2, 4},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Filter(testCase.input, testCase.condition)

			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected %v, recived %v", testCase.expected, result)
			}
		})
	}
}
