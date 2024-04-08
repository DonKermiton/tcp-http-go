package utils

import "strings"

func TrimElementsInArray(array []string) []string {
	trimmedArray := make([]string, len(array))

	for i, item := range array {
		trimmedArray[i] = strings.TrimSpace(item)
	}

	return trimmedArray
}
