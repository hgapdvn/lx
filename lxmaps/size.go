package lxmaps

// Size returns the number of entries in a map.
// If the map is nil, returns 0.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	size := Size(m)
//	// size: 3
//
//	m2 := map[string]int{}
//	size2 := Size(m2)
//	// size2: 0
//
//	var m3 map[string]int
//	size3 := Size(m3)
//	// size3: 0
func Size[K comparable, V any](m map[K]V) int {
	return len(m)
}
