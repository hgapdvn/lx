package lxmaps

import "github.com/hgapdvn/lx/types"

// Entries returns a slice of key-value pairs from a map.
// If the map is nil, it returns an empty slice.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := Entries(m)
//	// out: []lxtypes.Pair[string, int]{{First: "a", Second: 1}, {First: "b", Second: 2}, {First: "c", Second: 3}}
//
//	m2 := map[string]int{}
//	out2 := Entries(m2)
//	// out2: []lxtypes.Pair[string, int]{}
//
//	m3 := map[string]int{"a": 1, "b": 2, "c": 3}
//	out3 := Entries(m3, m3)
//	// out3: []lxtypes.Pair[string, int]{{First: "a", Second: 1}, {First: "b", Second: 2}, {First: "c", Second: 3}}
func Entries[K comparable, V any](in ...map[K]V) []lxtypes.Pair[K, V] {
	if len(in) == 0 {
		return nil
	}

	size := 0
	for i := range in {
		size += len(in[i])
	}

	entries := make([]lxtypes.Pair[K, V], 0, size)
	for _, m := range in {
		for k, v := range m {
			entries = append(entries, lxtypes.Pair[K, V]{First: k, Second: v})
		}
	}

	return entries
}
