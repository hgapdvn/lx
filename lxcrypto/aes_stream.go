package lxcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// gcmStreamChunkSize is the plaintext size of each chunk in EncryptGCMStream.
const gcmStreamChunkSize = 64 * 1024 // 64 KB

// maxGCMStreamChunkLen is the maximum allowed encrypted chunk length on the
// wire: nonce (12 B) + gcmStreamChunkSize plaintext bytes + GCM tag (16 B).
// Any chunk header claiming a larger size is treated as a malformed stream to
// prevent memory exhaustion via a crafted input.
const maxGCMStreamChunkLen = gcmStreamChunkSize + 12 + 16 // 65,564 bytes

// errChunkTooLarge is returned by DecryptGCMStream when a chunk length header
// exceeds maxGCMStreamChunkLen.
var errChunkTooLarge = errors.New("lxcrypto: chunk length exceeds maximum allowed size")

// EncryptGCMStream encrypts data from src and writes it to dst using AES-GCM
// in chunks of 64 KB. Each chunk is independently authenticated.
// key must be 16, 24, or 32 bytes (AES-128, AES-192, or AES-256).
//
// Wire format per chunk: [uint32 chunk_len][nonce (12 B)][ciphertext][tag (16 B)]
//
// Example:
//
//	f, _ := os.Open("large.bin")
//	out, _ := os.Create("large.bin.enc")
//	err := lxcrypto.EncryptGCMStream(f, out, key)
func EncryptGCMStream(src io.Reader, dst io.Writer, key []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("lxcrypto: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("lxcrypto: %w", err)
	}
	nonce := make([]byte, aead.NonceSize())
	buf := make([]byte, gcmStreamChunkSize)

	for {
		n, readErr := io.ReadFull(src, buf)
		if readErr != nil && !errors.Is(readErr, io.ErrUnexpectedEOF) && readErr != io.EOF {
			return readErr
		}
		if n == 0 {
			break
		}
		if _, err = rand.Read(nonce); err != nil {
			return fmt.Errorf("lxcrypto: %w", err)
		}
		sealed := aead.Seal(nonce, nonce, buf[:n], nil)
		if err = binary.Write(dst, binary.BigEndian, uint32(len(sealed))); err != nil {
			return fmt.Errorf("lxcrypto: %w", err)
		}
		if _, err = dst.Write(sealed); err != nil {
			return fmt.Errorf("lxcrypto: %w", err)
		}
		if errors.Is(readErr, io.ErrUnexpectedEOF) || readErr == io.EOF {
			break
		}
	}
	return nil
}

// DecryptGCMStream decrypts chunked AES-GCM data written by EncryptGCMStream.
// Returns an error if the key is invalid, any chunk's authentication tag fails,
// or the stream is truncated.
//
// Example:
//
//	enc, _ := os.Open("large.bin.enc")
//	out, _ := os.Create("large.bin")
//	err := lxcrypto.DecryptGCMStream(enc, out, key)
func DecryptGCMStream(src io.Reader, dst io.Writer, key []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("lxcrypto: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("lxcrypto: %w", err)
	}
	nonceSize := aead.NonceSize()

	for {
		var chunkLen uint32
		if err = binary.Read(src, binary.BigEndian, &chunkLen); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("lxcrypto: %w", err)
		}
		if chunkLen > maxGCMStreamChunkLen {
			return errChunkTooLarge
		}
		chunk := make([]byte, chunkLen)
		if _, err = io.ReadFull(src, chunk); err != nil {
			return fmt.Errorf("lxcrypto: %w", err)
		}
		if len(chunk) < nonceSize {
			return errCiphertextTooShort
		}
		nonce, body := chunk[:nonceSize], chunk[nonceSize:]
		plaintext, err := aead.Open(nil, nonce, body, nil)
		if err != nil {
			return fmt.Errorf("lxcrypto: %w", err)
		}
		if _, err = dst.Write(plaintext); err != nil {
			return fmt.Errorf("lxcrypto: %w", err)
		}
	}
	return nil
}
