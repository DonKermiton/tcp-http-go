package utils

import (
	"testing"
)

func TestTrimElementsInArray(t *testing.T) {
	testArray := []string{"  First ", "  Sec ond", "Third"}
	expect := []string{"First", "Sec ond", "Third"}
	result := TrimElementsInArray(testArray)

	for i := 0; i < len(result); i++ {
		if result[i] != expect[i] {
			t.Errorf("Expected: %s, got: %s", testArray[i], expect[i])
		}
	}

}
