package lxmaps

// MapKeys transforms the keys of a map using a transformation function.
// Returns a new map with transformed keys and original values.
// If the input map is nil, returns nil.
// If the transformation function produces duplicate keys, the last value wins.
//
// Example:
//
//	m := map[int]string{1: "a", 2: "b", 3: "c"}
//	result := MapKeys(m, func(k int) string {
//		return string(rune(k + 64)) // 1->'A', 2->'B', 3->'C'
//	})
//	// result: map[string]string{"A": "a", "B": "b", "C": "c"}
func MapKeys[K, J comparable, V any](m map[K]V, fn func(K) J) map[J]V {
	if m == nil {
		return nil
	}

	result := make(map[J]V, len(m))
	for k, v := range m {
		newKey := fn(k)
		result[newKey] = v
	}
	return result
}
