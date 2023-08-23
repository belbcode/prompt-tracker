package utils

func Includes[T comparable](array []T, value T) bool {
	for _, element := range array {
		if element == value {
			return true
		}
	}
	return false
}

func MapToString[T any](array []T, mapFn func(T) string) []string {
	mappedArray := make([]string, len(array))
	for index, element := range array {
		mappedArray[index] = mapFn(element)
	}
	return mappedArray
}

// func Sort[T any](array *[]T, compareFn func(T) int) {
// 	ref := *array
// 	valueArray := createValueArray[T](ref, compareFn)
// 	return
// }

// func swap[T any](array *[]T, indexA int, indexB int) {

// }

// func createValueArray[T any](array []T, compareFn func(T) int) []int {
// 	var valueRepr []int = make([]int, len(array))
// 	for index, element := range array {
// 		valueRepr[index] = compareFn(element)
// 	}
// 	return valueRepr
// }
