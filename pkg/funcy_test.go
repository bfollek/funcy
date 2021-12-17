// gob test ./...

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
	expected := []int{2, 4, 6, 8}
	sl := []int{2, 3, 4, 5, 6, 7, 8, 9}
	result := Filter(sl, func(i int) bool { return i%2 == 0 })
	require.Equal(expected, result)
}

func TestFilterWithStrings(t *testing.T) {
	require := require.New(t)
	expected := []string{"bat", "cat", "hat"}
	sl := []string{"bat", "monkey", "cat", "submarine", "me", "hat", "lunch"}
	result := Filter(sl, func(s string) bool { return len(s) == 3 })
	require.Equal(expected, result)
}

func TestMapWithInts(t *testing.T) {
	require := require.New(t)
	expected := []int{4, 6, 8, 10}
	sl := []int{2, 3, 4, 5}
	result := Map(sl, func(i int) int { return i * 2 })
	require.Equal(expected, result)
}

func TestMapWithStrings(t *testing.T) {
	require := require.New(t)
	expected := []string{"bat", "cat", "hat"}
	sl := []string{"BAT", "CaT", "haT"}
	result := Map(sl, strings.ToLower)
	require.Equal(expected, result)
}

func TestMapWithStringsToInts(t *testing.T) {
	require := require.New(t)
	expected := []int{1, 3, 5, 7, 9, 11}
	sl := []string{"1", "3", "5", "7", "9", "11"}
	result := Map(sl, func(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		return i
	})
	require.Equal(expected, result)
}
