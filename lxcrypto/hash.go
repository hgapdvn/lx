package lxcrypto

import (
	"crypto/md5"  //nolint:gosec // MD5 is provided for non-security uses (checksums, etc.)
	"crypto/sha1" //nolint:gosec // SHA1 is provided for non-security uses (checksums, etc.)
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
)

// MD5 returns the MD5 checksum of input as a lowercase hexadecimal string.
// input must be a []byte, string, or io.Reader.
// Returns an error only when input is an io.Reader and reading fails,
// or when the input type is unsupported.
//
// Note: MD5 is not suitable for security-sensitive purposes such as password
// hashing. Use it only for checksums or non-security use cases.
//
// Example:
//
//	lxcrypto.MD5([]byte("hello"))        // ("5d41402abc4b2a76b9719d911017c592", nil)
//	lxcrypto.MD5("hello")                // ("5d41402abc4b2a76b9719d911017c592", nil)
//	lxcrypto.MD5(os.Stdin)               // (digest, nil) or ("", err)
func MD5(input any) (string, error) {
	switch v := input.(type) {
	case []byte:
		h := md5.Sum(v) //nolint:gosec
		return hex.EncodeToString(h[:]), nil
	case string:
		h := md5.Sum([]byte(v)) //nolint:gosec
		return hex.EncodeToString(h[:]), nil
	case io.Reader:
		h := md5.New() //nolint:gosec
		if _, err := io.Copy(h, v); err != nil {
			return "", err
		}
		return hex.EncodeToString(h.Sum(nil)), nil
	default:
		return "", fmt.Errorf("lxcrypto: MD5 unsupported input type %T", input)
	}
}

// SHA1 returns the SHA-1 checksum of input as a lowercase hexadecimal string.
// input must be a []byte, string, or io.Reader.
// Returns an error only when input is an io.Reader and reading fails,
// or when the input type is unsupported.
//
// Note: SHA-1 is no longer considered secure against well-funded adversaries.
// Use SHA-256 or SHA-512 for security-sensitive applications.
//
// Example:
//
//	lxcrypto.SHA1([]byte("hello")) // ("aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d", nil)
//	lxcrypto.SHA1("hello")         // ("aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d", nil)
//	lxcrypto.SHA1(os.Stdin)        // (digest, nil) or ("", err)
func SHA1(input any) (string, error) {
	switch v := input.(type) {
	case []byte:
		h := sha1.Sum(v) //nolint:gosec
		return hex.EncodeToString(h[:]), nil
	case string:
		h := sha1.Sum([]byte(v)) //nolint:gosec
		return hex.EncodeToString(h[:]), nil
	case io.Reader:
		h := sha1.New() //nolint:gosec
		if _, err := io.Copy(h, v); err != nil {
			return "", err
		}
		return hex.EncodeToString(h.Sum(nil)), nil
	default:
		return "", fmt.Errorf("lxcrypto: SHA1 unsupported input type %T", input)
	}
}

// SHA256 returns the SHA-256 checksum of input as a lowercase hexadecimal string.
// input must be a []byte, string, or io.Reader.
// Returns an error only when input is an io.Reader and reading fails,
// or when the input type is unsupported.
//
// Example:
//
//	lxcrypto.SHA256([]byte("hello")) // ("2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", nil)
//	lxcrypto.SHA256("hello")         // ("2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", nil)
//	lxcrypto.SHA256(os.Stdin)        // (digest, nil) or ("", err)
func SHA256(input any) (string, error) {
	switch v := input.(type) {
	case []byte:
		h := sha256.Sum256(v)
		return hex.EncodeToString(h[:]), nil
	case string:
		h := sha256.Sum256([]byte(v))
		return hex.EncodeToString(h[:]), nil
	case io.Reader:
		h := sha256.New()
		if _, err := io.Copy(h, v); err != nil {
			return "", err
		}
		return hex.EncodeToString(h.Sum(nil)), nil
	default:
		return "", fmt.Errorf("lxcrypto: SHA256 unsupported input type %T", input)
	}
}

// SHA512 returns the SHA-512 checksum of input as a lowercase hexadecimal string.
// input must be a []byte, string, or io.Reader.
// Returns an error only when input is an io.Reader and reading fails,
// or when the input type is unsupported.
//
// Example:
//
//	lxcrypto.SHA512([]byte("hello")) // ("9b71d224...", nil)
//	lxcrypto.SHA512("hello")         // ("9b71d224...", nil)
//	lxcrypto.SHA512(os.Stdin)        // (digest, nil) or ("", err)
func SHA512(input any) (string, error) {
	switch v := input.(type) {
	case []byte:
		h := sha512.Sum512(v)
		return hex.EncodeToString(h[:]), nil
	case string:
		h := sha512.Sum512([]byte(v))
		return hex.EncodeToString(h[:]), nil
	case io.Reader:
		h := sha512.New()
		if _, err := io.Copy(h, v); err != nil {
			return "", err
		}
		return hex.EncodeToString(h.Sum(nil)), nil
	default:
		return "", fmt.Errorf("lxcrypto: SHA512 unsupported input type %T", input)
	}
}
