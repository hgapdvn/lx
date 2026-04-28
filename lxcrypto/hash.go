package lxcrypto

import (
	"crypto/md5"  //nolint:gosec // MD5 is provided for non-security uses (checksums, etc.)
	"crypto/sha1" //nolint:gosec // SHA1 is provided for non-security uses (checksums, etc.)
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
)

// MD5 returns the MD5 checksum of data as a lowercase hexadecimal string.
//
// Note: MD5 is not suitable for security-sensitive purposes such as password
// hashing. Use it only for checksums or non-security use cases.
//
// Example:
//
//	lxcrypto.MD5([]byte("hello")) // "5d41402abc4b2a76b9719d911017c592"
func MD5(data []byte) string {
	h := md5.Sum(data) //nolint:gosec
	return hex.EncodeToString(h[:])
}

// MD5String returns the MD5 checksum of s as a lowercase hexadecimal string.
//
// Example:
//
//	lxcrypto.MD5String("hello") // "5d41402abc4b2a76b9719d911017c592"
func MD5String(s string) string {
	return MD5([]byte(s))
}

// MD5Stream returns the MD5 checksum of all bytes read from src as a lowercase
// hexadecimal string. Returns an error if reading from src fails.
//
// Example:
//
//	digest, err := lxcrypto.MD5Stream(os.Stdin)
func MD5Stream(src io.Reader) (string, error) {
	h := md5.New() //nolint:gosec
	if _, err := io.Copy(h, src); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// SHA1 returns the SHA-1 checksum of data as a lowercase hexadecimal string.
//
// Note: SHA-1 is no longer considered secure against well-funded adversaries.
// Use SHA-256 or SHA-512 for security-sensitive applications.
//
// Example:
//
//	lxcrypto.SHA1([]byte("hello")) // "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"
func SHA1(data []byte) string {
	h := sha1.Sum(data) //nolint:gosec
	return hex.EncodeToString(h[:])
}

// SHA1String returns the SHA-1 checksum of s as a lowercase hexadecimal string.
//
// Example:
//
//	lxcrypto.SHA1String("hello") // "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"
func SHA1String(s string) string {
	return SHA1([]byte(s))
}

// SHA1Stream returns the SHA-1 checksum of all bytes read from src as a
// lowercase hexadecimal string. Returns an error if reading from src fails.
//
// Example:
//
//	digest, err := lxcrypto.SHA1Stream(os.Stdin)
func SHA1Stream(src io.Reader) (string, error) {
	h := sha1.New() //nolint:gosec
	if _, err := io.Copy(h, src); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// SHA256 returns the SHA-256 checksum of data as a lowercase hexadecimal string.
//
// Example:
//
//	lxcrypto.SHA256([]byte("hello")) // "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
func SHA256(data []byte) string {
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:])
}

// SHA256String returns the SHA-256 checksum of s as a lowercase hexadecimal string.
//
// Example:
//
//	lxcrypto.SHA256String("hello") // "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
func SHA256String(s string) string {
	return SHA256([]byte(s))
}

// SHA256Stream returns the SHA-256 checksum of all bytes read from src as a
// lowercase hexadecimal string. Returns an error if reading from src fails.
//
// Example:
//
//	digest, err := lxcrypto.SHA256Stream(file)
func SHA256Stream(src io.Reader) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, src); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// SHA512 returns the SHA-512 checksum of data as a lowercase hexadecimal string.
//
// Example:
//
//	lxcrypto.SHA512([]byte("hello")) // "9b71d224..."
func SHA512(data []byte) string {
	h := sha512.Sum512(data)
	return hex.EncodeToString(h[:])
}

// SHA512String returns the SHA-512 checksum of s as a lowercase hexadecimal string.
//
// Example:
//
//	lxcrypto.SHA512String("hello") // "9b71d224..."
func SHA512String(s string) string {
	return SHA512([]byte(s))
}

// SHA512Stream returns the SHA-512 checksum of all bytes read from src as a
// lowercase hexadecimal string. Returns an error if reading from src fails.
//
// Example:
//
//	digest, err := lxcrypto.SHA512Stream(file)
func SHA512Stream(src io.Reader) (string, error) {
	h := sha512.New()
	if _, err := io.Copy(h, src); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
