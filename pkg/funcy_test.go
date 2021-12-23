package funcy

import (
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilterWithInts(t *testing.T) {
	require := require.New(t)
	expecting := []int{2, 4, 6, 8}
	sl := []int{2, 3, 4, 5, 6, 7, 8, 9}
	result := Filter(sl, func(i int) bool { return i%2 == 0 })
	require.Equal(expecting, result)
}

func TestFilterWithStrings(t *testing.T) {
	require := require.New(t)
	expecting := []string{"bat", "cat", "hat"}
	sl := []string{"bat", "monkey", "cat", "submarine", "me", "hat", "lunch"}
	result := Filter(sl, func(s string) bool { return len(s) == 3 })
	require.Equal(expecting, result)
}

func TestFilterWithEmptySlice(t *testing.T) {
	require := require.New(t)
	expecting := []string{}
	sl := []string{}
	result := Filter(sl, func(s string) bool { return len(s) == 3 })
	require.Equal(expecting, result)
}

func TestMapWithInts(t *testing.T) {
	require := require.New(t)
	expecting := []int{4, 6, 8, 10}
	sl := []int{2, 3, 4, 5}
	result := Map(sl, func(i int) int { return i * 2 })
	require.Equal(expecting, result)
}

func TestMapWithStrings(t *testing.T) {
	require := require.New(t)
	expecting := []string{"bat", "cat", "hat"}
	sl := []string{"BAT", "CaT", "haT"}
	result := Map(sl, strings.ToLower)
	require.Equal(expecting, result)
}

func TestMapWithStringsToInts(t *testing.T) {
	require := require.New(t)
	expecting := []int{1, 3, 5, 7, 9, 11}
	sl := []string{"1", "3", "5", "7", "9", "11"}
	result := Map(sl, func(s string) int {
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
	expecting := []int{}
	sl := []int{}
	result := Map(sl, func(i int) int { return i * 2 })
	require.Equal(expecting, result)
}

func TestReduceAddInts(t *testing.T) {
	require := require.New(t)
	expecting := 14
	sl := []int{2, 3, 4, 5}
	result := Reduce(sl, 0, func(i, j int) int { return i + j })
	require.Equal(expecting, result)
}

func TestReduceAddStrings(t *testing.T) {
	require := require.New(t)
	expecting := "home-made"
	sl := []string{"home", "-", "made"}
	result := Reduce(sl, "", func(s, t string) string { return s + t })
	require.Equal(expecting, result)
}

func TestReduceCanDoMap(t *testing.T) {
	require := require.New(t)
	expecting := []string{"bat", "cat", "hat"}
	sl := []string{"BAT", "CaT", "haT"}
	result := Reduce(sl, []string{}, func(acc []string, s string) []string {
		return append(acc, strings.ToLower(s))
	})
	require.Equal(expecting, result)
}

func TestReduceWithEmptySlice(t *testing.T) {
	require := require.New(t)
	startValue := ""
	expecting := startValue
	sl := []string{}
	result := Reduce(sl, startValue, func(s, t string) string { return s + t })
	require.Equal(expecting, result)
}

func TestTransposeSquareInts(t *testing.T) {
	require := require.New(t)
	expecting := [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}
	sl := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	result := Transpose(sl)
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
	result := Transpose(input)
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
	result := Transpose(input)
	require.Equal(expecting, result)
}
