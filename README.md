# funcy

Experimenting with golang generics to implement functional favorites like `filter`, `map`, &amp;&amp; `reduce`. 

## Notes

### Compile error, as expected, when you try something that doesn't make sense.

For example, a call to `strings.ToLower` on ints:

```go
func TestMapBadFuncType(t *testing.T) {
	require := require.New(t)
	expected := []string{"bat", "cat", "hat"}
 	sl := []int{1, 2, 3, 4}
 	result := Map(sl, strings.ToLower)
 	require.Equal(expected, result)
}
```

gets this when you test:

```
# github.com/bfollek/funcy/pkg [github.com/bfollek/funcy/pkg.test]
pkg/funcy_test.go:64:20: type func(s string) string of strings.ToLower does not match inferred type func(int) T2 for func(T1) T2
FAIL	github.com/bfollek/funcy/pkg [build failed]
FAIL
```

## Todo


