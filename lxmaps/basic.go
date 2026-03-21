package lxmaps

import "github.com/nthanhhai2909/lx/lxtypes"

// Keys returns a slice of keys from a map.
func Keys[K comparable, V any](in ...map[K]V) []K {
	if len(in) == 0 {
		return nil
	}
	size := 0
	for i := range in {
		size += len(in[i])
	}

	keys := make([]K, 0, size)
	for _, m := range in {
		for k := range m {
			keys = append(keys, k)
		}
	}
	return keys
}

// UniqKeys returns a slice of unique keys from a map.
func UniqKeys[K comparable, V any](in ...map[K]V) []K {
	if len(in) == 0 {
		return nil
	}

	seen := make(map[K]struct{})
	for _, m := range in {
		for k := range m {
			seen[k] = struct{}{}
		}
	}
	return Keys(seen)
}

// Values returns a slice of values from a map.
func Values[K comparable, V any](in ...map[K]V) []V {
	if len(in) == 0 {
		return nil
	}

	size := 0
	for i := range in {
		size += len(in[i])
	}

	values := make([]V, 0, size)
	for _, m := range in {
		for _, v := range m {
			values = append(values, v)
		}
	}

	return values
}

// Entries returns a slice of key-value pairs from a map.
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
