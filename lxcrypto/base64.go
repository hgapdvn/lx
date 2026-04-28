package lxcrypto

import (
	"encoding/base64"
)

// Base64Encode returns the standard base64 encoding of data.
//
// Example:
//
//	lxcrypto.Base64Encode([]byte("hello")) // "aGVsbG8="
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64EncodeString returns the standard base64 encoding of s.
//
// Example:
//
//	lxcrypto.Base64EncodeString("hello") // "aGVsbG8="
func Base64EncodeString(s string) string {
	return Base64Encode([]byte(s))
}

// Base64Decode decodes a standard base64-encoded string and returns the raw bytes.
// Returns an error if s is not valid base64.
//
// Example:
//
//	b, err := lxcrypto.Base64Decode("aGVsbG8=")
//	// b: []byte("hello"), err: nil
func Base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// Base64DecodeString decodes a standard base64-encoded string and returns the
// result as a string. Returns an error if s is not valid base64.
//
// Example:
//
//	s, err := lxcrypto.Base64DecodeString("aGVsbG8=")
//	// s: "hello", err: nil
func Base64DecodeString(s string) (string, error) {
	b, err := Base64Decode(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Base64URLEncode returns the URL-safe base64 encoding of data (RFC 4648 §5).
//
// Example:
//
//	lxcrypto.Base64URLEncode([]byte("hello")) // "aGVsbG8="
func Base64URLEncode(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

// Base64URLEncodeString returns the URL-safe base64 encoding of s (RFC 4648 §5).
//
// Example:
//
//	lxcrypto.Base64URLEncodeString("hello") // "aGVsbG8="
func Base64URLEncodeString(s string) string {
	return Base64URLEncode([]byte(s))
}

// Base64URLDecode decodes a URL-safe base64-encoded string and returns the raw bytes.
// Returns an error if s is not valid URL-safe base64.
//
// Example:
//
//	b, err := lxcrypto.Base64URLDecode("aGVsbG8=")
//	// b: []byte("hello"), err: nil
func Base64URLDecode(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}

// Base64URLDecodeString decodes a URL-safe base64-encoded string and returns
// the result as a string. Returns an error if s is not valid URL-safe base64.
//
// Example:
//
//	s, err := lxcrypto.Base64URLDecodeString("aGVsbG8=")
//	// s: "hello", err: nil
func Base64URLDecodeString(s string) (string, error) {
	b, err := Base64URLDecode(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
