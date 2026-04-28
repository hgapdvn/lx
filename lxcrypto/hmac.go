package lxcrypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// HMAC256 returns the HMAC-SHA-256 of data using key as a lowercase hexadecimal
// string. Both data and key may be a []byte or a string; the types are inferred
// at the call site.
//
// Example:
//
//	tag := lxcrypto.HMAC256("message", "secret")
//	tag := lxcrypto.HMAC256([]byte("message"), []byte("secret"))
func HMAC256[D, K BytesOrString](data D, key K) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// HMAC512 returns the HMAC-SHA-512 of data using key as a lowercase hexadecimal
// string. Both data and key may be a []byte or a string; the types are inferred
// at the call site.
//
// Example:
//
//	tag := lxcrypto.HMAC512("message", "secret")
//	tag := lxcrypto.HMAC512([]byte("message"), []byte("secret"))
func HMAC512[D, K BytesOrString](data D, key K) string {
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
