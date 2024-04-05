package utils

import (
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	cases := []struct {
		name      string
		input     []string
		expected  string
		condition func(string) bool
	}{
		{"Find element", []string{"w", "3", "test"}, "test", func(s string) bool {
			return s == "test"
		}},
		{"Not found", []string{"1,3", "test", "value"}, "", func(s string) bool {
			return s == "q"
		}},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			result, _ := Find(testCase.input, testCase.condition)

			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected %v, recived %v", testCase.expected, result)
			}
		})
	}
}
