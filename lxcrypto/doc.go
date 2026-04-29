// Package lxcrypto provides common cryptographic utilities built on top of Go's
// standard library crypto packages.
//
// All functions follow a consistent naming convention:
//
//   - Plain name (e.g. MD5, HMAC256)      → accepts []byte
//   - String suffix (e.g. MD5String)       → accepts / returns string
//   - Stream suffix (e.g. MD5Stream)       → accepts io.Reader / io.Writer
//
// # Hashing (hash.go)
//
// Four algorithms are available, each in three variants:
//
//	MD5(data []byte) string
//	MD5String(s string) string
//	MD5Stream(src io.Reader) (string, error)
//
//	SHA1(data []byte) string
//	SHA1String(s string) string
//	SHA1Stream(src io.Reader) (string, error)
//
//	SHA256(data []byte) string
//	SHA256String(s string) string
//	SHA256Stream(src io.Reader) (string, error)
//
//	SHA512(data []byte) string
//	SHA512String(s string) string
//	SHA512Stream(src io.Reader) (string, error)
//
// Bytes and string variants never return an error.
// Stream variants return an error only when reading from src fails.
// All return a lowercase hex-encoded digest.
//
// Example:
//
//	lxcrypto.SHA256([]byte("hello"))  // "2cf24dba..."
//	lxcrypto.SHA256String("hello")    // "2cf24dba..."
//
//	f, _ := os.Open("large.bin")
//	defer f.Close()
//	sum, err := lxcrypto.SHA256Stream(f)
//
// # HMAC (hmac.go)
//
// Two algorithms, each in three variants plus a timing-safe verifier:
//
//	HMAC256(data, key []byte) string
//	HMAC256String(data, key string) string
//	HMAC256Stream(src io.Reader, key []byte) (string, error)
//	VerifyHMAC256(data, key []byte, tag string) bool
//
//	HMAC512(data, key []byte) string
//	HMAC512String(data, key string) string
//	HMAC512Stream(src io.Reader, key []byte) (string, error)
//	VerifyHMAC512(data, key []byte, tag string) bool
//
// Example:
//
//	tag := lxcrypto.HMAC256([]byte("message"), []byte("secret"))
//	tag := lxcrypto.HMAC256String("message", "secret")
//	ok  := lxcrypto.VerifyHMAC256([]byte("message"), []byte("secret"), tag)
//
//	f, _ := os.Open("large.bin")
//	tag, err := lxcrypto.HMAC256Stream(f, []byte("secret"))
//
// # AES Encryption (aes.go)
//
// GCM mode (authenticated):
//
//	EncryptGCM(plaintext, key []byte) ([]byte, error)
//	DecryptGCM(ciphertext, key []byte) ([]byte, error)
//	EncryptGCMString(plaintext string, key []byte) (string, error)
//	DecryptGCMString(ciphertext string, key []byte) (string, error)
//
// CBC mode (PKCS7-padded, no authentication):
//
//	EncryptCBC(plaintext, key []byte) ([]byte, error)
//	DecryptCBC(ciphertext, key []byte) ([]byte, error)
//	EncryptCBCString(plaintext string, key []byte) (string, error)
//	DecryptCBCString(ciphertext string, key []byte) (string, error)
//
// key must be 16, 24, or 32 bytes (AES-128/192/256).
// GCM String variants base64 URL-encode the ciphertext for safe transport.
// Prefer GCM over CBC; CBC provides no tamper detection.
//
// # AES Streaming (aes_stream.go)
//
// For large data without loading it entirely into memory:
//
//	EncryptGCMStream(src io.Reader, dst io.Writer, key []byte) error
//	DecryptGCMStream(src io.Reader, dst io.Writer, key []byte) error
//
// Data is processed in 64 KB chunks; each chunk is independently authenticated.
//
// Example:
//
//	in, _  := os.Open("large.bin")
//	out, _ := os.Create("large.bin.enc")
//	err := lxcrypto.EncryptGCMStream(in, out, key)
//
// # Base64 (base64.go)
//
// Standard and URL-safe encoding, each in three variants:
//
//	Base64Encode(data []byte) string
//	Base64EncodeString(s string) string
//	Base64EncodeStream(src io.Reader, dst io.Writer) error
//
//	Base64Decode(s string) ([]byte, error)
//	Base64DecodeString(s string) (string, error)
//	Base64DecodeStream(src io.Reader, dst io.Writer) error
//
//	Base64URLEncode(data []byte) string
//	Base64URLEncodeString(s string) string
//	Base64URLEncodeStream(src io.Reader, dst io.Writer) error
//
//	Base64URLDecode(s string) ([]byte, error)
//	Base64URLDecodeString(s string) (string, error)
//	Base64URLDecodeStream(src io.Reader, dst io.Writer) error
//
// # Random Generation (random.go)
//
// Three functions generate cryptographically secure random values using
// crypto/rand as the underlying source:
//
//	Random(n int) ([]byte, error)        – n raw random bytes
//	RandomString(n int) (string, error)  – n-char alphanumeric string (A-Z a-z 0-9)
//	SecureToken(n int) (string, error)   – URL-safe base64 token from n random bytes
//
// All return an error when n ≤ 0 or when the OS random source fails.
// RandomString uses rejection sampling to avoid modulo bias.
//
// Example:
//
//	b, _     := lxcrypto.Random(16)       // 16 raw bytes
//	s, _     := lxcrypto.RandomString(24) // "aB3kLmN9pQrS2tUvXyZ01234"
//	token, _ := lxcrypto.SecureToken(32)  // URL-safe base64-encoded string
package lxcrypto
