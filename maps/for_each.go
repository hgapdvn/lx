package lxmaps

// ForEach iterates over a map and calls the provided function for each key-value pair.
// The function is called with the key and value as arguments.
//
// Example:
//
//	ForEach(map[string]int{"a": 1, "b": 2}, func(k string, v int) {
//		fmt.Printf("Key: %s, Value: %d\n", k, v)
//	})
// outputs:
//  Key: a, Value: 1
//  Key: b, Value: 2
func ForEach[K comparable, V any](m map[K]V, fn func(k K, v V)) {
	for k, v := range m {
		fn(k, v)
	}
}
