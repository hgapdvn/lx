package lxmaps

import "github.com/hgapdvn/lx/types"

// FromEntries builds a map from a slice of key-value pairs.
// Pair.First is the key and Pair.Second is the value.
//
// For nil entries, returns nil. For an empty non-nil slice, returns a new empty map.
// If the same key appears more than once, the later pair wins (slice order).
//
// Example:
//
//	pairs := []lxtypes.Pair[string, int]{
//	    {First: "a", Second: 1},
//	    {First: "b", Second: 2},
//	}
//	FromEntries(pairs) // map[string]int{"a": 1, "b": 2}
func FromEntries[K comparable, V any](entries []lxtypes.Pair[K, V]) map[K]V {
	if entries == nil {
		return nil
	}
	out := make(map[K]V, len(entries))
	for _, e := range entries {
		out[e.First] = e.Second
	}
	return out
}
