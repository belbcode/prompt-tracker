package utils

func Includes[T comparable](array []T, value T) bool {
	for _, element := range array {
		if element == value {
			return true
		}
	}
	return false
}
