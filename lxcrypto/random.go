package lxcrypto

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

// randomCharset is the set of characters used by RandomString.
// It consists of URL-safe alphanumeric characters (A-Z, a-z, 0-9).
const randomCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

var errNonPositiveLength = errors.New("lxcrypto: n must be greater than zero")

// Random returns a slice of n cryptographically secure random bytes.
// Returns an error if n is less than 1 or if the underlying random source fails.
//
// Example:
//
//	b, err := lxcrypto.RandomBytes(16)
//	// b: []byte{...} (16 random bytes), err: nil
func Random(n int) ([]byte, error) {
	if n <= 0 {
		return nil, errNonPositiveLength
	}
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

// RandomString returns a cryptographically secure random string of n characters
// drawn from the URL-safe alphanumeric charset (A-Z, a-z, 0-9).
// Returns an error if n is less than 1 or if the underlying random source fails.
//
// Example:
//
//	s, err := lxcrypto.RandomString(16)
//	// s: "aB3kLmN9pQrS2tUv" (16 random characters), err: nil
func RandomString(n int) (string, error) {
	if n <= 0 {
		return "", errNonPositiveLength
	}

	const charsetLen = byte(len(randomCharset))
	const maxByte = 255 - (256 % int(charsetLen))

	result := make([]byte, n)
	generated := 0

	// Use a buffer to batch random reads.
	buf := make([]byte, n+n/2+8)
	for generated < n {
		if _, err := rand.Read(buf); err != nil {
			return "", err
		}
		for _, b := range buf {
			if int(b) > maxByte {
				continue // discard biased value
			}
			result[generated] = randomCharset[int(b)%int(charsetLen)]
			generated++
			if generated == n {
				break
			}
		}
	}

	return string(result), nil
}

// SecureToken returns a URL-safe, base64-encoded cryptographically secure token
// derived from n random bytes. The resulting string length is ceil(n*4/3),
// padded to a multiple of 4 with '=' characters.
// Returns an error if n is less than 1 or if the underlying random source fails.
//
// Example:
//
//	token, err := lxcrypto.SecureToken(32)
//	// token: "3q2-7w==" (URL-safe base64-encoded 32 random bytes), err: nil
func SecureToken(n int) (string, error) {
	b, err := Random(n)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
