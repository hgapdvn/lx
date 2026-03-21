package lxptrs

// Ref returns the address of the given value.
func Ref[V any](v V) *V {
	return &v
}

// Deref returns the value pointed to by the given pointer.
func Deref[V any](v *V) V {
	return *v
}
