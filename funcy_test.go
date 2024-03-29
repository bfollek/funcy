package funcy

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilterWithInts(t *testing.T) {
	require := require.New(t)
	input := []int{2, 3, 4, 5, 6, 7, 8, 9}
	expecting := []int{2, 4, 6, 8}
	result := Filter(input, func(i int) bool { return i%2 == 0 })
	require.Equal(expecting, result)
}

func TestFilterWithStrings(t *testing.T) {
	require := require.New(t)
	input := []string{"bat", "monkey", "cat", "submarine", "me", "hat", "lunch"}
	expecting := []string{"bat", "cat", "hat"}
	result := Filter(input, func(s string) bool { return len(s) == 3 })
	require.Equal(expecting, result)
}

func TestFilterWithEmptySlice(t *testing.T) {
	require := require.New(t)
	input := []string{}
	expecting := []string{}
	result := Filter(input, func(s string) bool { return len(s) == 3 })
	require.Equal(expecting, result)
}

func TestFilterWithIndexWithInts(t *testing.T) {
	require := require.New(t)
	input := []int{1, 3, 5, 7, 9}
	expecting := []int{1, 5, 9}
	result := FilterWithIndex(input, func(index, _ int) bool { return index%2 == 0 })
	require.Equal(expecting, result)
}

func TestMapWithInts(t *testing.T) {
	require := require.New(t)
	input := []int{2, 3, 4, 5}
	expecting := []int{4, 6, 8, 10}
	result := Map(input, func(i int) int { return i * 2 })
	require.Equal(expecting, result)
}

func TestMapWithStrings(t *testing.T) {
	require := require.New(t)
	input := []string{"BAT", "CaT", "haT"}
	expecting := []string{"bat", "cat", "hat"}
	result := Map(input, strings.ToLower)
	require.Equal(expecting, result)
}

func TestMapWithStringsToInts(t *testing.T) {
	require := require.New(t)
	input := []string{"1", "3", "5", "7", "9", "11"}
	expecting := []int{1, 3, 5, 7, 9, 11}
	result := Map(input, func(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		return i
	})
	require.Equal(expecting, result)
}

func TestMapWithEmptySlice(t *testing.T) {
	require := require.New(t)
	input := []int{}
	expecting := []int{}
	result := Map(input, func(i int) int { return i * 2 })
	require.Equal(expecting, result)
}

func TestReduceAddInts(t *testing.T) {
	require := require.New(t)
	input := []int{2, 3, 4, 5}
	expecting := 14
	result := Reduce(input, 0, func(i, j int) int { return i + j })
	require.Equal(expecting, result)
}

func TestReduceAddStrings(t *testing.T) {
	require := require.New(t)
	input := []string{"home", "-", "made"}
	expecting := "home-made"
	result := Reduce(input, "", func(s, t string) string { return s + t })
	require.Equal(expecting, result)
}

func TestReduceCanDoMap(t *testing.T) {
	require := require.New(t)
	input := []string{"BAT", "CaT", "haT"}
	expecting := []string{"bat", "cat", "hat"}
	result := Reduce(input, []string{}, func(acc []string, s string) []string {
		return append(acc, strings.ToLower(s))
	})
	require.Equal(expecting, result)
}

func TestReduceWithEmptySlice(t *testing.T) {
	require := require.New(t)
	input := []string{}
	startValue := ""
	expecting := startValue
	result := Reduce(input, startValue, func(s, t string) string { return s + t })
	require.Equal(expecting, result)
}

func TestSumInts(t *testing.T) {
	require := require.New(t)
	input := []int{1, 2, 3, 4, 5}
	expecting := 15
	result := Sum(input)
	require.Equal(expecting, result)
}

func TestSumFloat64s(t *testing.T) {
	require := require.New(t)
	input := []float64{1.3, 2.7, 3.2, 4.8, 5.5}
	expecting := 17.5
	result := Sum(input)
	require.Equal(expecting, result)
}

func TestSumStrings(t *testing.T) {
	require := require.New(t)
	input := []string{"That ", "crazy ", "cat"}
	expecting := "That crazy cat"
	result := Sum(input)
	require.Equal(expecting, result)
}

func TestTransposeSquareInts(t *testing.T) {
	require := require.New(t)
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expecting := [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}
	result, err := Transpose(input)
	require.Nil(err)
	require.Equal(expecting, result)
}

func TestTransposeOneRowOneColInt(t *testing.T) {
	require := require.New(t)
	input := [][]int{
		{99},
	}
	expecting := input
	result, err := Transpose(input)
	require.Nil(err)
	require.Equal(expecting, result)
}

func TestTransposeMoreColsThanRowsInts(t *testing.T) {
	require := require.New(t)
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}
	expecting := [][]int{
		{1, 4, 7, 10},
		{2, 5, 8, 11},
		{3, 6, 9, 12},
	}
	result, err := Transpose(input)
	require.Nil(err)
	require.Equal(expecting, result)
}

func TestTransposeMoreRowsThanColsInts(t *testing.T) {
	require := require.New(t)
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	expecting := [][]int{
		{1, 4},
		{2, 5},
		{3, 6},
	}
	result, err := Transpose(input)
	require.Nil(err)
	require.Equal(expecting, result)
}

func TestTransposeEmptyRowsInts(t *testing.T) {
	require := require.New(t)
	input := [][]int{}
	expecting := input
	result, err := Transpose(input)
	require.Nil(err)
	require.Equal(expecting, result)
}

func TestTransposeEmptyColsInts(t *testing.T) {
	require := require.New(t)
	input := [][]int{{}}
	expecting := input
	result, err := Transpose(input)
	require.Nil(err)
	require.Equal(expecting, result)
}

func TestTransposeJaggedInts(t *testing.T) {
	require := require.New(t)
	input := [][]int{
		{1, 2, 3},
		{4, 5},
		{7, 8, 9},
	}
	expectingError := fmt.Sprintf(jaggedTransposeError, len(input[0]), 1, len(input[1]))
	_, err := Transpose(input)
	require.NotNil(err)
	require.Equal(expectingError, err.Error())
}

func TestTransposeJaggedAndEmptyInts(t *testing.T) {
	require := require.New(t)
	input := [][]int{
		{1, 2, 3},
		{},
		{7, 8, 9},
	}
	expectingError := fmt.Sprintf(jaggedTransposeError, len(input[0]), 1, len(input[1]))
	_, err := Transpose(input)
	require.NotNil(err)
	require.Equal(expectingError, err.Error())
}

func TestTransposeJaggedEmptyZerorowInts(t *testing.T) {
	require := require.New(t)
	input := [][]int{
		{},
		{4, 5},
		{6, 7},
	}
	expectingError := fmt.Sprintf(jaggedTransposeError, len(input[0]), 1, len(input[1]))
	_, err := Transpose(input)
	require.NotNil(err)
	require.Equal(expectingError, err.Error())
}

// https://stackoverflow.com/questions/31595791/how-to-test-panics
func TestMustTransposePanics(t *testing.T) {
	input := [][]int{
		{3},
		{4, 5},
		{6, 7},
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustTranspose did not panic with this jagged matrix: %v\n", input)
		}
	}()

	MustTranspose(input)
}
