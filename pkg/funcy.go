package funcy

func Filter[T any](sl []T, test func(T) bool) []T {
	rv := make([]T, 0, len(sl))
	for _, elem := range sl {
		if test(elem) {
			rv = append(rv, elem)
		}
	}
	return rv
}

func Map[T1, T2 any](sl []T1, transform func(T1) T2) []T2 {
	rv := make([]T2, len(sl))
	for i, elem := range sl {
			rv[i] = transform(elem)
	}
	return rv
}

func Reduce[T1, T2 any](sl []T1, startValue T2, fReduce func(T2, T1) T2) T2 {
	accumulator := startValue
	for _, elem := range sl {
		accumulator = fReduce(accumulator, elem)
	}
	return accumulator
}

func Transpose[T any](sl [][]T)[][]T {
	return sl
}
