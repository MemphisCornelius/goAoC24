package util

func Swap[T any](slice []T, i, j int) []T {
	slice[i], slice[j] = slice[j], slice[i]
	return slice
}

func Index[T comparable](slice []T, f T) int {
	for i, v := range slice {
		if v == f {
			return i
		}
	}
	return -1
}

func Remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}
