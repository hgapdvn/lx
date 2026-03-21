package lxmaps

import "github.com/nthanhhai2909/lx/lxtypes"

// Keys returns a slice of keys from a map.
func Keys[K comparable, V any](m map[K]V) []K {
	if m == nil {
		return nil
	}
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice of values from a map.
func Values[K comparable, V any](m map[K]V) []V {
	if m == nil {
		return nil
	}
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Entries returns a slice of key-value pairs from a map.
func Entries[K comparable, V any](m map[K]V) []lxtypes.Pair[K, V] {
	if m == nil {
		return nil
	}
	entries := make([]lxtypes.Pair[K, V], 0, len(m))
	for k, v := range m {
		entries = append(entries, lxtypes.Pair[K, V]{First: k, Second: v})
	}
	return entries
}
