package utils

// func Filter[T](array []T, filterFn func(T) bool) []T {
// 	var result []T
// 	for _, t := range array {
// 		if filterFn(str) {
// 			result = append(result, str)
// 		}
// 	}
// 	return result
// }

func Includes[T comparable](array []T, value T) bool {
	for _, element := range array {
		if element == value {
			return true
		}
	}
	return false
}
