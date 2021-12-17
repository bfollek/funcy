// gob build pkg/funcy.go

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
