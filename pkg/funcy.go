// gob build pkg/funcy.go
// gob build ./...

package funcy

func Filter[T any](sl []T, test func(T) bool) []T {
	rv := []T{}
	for _, elem := range sl {
		if test(elem) {
			rv = append(rv, elem)
		}
	}
	return rv
}

func Map[T1, T2 any](sl []T1, transform func(T1) T2) []T2 {
	rv := make([]T2, len (sl))
	for i, elem := range sl {
			rv[i] = transform(elem)
	}
	return rv
}
