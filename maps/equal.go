package lxmaps

// Equal reports whether two maps contain the same keys and values.
// Values are compared with ==, so K and V must be comparable.
//
// Two nil maps are equal. A nil map is not equal to an empty non-nil map.
//
// Example:
//
//	Equal(map[string]int{"a": 1}, map[string]int{"a": 1})     // true
//	Equal(map[string]int{"a": 1}, map[string]int{"a": 2})     // false
//	Equal[int, int](nil, nil)                                 // true
//	Equal(map[string]int(nil), map[string]int{})             // false
func Equal[K comparable, V comparable](m1, m2 map[K]V) bool {
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
		if v2, ok := m2[k]; !ok || v2 != v {
			return false
		}
	}
	return true
}
