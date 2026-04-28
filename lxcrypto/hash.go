package lxcrypto

import (
	"crypto/md5"  //nolint:gosec // MD5 is provided for non-security uses (checksums, etc.)
	"crypto/sha1" //nolint:gosec // SHA1 is provided for non-security uses (checksums, etc.)
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// MD5 returns the MD5 checksum of data as a lowercase hexadecimal string.
//
// Note: MD5 is not suitable for security-sensitive purposes such as password
// hashing. Use it only for checksums or non-security use cases.
//
// Example:
//
//	sum := lxcrypto.MD5([]byte("hello"))
//	// sum: "5d41402abc4b2a76b9719d911017c592"
func MD5(data []byte) string {
	h := md5.Sum(data) //nolint:gosec
	return hex.EncodeToString(h[:])
}

// MD5String returns the MD5 checksum of s as a lowercase hexadecimal string.
//
// Note: MD5 is not suitable for security-sensitive purposes such as password
// hashing. Use it only for checksums or non-security use cases.
//
// Example:
//
//	sum := lxcrypto.MD5String("hello")
//	// sum: "5d41402abc4b2a76b9719d911017c592"
func MD5String(s string) string {
	return MD5([]byte(s))
}

// SHA1 returns the SHA-1 checksum of data as a lowercase hexadecimal string.
//
// Note: SHA-1 is no longer considered secure against well-funded adversaries.
// Use SHA-256 or SHA-512 for security-sensitive applications.
//
// Example:
//
//	sum := lxcrypto.SHA1([]byte("hello"))
//	// sum: "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"
func SHA1(data []byte) string {
	h := sha1.Sum(data) //nolint:gosec
	return hex.EncodeToString(h[:])
}

// SHA1String returns the SHA-1 checksum of s as a lowercase hexadecimal string.
//
// Note: SHA-1 is no longer considered secure against well-funded adversaries.
// Use SHA-256 or SHA-512 for security-sensitive applications.
//
// Example:
//
//	sum := lxcrypto.SHA1String("hello")
//	// sum: "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"
func SHA1String(s string) string {
	return SHA1([]byte(s))
}

// SHA256 returns the SHA-256 checksum of data as a lowercase hexadecimal string.
//
// Example:
//
//	sum := lxcrypto.SHA256([]byte("hello"))
//	// sum: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
func SHA256(data []byte) string {
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:])
}

// SHA256String returns the SHA-256 checksum of s as a lowercase hexadecimal string.
//
// Example:
//
//	sum := lxcrypto.SHA256String("hello")
//	// sum: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
func SHA256String(s string) string {
	return SHA256([]byte(s))
}

// SHA512 returns the SHA-512 checksum of data as a lowercase hexadecimal string.
//
// Example:
//
//	sum := lxcrypto.SHA512([]byte("hello"))
//	// sum: "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"
func SHA512(data []byte) string {
	h := sha512.Sum512(data)
	return hex.EncodeToString(h[:])
}

// SHA512String returns the SHA-512 checksum of s as a lowercase hexadecimal string.
//
// Example:
//
//	sum := lxcrypto.SHA512String("hello")
//	// sum: "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"
func SHA512String(s string) string {
	return SHA512([]byte(s))
}
