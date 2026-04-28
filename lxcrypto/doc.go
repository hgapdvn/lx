// Package lxcrypto provides common cryptographic utilities built on top of Go's
// standard library crypto packages.
//
// The package is organized into the following categories:
//
// 1. Hashing (hash.go)
//   - MD5, MD5String   - MD5 digest (hex-encoded)
//   - SHA1, SHA1String - SHA-1 digest (hex-encoded)
//   - SHA256, SHA256String - SHA-256 digest (hex-encoded)
//   - SHA512, SHA512String - SHA-512 digest (hex-encoded)
//
// All hash functions accept a []byte input and return a lowercase hex-encoded
// string. Convenience *String variants accept a plain string input.
//
// Example:
//
//	sum := lxcrypto.SHA256([]byte("hello"))
//	// sum: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
//
//	sum2 := lxcrypto.SHA256String("hello")
//	// sum2: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
package lxcrypto
