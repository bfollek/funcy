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

// Transpose converts a matrix from T[rows][columns] to T[columns][rows].
// The matrix cannot be jagged, i.e. all rows must have the same number 
// of elements.

func Transpose[T any](sl [][]T)([][]T, error) {
	num_rows := len(sl)
	num_cols := len(sl[0])
	if num_rows == 0 {
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
			//errors.New("")
		}
		for j := 0; j < num_cols; j++ {
			rv[j][i] = sl[i][j]
		}
	}
	return rv, nil
}
