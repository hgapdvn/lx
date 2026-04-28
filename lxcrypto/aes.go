package lxcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
)

var (
	errCiphertextTooShort = errors.New("lxcrypto: ciphertext too short")
	errInvalidBlockSize   = errors.New("lxcrypto: ciphertext length is not a multiple of block size")
	errInvalidPadding     = errors.New("lxcrypto: invalid PKCS7 padding")
)

// EncryptGCM encrypts plaintext using AES-GCM with the given key.
// key must be 16, 24, or 32 bytes (AES-128, AES-192, or AES-256).
// A random 12-byte nonce is generated and prepended to the returned ciphertext:
// [nonce (12 bytes)][ciphertext][auth tag (16 bytes)].
//
// Example:
//
//	ct, err := lxcrypto.EncryptGCM([]byte("hello"), key)
func EncryptGCM(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("lxcrypto: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("lxcrypto: %w", err)
	}
	nonce := make([]byte, aead.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}
	return aead.Seal(nonce, nonce, plaintext, nil), nil
}

// DecryptGCM decrypts AES-GCM ciphertext produced by EncryptGCM.
// The nonce is read from the first 12 bytes. Returns an error if the key is
// invalid, ciphertext is too short, or authentication fails (tampered data).
//
// Example:
//
//	pt, err := lxcrypto.DecryptGCM(ct, key)
//	// pt: []byte("hello")
func DecryptGCM(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("lxcrypto: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("lxcrypto: %w", err)
	}
	nonceSize := aead.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errCiphertextTooShort
	}
	nonce, body := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aead.Open(nil, nonce, body, nil)
	if err != nil {
		return nil, fmt.Errorf("lxcrypto: %w", err)
	}
	return plaintext, nil
}

// EncryptGCMString encrypts plaintext using AES-GCM and returns the result
// as a URL-safe base64-encoded string.
//
// Example:
//
//	s, err := lxcrypto.EncryptGCMString("hello", key)
func EncryptGCMString(plaintext string, key []byte) (string, error) {
	b, err := EncryptGCM([]byte(plaintext), key)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// DecryptGCMString decrypts a URL-safe base64-encoded ciphertext produced by
// EncryptGCMString and returns the original plaintext string.
//
// Example:
//
//	s, err := lxcrypto.DecryptGCMString(encoded, key)
//	// s: "hello"
func DecryptGCMString(ciphertext string, key []byte) (string, error) {
	b, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("lxcrypto: %w", err)
	}
	plaintext, err := DecryptGCM(b, key)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// EncryptCBC encrypts plaintext using AES-CBC with PKCS7 padding.
// key must be 16, 24, or 32 bytes. A random 16-byte IV is prepended to the
// returned ciphertext: [IV (16 bytes)][ciphertext].
//
// Note: CBC does not provide authentication. Use EncryptGCM when tamper
// detection is required.
//
// Example:
//
//	ct, err := lxcrypto.EncryptCBC([]byte("hello"), key)
func EncryptCBC(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("lxcrypto: %w", err)
	}
	padded := pkcs7Pad(plaintext, aes.BlockSize)
	out := make([]byte, aes.BlockSize+len(padded))
	iv := out[:aes.BlockSize]
	if _, err = rand.Read(iv); err != nil {
		return nil, err
	}
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(out[aes.BlockSize:], padded)
	return out, nil
}

// DecryptCBC decrypts AES-CBC ciphertext produced by EncryptCBC.
// The IV is read from the first 16 bytes. Returns an error if the key is
// invalid, the ciphertext length is not a multiple of the block size, or
// PKCS7 padding is corrupt.
//
// Example:
//
//	pt, err := lxcrypto.DecryptCBC(ct, key)
//	// pt: []byte("hello")
func DecryptCBC(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("lxcrypto: %w", err)
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, errCiphertextTooShort
	}
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errInvalidBlockSize
	}
	iv, body := ciphertext[:aes.BlockSize], ciphertext[aes.BlockSize:]
	plaintext := make([]byte, len(body))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(plaintext, body)
	return pkcs7Unpad(plaintext)
}

// EncryptCBCString encrypts plaintext using AES-CBC and returns the result
// as a URL-safe base64-encoded string.
//
// Example:
//
//	s, err := lxcrypto.EncryptCBCString("hello", key)
func EncryptCBCString(plaintext string, key []byte) (string, error) {
	b, err := EncryptCBC([]byte(plaintext), key)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// DecryptCBCString decrypts a URL-safe base64-encoded ciphertext produced by
// EncryptCBCString and returns the original plaintext string.
//
// Example:
//
//	s, err := lxcrypto.DecryptCBCString(encoded, key)
//	// s: "hello"
func DecryptCBCString(ciphertext string, key []byte) (string, error) {
	b, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("lxcrypto: %w", err)
	}
	plaintext, err := DecryptCBC(b, key)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// pkcs7Pad pads src to a multiple of blockSize using PKCS7.
func pkcs7Pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	out := make([]byte, len(src)+padding)
	copy(out, src)
	for i := len(src); i < len(out); i++ {
		out[i] = byte(padding)
	}
	return out
}

// pkcs7Unpad removes PKCS7 padding and returns the unpadded slice.
func pkcs7Unpad(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return nil, errInvalidPadding
	}
	padding := int(src[len(src)-1])
	if padding == 0 || padding > aes.BlockSize || padding > len(src) {
		return nil, errInvalidPadding
	}
	for _, b := range src[len(src)-padding:] {
		if int(b) != padding {
			return nil, errInvalidPadding
		}
	}
	return src[:len(src)-padding], nil
}
