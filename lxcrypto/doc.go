// Package lxcrypto provides common cryptographic utilities built on top of Go's
// standard library crypto packages.
//
// # Hashing (hash.go)
//
// Four hash functions are available, each accepting a []byte, string, or
// io.Reader as input:
//
//	MD5(input any)    (string, error)
//	SHA1(input any)   (string, error)
//	SHA256(input any) (string, error)
//	SHA512(input any) (string, error)
//
// All functions return a lowercase hex-encoded digest.
// An error is returned only when input is an io.Reader and reading fails,
// or when the input type is not one of the three supported types.
//
// Supported input types:
//
//   - []byte  — in-memory byte slice (never errors)
//   - string  — plain text (never errors)
//   - io.Reader — streamed chunk-by-chunk; suitable for large files or HTTP
//     bodies without loading them into memory
//
// Example ([]byte):
//
//	sum, _ := lxcrypto.SHA256([]byte("hello"))
//	// sum: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
//
// Example (string):
//
//	sum, _ := lxcrypto.SHA256("hello")
//	// sum: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
//
// Example (streaming a file):
//
//	f, err := os.Open("large.bin")
//	if err != nil { ... }
//	defer f.Close()
//	sum, err := lxcrypto.SHA256(f)
//
// Example (streaming an HTTP response body):
//
//	resp, err := http.Get("https://example.com/file")
//	if err != nil { ... }
//	defer resp.Body.Close()
//	sum, err := lxcrypto.SHA256(resp.Body)
package lxcrypto
