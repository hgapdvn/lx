package lxcrypto_test

import (
	"bytes"
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
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.run)
	}
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
