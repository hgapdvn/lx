package lxmaps

// EqualBy compares two maps by the given equality function.
// If the maps are nil, they are considered equal.
//
// Example:
//
//	EqualBy(map[string]int{"a": 1}, map[string]int{"a": 1}, func(v1, v2 int) bool { return v1 == v2 }) // true
//	EqualBy(map[string]int{"a": 1}, map[string]int{"a": 2}, func(v1, v2 int) bool { return v1 == v2 }) // false
func EqualBy[K comparable, V any](m1, m2 map[K]V, eq func(v1, v2 V) bool) bool {
	if m1 == nil && m2 == nil {
		return true
	}
	if m1 == nil || m2 == nil {
		return false
	}
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		v2, exists := m2[k]
		if !exists {
			return false
		}
		if !eq(v, v2) {
			return false
		}
	}
	return true
}
