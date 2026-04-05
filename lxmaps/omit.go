package lxmaps

// Omit returns a shallow copy of m without the listed keys.
// Keys to omit that are not in m have no effect.
// For nil m, returns nil. When keys is empty, returns the same result as Clone(m).
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	Omit(m, "b") // map[string]int{"a": 1}
func Omit[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	if m == nil {
		return nil
	}
	if len(keys) == 0 {
		return Clone(m)
	}
	skip := make(map[K]struct{}, len(keys))
	for _, k := range keys {
		skip[k] = struct{}{}
	}
	out := make(map[K]V, len(m))
	for k, v := range m {
		if _, drop := skip[k]; !drop {
			out[k] = v
		}
	}
	return out
}
