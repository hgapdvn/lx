package lxmaps

// Invert returns a new map swapping keys and values: each entry m[k]=v becomes out[v]=k.
// K and V must both be comparable so they can be used as map keys.
//
// For nil input, returns nil. For an empty non-nil map, returns a new empty map.
//
// If several keys in m hold the same value, only one of those keys appears in the
// result (the last assignment during iteration wins; which key that is is unspecified).
//
// Example:
//
//	Invert(map[string]int{"a": 1, "b": 2}) // map[int]string{1: "a", 2: "b"} (key order may vary)
func Invert[K comparable, V comparable](m map[K]V) map[V]K {
	if m == nil {
		return nil
	}
	out := make(map[V]K, len(m))
	for k, v := range m {
		out[v] = k
	}
	return out
}
