package utils

func Filter[T any](collection []T, condition func(T) bool) (result []T) {
	for _, item := range collection {
		if condition(item) {
			result = append(result, item)
		}
	}
	return result
}
