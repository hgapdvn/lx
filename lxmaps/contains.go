package lxmaps

import "reflect"

// Contains returns true if the map contains the key and the value is not nil/zero/empty.
// For pointers and interfaces, it checks if the value is not nil.
// For numeric types, it checks if the value is not zero.
// For booleans, it checks if the value is true.
// For strings, it checks if the value is not empty.
// For slices, maps, and channels, it checks if the value is not nil or empty.
func Contains[K comparable, V any](m map[K]V, key K) bool {
	v, ok := m[key]
	if !ok {
		return false
	}

	// Use reflection to check the value
	rv := reflect.ValueOf(v)
	return checkReflectValue(rv)
}

func checkReflectValue(rv reflect.Value) bool {
	// First, check if the value itself is invalid or nil-like
	if !rv.IsValid() {
		return false
	}

	switch rv.Kind() {
	case reflect.Ptr, reflect.Chan, reflect.Func:
		return !rv.IsNil()
	case reflect.Interface:
		// For interface types, check if they contain nil
		if rv.IsNil() {
			return false
		}
		// Recursively check the underlying value
		return checkReflectValue(rv.Elem())
	case reflect.Slice, reflect.Map:
		// For slices and maps, return true if not nil and not empty
		if rv.IsNil() {
			return false
		}
		return rv.Len() > 0
	case reflect.Bool:
		return rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() != 0
	case reflect.String:
		return rv.String() != ""
	default:
		// For other types, return true if the value exists
		return true
	}
}
