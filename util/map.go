package util

func CopyMap[T comparable, S any](m map[T]S) map[T]S {
	m2 := make(map[T]S, len(m))

	for k, v := range m {
		m2[k] = v
	}

	return m2
}
