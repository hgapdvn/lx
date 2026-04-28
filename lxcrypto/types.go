package lxcrypto

// BytesOrString is the type constraint for functions that accept either
// a []byte or a string as input.
type BytesOrString interface {
	~[]byte | ~string
}
