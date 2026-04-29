package lxcrypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
)

// HMAC256 returns the HMAC-SHA-256 of data using key as a lowercase hexadecimal string.
//
// Example:
//
//	tag := lxcrypto.HMAC256([]byte("message"), []byte("secret"))
func HMAC256(data, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// HMAC256String returns the HMAC-SHA-256 of data using key as a lowercase hexadecimal string.
// Both data and key are strings.
//
// Example:
//
//	tag := lxcrypto.HMAC256String("message", "secret")
func HMAC256String(data, key string) string {
	return HMAC256([]byte(data), []byte(key))
}

// HMAC256Stream returns the HMAC-SHA-256 of all bytes read from src using key
// as a lowercase hexadecimal string. Returns an error if reading from src fails.
//
// Example:
//
//	tag, err := lxcrypto.HMAC256Stream(file, []byte("secret"))
func HMAC256Stream(src io.Reader, key []byte) (string, error) {
	h := hmac.New(sha256.New, key)
	if _, err := io.Copy(h, src); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// VerifyHMAC256 reports whether tag equals the HMAC-SHA-256 of data under key.
// The comparison is performed in constant time to prevent timing side-channel attacks.
//
// Example:
//
//	ok := lxcrypto.VerifyHMAC256([]byte("message"), []byte("secret"), tag)
func VerifyHMAC256(data, key []byte, tag string) bool {
	expected := HMAC256(data, key)
	return hmac.Equal([]byte(expected), []byte(tag))
}

// HMAC512 returns the HMAC-SHA-512 of data using key as a lowercase hexadecimal string.
//
// Example:
//
//	tag := lxcrypto.HMAC512([]byte("message"), []byte("secret"))
func HMAC512(data, key []byte) string {
	h := hmac.New(sha512.New, key)
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// HMAC512String returns the HMAC-SHA-512 of data using key as a lowercase hexadecimal string.
// Both data and key are strings.
//
// Example:
//
//	tag := lxcrypto.HMAC512String("message", "secret")
func HMAC512String(data, key string) string {
	return HMAC512([]byte(data), []byte(key))
}

// HMAC512Stream returns the HMAC-SHA-512 of all bytes read from src using key
// as a lowercase hexadecimal string. Returns an error if reading from src fails.
//
// Example:
//
//	tag, err := lxcrypto.HMAC512Stream(file, []byte("secret"))
func HMAC512Stream(src io.Reader, key []byte) (string, error) {
	h := hmac.New(sha512.New, key)
	if _, err := io.Copy(h, src); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// VerifyHMAC512 reports whether tag equals the HMAC-SHA-512 of data under key.
// The comparison is performed in constant time to prevent timing side-channel attacks.
//
// Example:
//
//	ok := lxcrypto.VerifyHMAC512([]byte("message"), []byte("secret"), tag)
func VerifyHMAC512(data, key []byte, tag string) bool {
	expected := HMAC512(data, key)
	return hmac.Equal([]byte(expected), []byte(tag))
}
