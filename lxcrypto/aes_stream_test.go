package lxcrypto_test

import (
	"bytes"
	"encoding/binary"
	"errors"
	"testing"

	"github.com/hgapdvn/lx/lxcrypto"
)

// streamErrReader always returns the given error on Read.
type streamErrReader struct{ err error }

func (r streamErrReader) Read(_ []byte) (int, error) { return 0, r.err }

// errWriterAfterN fails on the (n+1)-th Write call.
type errWriterAfterN struct {
	n   int
	err error
	buf bytes.Buffer
}

func (w *errWriterAfterN) Write(p []byte) (int, error) {
	if w.n <= 0 {
		return 0, w.err
	}
	w.n--
	return w.buf.Write(p)
}

// shortWriter accepts only part of each write without returning an error.
// It exercises callers that correctly retry short writes.
type shortWriter struct {
	limit int
	buf   bytes.Buffer
}

func (w *shortWriter) Write(p []byte) (int, error) {
	if len(p) > w.limit {
		p = p[:w.limit]
	}
	return w.buf.Write(p)
}

// TestGCMStream covers EncryptGCMStream and DecryptGCMStream.
func TestGCMStream(t *testing.T) {
	errSentinel := errors.New("test error")

	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		// --- round-trip cases ---
		{name: "small data", run: func(t *testing.T) { roundTrip(t, []byte("hello world"), aesKey256) }},
		{name: "empty data", run: func(t *testing.T) { roundTrip(t, []byte{}, aesKey256) }},
		{name: "exactly one chunk (64KB)", run: func(t *testing.T) {
			roundTrip(t, bytes.Repeat([]byte("a"), 64*1024), aesKey256)
		}},
		{name: "multi-chunk (130KB)", run: func(t *testing.T) {
			roundTrip(t, bytes.Repeat([]byte("b"), 130*1024), aesKey256)
		}},
		{name: "AES-128 key", run: func(t *testing.T) { roundTrip(t, []byte("hello"), aesKey128) }},
		// --- EncryptGCMStream error paths ---
		{
			name: "encrypt: invalid key",
			run: func(t *testing.T) {
				err := lxcrypto.EncryptGCMStream(bytes.NewReader([]byte("hello")), &bytes.Buffer{}, aesKeyBad)
				if err == nil {
					t.Error("expected error, got nil")
				}
			},
		},
		{
			name: "encrypt: reader error",
			run: func(t *testing.T) {
				// non-EOF/ErrUnexpectedEOF error from src triggers the return readErr branch.
				err := lxcrypto.EncryptGCMStream(streamErrReader{err: errSentinel}, &bytes.Buffer{}, aesKey256)
				if err == nil {
					t.Error("expected error, got nil")
				}
			},
		},
		{
			name: "encrypt: dst fails on length prefix",
			run: func(t *testing.T) {
				dst := &errWriterAfterN{n: 0, err: errSentinel}
				err := lxcrypto.EncryptGCMStream(bytes.NewReader([]byte("hello")), dst, aesKey256)
				if err == nil {
					t.Error("expected error, got nil")
				}
			},
		},
		{
			name: "encrypt: dst fails on sealed data",
			run: func(t *testing.T) {
				// Succeeds on first Write (length prefix), fails on second (sealed bytes).
				dst := &errWriterAfterN{n: 1, err: errSentinel}
				err := lxcrypto.EncryptGCMStream(bytes.NewReader([]byte("hello")), dst, aesKey256)
				if err == nil {
					t.Error("expected error, got nil")
				}
			},
		},
		// --- DecryptGCMStream error paths ---
		{
			name: "decrypt: invalid key",
			run: func(t *testing.T) {
				var buf bytes.Buffer
				if err := lxcrypto.EncryptGCMStream(bytes.NewReader([]byte("hello")), &buf, aesKey256); err != nil {
					t.Fatalf("encrypt: %v", err)
				}
				if err := lxcrypto.DecryptGCMStream(&buf, &bytes.Buffer{}, aesKeyBad); err == nil {
					t.Error("expected error, got nil")
				}
			},
		},
		{
			name: "decrypt: partial header (non-EOF binary.Read error)",
			run: func(t *testing.T) {
				// 3 bytes — binary.Read expects 4 — returns io.ErrUnexpectedEOF ≠ io.EOF.
				if err := lxcrypto.DecryptGCMStream(bytes.NewReader([]byte{0, 0, 0}), &bytes.Buffer{}, aesKey256); err == nil {
					t.Error("expected error, got nil")
				}
			},
		},
		{
			name: "decrypt: truncated chunk body",
			run: func(t *testing.T) {
				// Length prefix = 50, only 3 bytes of body → io.ReadFull fails.
				if err := lxcrypto.DecryptGCMStream(bytes.NewReader([]byte{0, 0, 0, 50, 1, 2, 3}), &bytes.Buffer{}, aesKey256); err == nil {
					t.Error("expected error, got nil")
				}
			},
		},
		{
			name: "decrypt: chunk smaller than nonce size",
			run: func(t *testing.T) {
				// Length prefix = 5, body = 5 bytes < nonce size (12) → errCiphertextTooShort.
				if err := lxcrypto.DecryptGCMStream(bytes.NewReader([]byte{0, 0, 0, 5, 1, 2, 3, 4, 5}), &bytes.Buffer{}, aesKey256); err == nil {
					t.Error("expected error, got nil")
				}
			},
		},
		{
			name: "decrypt: tampered ciphertext",
			run: func(t *testing.T) {
				var buf bytes.Buffer
				if err := lxcrypto.EncryptGCMStream(bytes.NewReader([]byte("hello")), &buf, aesKey256); err != nil {
					t.Fatalf("encrypt: %v", err)
				}
				b := buf.Bytes()
				b[len(b)-1] ^= 0xff
				if err := lxcrypto.DecryptGCMStream(bytes.NewReader(b), &bytes.Buffer{}, aesKey256); err == nil {
					t.Error("expected error, got nil")
				}
			},
		},
		{
			name: "decrypt: dst write fails",
			run: func(t *testing.T) {
				var buf bytes.Buffer
				if err := lxcrypto.EncryptGCMStream(bytes.NewReader([]byte("hello")), &buf, aesKey256); err != nil {
					t.Fatalf("encrypt: %v", err)
				}
				dst := &errWriterAfterN{n: 0, err: errSentinel}
				if err := lxcrypto.DecryptGCMStream(&buf, dst, aesKey256); err == nil {
					t.Error("expected error, got nil")
				}
			},
		},
		{
			name: "decrypt: chunk length exceeds maximum",
			run: func(t *testing.T) {
				// Craft a stream with a chunk-length header of 0xFFFFFFFF (4 GiB).
				// DecryptGCMStream must reject this before allocating any memory.
				var buf bytes.Buffer
				// Write uint32 big-endian 0xFFFFFFFF
				buf.Write([]byte{0xFF, 0xFF, 0xFF, 0xFF})
				if err := lxcrypto.DecryptGCMStream(&buf, &bytes.Buffer{}, aesKey256); err == nil {
					t.Error("expected error for oversized chunk, got nil")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.run)
	}
}

func TestGCMStreamIntegrity(t *testing.T) {
	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		{
			name: "round trip multi-chunk data",
			run: func(t *testing.T) {
				plaintext := bytes.Repeat([]byte("a"), 2*64*1024+23)
				ciphertext := encryptGCMStream(t, plaintext)

				var output bytes.Buffer
				if err := lxcrypto.DecryptGCMStream(bytes.NewReader(ciphertext), &output, aesKey256); err != nil {
					t.Fatalf("DecryptGCMStream() error = %v", err)
				}
				if !bytes.Equal(output.Bytes(), plaintext) {
					t.Error("DecryptGCMStream() output does not match plaintext")
				}
			},
		},
		{
			name: "round trip empty data",
			run: func(t *testing.T) {
				ciphertext := encryptGCMStream(t, nil)
				var output bytes.Buffer
				if err := lxcrypto.DecryptGCMStream(bytes.NewReader(ciphertext), &output, aesKey256); err != nil {
					t.Fatalf("DecryptGCMStream() error = %v", err)
				}
				if output.Len() != 0 {
					t.Errorf("DecryptGCMStream() output length = %d, want 0", output.Len())
				}
			},
		},
		{
			name: "encrypt rejects invalid key",
			run: func(t *testing.T) {
				if err := lxcrypto.EncryptGCMStream(bytes.NewReader([]byte("hello")), &bytes.Buffer{}, aesKeyBad); err == nil {
					t.Error("EncryptGCMStream() error = nil for an invalid key")
				}
			},
		},
		{
			name: "decrypt rejects invalid key",
			run: func(t *testing.T) {
				ciphertext := encryptGCMStream(t, []byte("hello"))
				if err := lxcrypto.DecryptGCMStream(bytes.NewReader(ciphertext), &bytes.Buffer{}, aesKeyBad); err == nil {
					t.Error("DecryptGCMStream() error = nil for an invalid key")
				}
			},
		},
		{
			name: "rejects invalid header",
			run: func(t *testing.T) {
				invalidHeader := make([]byte, gcmStreamHeaderSize)
				if err := lxcrypto.DecryptGCMStream(bytes.NewReader(invalidHeader), &bytes.Buffer{}, aesKey256); err == nil {
					t.Error("DecryptGCMStream() error = nil for an invalid header")
				}
			},
		},
		{
			name: "rejects a cleanly truncated final record",
			run: func(t *testing.T) {
				ciphertext := encryptGCMStream(t, bytes.Repeat([]byte("a"), 64*1024+1))
				records := splitGCMStreamRecords(t, ciphertext)
				if len(records) < 2 {
					t.Fatal("expected at least one data record and one final record")
				}

				truncated := append([]byte(nil), ciphertext[:gcmStreamHeaderSize]...)
				for _, record := range records[:len(records)-1] {
					truncated = append(truncated, record...)
				}
				if err := lxcrypto.DecryptGCMStream(bytes.NewReader(truncated), &bytes.Buffer{}, aesKey256); err == nil {
					t.Error("DecryptGCMStream() error = nil for a stream without its final record")
				}
			},
		},
		{
			name: "rejects reordered data records",
			run: func(t *testing.T) {
				ciphertext := encryptGCMStream(t, bytes.Repeat([]byte("a"), 2*64*1024+1))
				records := splitGCMStreamRecords(t, ciphertext)
				if len(records) < 3 {
					t.Fatal("expected at least two data records and one final record")
				}

				reordered := append([]byte(nil), ciphertext[:gcmStreamHeaderSize]...)
				reordered = append(reordered, records[1]...)
				reordered = append(reordered, records[0]...)
				for _, record := range records[2:] {
					reordered = append(reordered, record...)
				}
				if err := lxcrypto.DecryptGCMStream(bytes.NewReader(reordered), &bytes.Buffer{}, aesKey256); err == nil {
					t.Error("DecryptGCMStream() error = nil for reordered data records")
				}
			},
		},
		{
			name: "rejects a record copied from another stream",
			run: func(t *testing.T) {
				plaintext := bytes.Repeat([]byte("a"), 64*1024+1)
				firstCiphertext := encryptGCMStream(t, plaintext)
				secondCiphertext := encryptGCMStream(t, plaintext)
				firstRecords := splitGCMStreamRecords(t, firstCiphertext)
				secondRecords := splitGCMStreamRecords(t, secondCiphertext)
				if len(firstRecords) < 2 || len(secondRecords) < 1 {
					t.Fatal("expected encrypted data records")
				}

				mixed := append([]byte(nil), firstCiphertext[:gcmStreamHeaderSize]...)
				mixed = append(mixed, secondRecords[0]...)
				for _, record := range firstRecords[1:] {
					mixed = append(mixed, record...)
				}
				if err := lxcrypto.DecryptGCMStream(bytes.NewReader(mixed), &bytes.Buffer{}, aesKey256); err == nil {
					t.Error("DecryptGCMStream() error = nil for a record from another stream")
				}
			},
		},
		{
			name: "rejects trailing data after final record",
			run: func(t *testing.T) {
				ciphertext := append(encryptGCMStream(t, []byte("hello")), 1)
				if err := lxcrypto.DecryptGCMStream(bytes.NewReader(ciphertext), &bytes.Buffer{}, aesKey256); err == nil {
					t.Error("DecryptGCMStream() error = nil for trailing data")
				}
			},
		},
		{
			name: "rejects oversized record length before allocating",
			run: func(t *testing.T) {
				ciphertext := append([]byte("LXGCM1"), make([]byte, 16)...)
				ciphertext = append(ciphertext, 1, 0, 1, 0, 29)
				if err := lxcrypto.DecryptGCMStream(bytes.NewReader(ciphertext), &bytes.Buffer{}, aesKey256); err == nil {
					t.Error("DecryptGCMStream() error = nil for an oversized record length")
				}
			},
		},
		{
			name: "supports short writes",
			run: func(t *testing.T) {
				plaintext := bytes.Repeat([]byte("a"), 64*1024+1)
				encrypted := &shortWriter{limit: 1}
				if err := lxcrypto.EncryptGCMStream(bytes.NewReader(plaintext), encrypted, aesKey256); err != nil {
					t.Fatalf("EncryptGCMStream() error = %v", err)
				}

				decrypted := &shortWriter{limit: 1}
				if err := lxcrypto.DecryptGCMStream(bytes.NewReader(encrypted.buf.Bytes()), decrypted, aesKey256); err != nil {
					t.Fatalf("DecryptGCMStream() error = %v", err)
				}
				if !bytes.Equal(decrypted.buf.Bytes(), plaintext) {
					t.Error("DecryptGCMStream() did not write the complete plaintext")
				}
			},
		},
		{
			name: "propagates destination write errors",
			run: func(t *testing.T) {
				writeErr := errors.New("write error")
				if err := lxcrypto.EncryptGCMStream(bytes.NewReader([]byte("hello")), &errWriterAfterN{n: 0, err: writeErr}, aesKey256); err == nil {
					t.Error("EncryptGCMStream() error = nil for a failing destination")
				}

				ciphertext := encryptGCMStream(t, []byte("hello"))
				if err := lxcrypto.DecryptGCMStream(bytes.NewReader(ciphertext), &errWriterAfterN{n: 0, err: writeErr}, aesKey256); err == nil {
					t.Error("DecryptGCMStream() error = nil for a failing destination")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.run)
	}
}

const gcmStreamHeaderSize = len("LXGCM1") + 16

func encryptGCMStream(t *testing.T, plaintext []byte) []byte {
	t.Helper()
	var ciphertext bytes.Buffer
	if err := lxcrypto.EncryptGCMStream(bytes.NewReader(plaintext), &ciphertext, aesKey256); err != nil {
		t.Fatalf("EncryptGCMStream() error = %v", err)
	}
	return ciphertext.Bytes()
}

func splitGCMStreamRecords(t *testing.T, ciphertext []byte) [][]byte {
	t.Helper()
	if len(ciphertext) < gcmStreamHeaderSize {
		t.Fatal("ciphertext is missing the stream header")
	}

	var records [][]byte
	for offset := gcmStreamHeaderSize; offset < len(ciphertext); {
		if len(ciphertext)-offset < 5 {
			t.Fatal("ciphertext has a truncated stream record header")
		}
		payloadLen := int(binary.BigEndian.Uint32(ciphertext[offset+1 : offset+5]))
		end := offset + 5 + payloadLen
		if end > len(ciphertext) {
			t.Fatal("ciphertext has a truncated stream record payload")
		}
		records = append(records, append([]byte(nil), ciphertext[offset:end]...))
		offset = end
	}
	return records
}

// roundTrip encrypts then decrypts and verifies the result matches the original plaintext.
func roundTrip(t *testing.T, plaintext, key []byte) {
	t.Helper()
	var buf bytes.Buffer
	if err := lxcrypto.EncryptGCMStream(bytes.NewReader(plaintext), &buf, key); err != nil {
		t.Fatalf("EncryptGCMStream: %v", err)
	}
	var out bytes.Buffer
	if err := lxcrypto.DecryptGCMStream(&buf, &out, key); err != nil {
		t.Fatalf("DecryptGCMStream: %v", err)
	}
	if !bytes.Equal(out.Bytes(), plaintext) {
		t.Errorf("round-trip mismatch: got %d bytes, want %d bytes", out.Len(), len(plaintext))
	}
}
