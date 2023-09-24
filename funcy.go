// package funcy implements functional favorites like filter, map, and reduce.
//
// You'll get a compile error if you try something that doesn't make sense.
// For example, using map to run strings.ToLower on a slice of ints:
//
//	sl := []int{1, 2, 3, 4}
//	result := Map(sl, strings.ToLower)
//
// will get you an error like
//
//	pkg/funcy_test.go:64:20: type func(s string) string of strings.ToLower
//	does not match inferred type func(int) T2 for func(T1) T2
package funcy

import (
	"fmt"
	"log"

	"golang.org/x/exp/constraints"
)

const jaggedTransposeError = "all rows must be the same size as the zero row (len == %d). Row %d is not the same size (len == %d)"

// Filter takes a slice of any type and a predicate function. It passes each
// slice item to the predicate function, and returns a slice of the items
// for which the predicate function returns true.
func Filter[T any](sl []T, test func(T) bool) []T {
	// Just use FilterWithIndex and a wrapper func that ignores the index.
	return FilterWithIndex(sl, func(_ int, item T) bool {
		return test(item)
	})
}

// FilterWithIndex is like Filter, but the predicate function receives
// two arguments. The second argument is the slice item; the first
// argument is the item's index within the slice. FilterWithIndex
// returns a list of the items for which the predicate function returns true.
func FilterWithIndex[T any](sl []T, test func(int, T) bool) []T {
	rv := make([]T, 0, len(sl))
	for i, elem := range sl {
		if test(i, elem) {
			rv = append(rv, elem)
		}
	}
	return rv
}

// Map runs each item in a slice through a transformation function,
// and returns a slice of the transformed items. The transformed items
// may be a different type from the input items, e.g. strings to ints
// using strconv.Atoi.
//
// Note: For tranformation functions that should check errors, like
// strconv.Atoi, you can wrap the transformation function in a function
// that handles the error:
//
//	result := Map(input, func(s string) int {
//		i, err := strconv.Atoi(s)
//		if err != nil {
//			log.Fatal(err)
//		}
//		return i
//	})
func Map[T1, T2 any](sl []T1, transform func(T1) T2) []T2 {
	rv := make([]T2, len(sl))
	for i, elem := range sl {
		rv[i] = transform(elem)
	}
	return rv
}

// Reduce reduces a slice to a single value by running each item of the slice
// through a function that takes an accumulator and the next item
// as its paramaters. The classic example is reducing a slice of numbers by
// adding them together.
func Reduce[T1, T2 any](sl []T1, startValue T2, fReduce func(T2, T1) T2) T2 {
	accumulator := startValue
	for _, elem := range sl {
		accumulator = fReduce(accumulator, elem)
	}
	return accumulator
}

// Sum adds together the items in a slice.
func Sum[T constraints.Ordered](sl []T) T {
	// Alternative: reflect.Zero(T)
	var zeroValue T
	return Reduce(sl, zeroValue, func(acc T, nxt T) T {
		return acc + nxt
	})
}

// Transpose converts a matrix from T[rows][columns] to T[columns][rows].
// The matrix cannot be jagged, i.e. all rows must have the same number
// of elements.
func Transpose[T any](sl [][]T) ([][]T, error) {
	numRows := len(sl)
	var numCols int
	if numRows > 0 {
		numCols = len(sl[0])
	}
	// Edge cases where there's nothing to do.
	if numRows == 0 || (numRows == 1 && numCols == 0) {
		return sl, nil
	}
	// Create `rv`, an empty slice of slices.
	rv := make([][]T, numCols) // Columns transposed to rows.
	for i := range rv {
		rv[i] = make([]T, numRows) // Rows transposed to columns.
	}
	// Fill in `rv`.
	for i, row := range sl {
		if len(row) != numCols {
			return nil, fmt.Errorf(jaggedTransposeError, numCols, i, len(row))
		}
		for j := 0; j < numCols; j++ {
			rv[j][i] = sl[i][j]
		}
	}
	return rv, nil
}

// MustTranspose calls Transpose and panics on error.
func MustTranspose[T any](sl [][]T) [][]T {
	rv, err := Transpose(sl)
	if err != nil {
		log.Panic(err)
	}
	return rv
}
