package utils

func Find[T any](collection []T, condition func(T) bool) (T, int) {
	for index, item := range collection {
		if condition(item) {
			return item, index
		}
	}

	var zeroValue T
	return zeroValue, -1
}
