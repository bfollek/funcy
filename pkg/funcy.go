package funcy

import (
	"constraints"
	"fmt"
)

const jaggedTransposeErrorFmt = "All rows must be the same size as the zero row (len == %d). Row %d is not the same size (len == %d)."

// Filter returns items from a slice that satisfy a predicate function.
func Filter[T any](sl []T, test func(T) bool) []T {
	rv := make([]T, 0, len(sl))
	for _, elem := range sl {
		if test(elem) {
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
		return acc + nxt})

}

// Transpose converts a matrix from T[rows][columns] to T[columns][rows].
// The matrix cannot be jagged, i.e. all rows must have the same number 
// of elements.
func Transpose[T any](sl [][]T)([][]T, error) {
	num_rows := len(sl)
	var num_cols int
	if num_rows > 0 {
		num_cols = len(sl[0])
	}
	// Edge cases where there's nothing to do.
	if num_rows == 0 || (num_rows == 1 && num_cols == 0) {
		return sl, nil
	}
	// Create `rv`, an empty slice of slices.
	rv := make([][]T, num_cols) // Columns transposed to rows.
	for i := range rv {
    	rv[i] = make([]T, num_rows) // Rows transposed to columns.
	}
	// Fill in `rv`.
	for i, row := range sl {
		if len(row) != num_cols {
			return nil, fmt.Errorf(jaggedTransposeErrorFmt, num_cols, i, len(row))
		}
		for j := 0; j < num_cols; j++ {
			rv[j][i] = sl[i][j]
		}
	}
	return rv, nil
}

