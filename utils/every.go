package utils

func Every[T any](collection []T, condition func(T) bool) bool {
	for _, item := range collection {
		if !condition(item) {
			return false
		}
	}
	return true
}
