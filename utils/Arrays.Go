package arrays

func filter[T](array []T, filterFn func(T) bool) []T {
	var result []T
	for _, t := range array {
		if filterFn(str) {
			result = append(result, str)
		}
	}
	return result
}
