package funcy

import (
	"fmt"
	"log"

	"golang.org/x/exp/constraints"
)

const jaggedTransposeError = "All rows must be the same size as the zero row (len == %d). Row %d is not the same size (len == %d)."

// Filter returns items from a slice that satisfy a predicate function.
func Filter[T any](sl []T, test func(T) bool) []T {
	// Just use FilterWithIndex and a wrapper func that ignores the index.
	return FilterWithIndex(sl, func(_ int, item T) bool {
		return test(item)
	})
}

// FilterWithIndex is like Filter, but the predicate function receives
// two arguments. The first is the int index of the second argument.
// Returns items from a slice that satisfy the predicate function.
func FilterWithIndex[T any](sl []T, test func(int, T) bool) []T {
	rv := make([]T, 0, len(sl))
	for i, elem := range sl {
		if test(i, elem) {
			rv = append(rv, elem)
		}
	}
	return rv
}

// Map runs each item in a slice through a transform function, and
// returns a slice of the transformed items. The transformed items may be a different
// type, e.g. strings to ints using strconv.Atoi.
func Map[T1, T2 any](sl []T1, transform func(T1) T2) []T2 {
	rv := make([]T2, len(sl))
	for i, elem := range sl {
		rv[i] = transform(elem)
	}
	return rv
}

// Reduce reduces each item of a slice to a single value, by running
// each item through a function that takes an accumulator and the next item
// as its paramaters. The classic example is reducing a slice of numbers by
// adding them together.
func Reduce[T1, T2 any](sl []T1, startValue T2, fReduce func(T2, T1) T2) T2 {
	accumulator := startValue
	for _, elem := range sl {
		accumulator = fReduce(accumulator, elem)
	}
	return accumulator
}

// Sum can be defined using Reduce.
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

// MustTranspose wraps `Transpose` and panics on error.
func MustTranspose[T any](sl [][]T) [][]T {
	rv, err := Transpose(sl)
	if err != nil {
		log.Panic(err)
	}
	return rv
}
