package slices

func Map[T, U any](values []T, f func(T) U) []U {
	mappedValues := make([]U, len(values))

	for i := range values {
		mappedValues[i] = f(values[i])
	}

	return mappedValues
}
