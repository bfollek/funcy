package funcy

import (
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
