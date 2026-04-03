package lxmaps

// Reduce applies the given reducer function to each key-value pair in the map,
// accumulating a result of type R, starting with the initial value.
// If the map is nil, returns the initial value.
// The order of iteration over map entries is not guaranteed.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	sum := Reduce(m, func(acc int, k string, v int) int {
//		return acc + v
//	}, 0)
//	// sum: 6
func Reduce[K comparable, V, R any](m map[K]V, fn func(R, K, V) R, initial R) R {
	result := initial
	for k, v := range m {
		result = fn(result, k, v)
	}
	return result
}
