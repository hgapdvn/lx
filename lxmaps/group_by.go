package lxmaps

// GroupBy groups map entries by the result of the grouping function.
// Returns a map where each key is the result of applying the grouping function,
// and each value is a map containing the original key-value pairs that produced that group key.
// If the map is nil, returns nil.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	groups := GroupBy(m, func(k string, v int) string {
//		if v%2 == 0 { return "even" }
//		return "odd"
//	})
//	// groups: map[string]map[string]int{
//	//   "odd":  {"a": 1, "c": 3},
//	//   "even": {"b": 2, "d": 4},
//	// }
func GroupBy[K comparable, V any, G comparable](m map[K]V, fn func(K, V) G) map[G]map[K]V {
	if m == nil {
		return nil
	}
	result := make(map[G]map[K]V)
	for k, v := range m {
		group := fn(k, v)
		if result[group] == nil {
			result[group] = make(map[K]V)
		}
		result[group][k] = v
	}
	return result
}
