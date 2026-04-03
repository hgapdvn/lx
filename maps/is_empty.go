package lxmaps

// IsEmpty checks if a map is empty (has no entries).
// Returns true if the map is nil or has no entries, false otherwise.
//
// Example:
//
//	m := map[string]int{"a": 1}
//	IsEmpty(m)
//	// false
//
//	m2 := map[string]int{}
//	IsEmpty(m2)
//	// true
//
//	var m3 map[string]int
//	IsEmpty(m3)
//	// true
func IsEmpty[K comparable, V any](m map[K]V) bool {
	return len(m) == 0
}
