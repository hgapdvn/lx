package lxmaps

// Clone returns a shallow copy of the map.
// Keys and values are copied by assignment; pointer, slice, and map values
// are not deep-copied.
// For nil input, returns nil. For an empty non-nil map, returns a new empty map.
//
// Example:
//
//	orig := map[string]int{"a": 1, "b": 2}
//	cp := Clone(orig)
//	delete(cp, "a")
//	// orig still has "a"; cp does not
func Clone[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}
	out := make(map[K]V, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}
