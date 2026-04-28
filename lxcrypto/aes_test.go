package lxcrypto_test

import (
	"bytes"
	"encoding/base64"
	"testing"

	"github.com/hgapdvn/lx/lxcrypto"
)

var (
	aesKey128 = []byte("0123456789abcdef")                 // 16 bytes
	aesKey192 = []byte("0123456789abcdef01234567")         // 24 bytes
	aesKey256 = []byte("0123456789abcdef0123456789abcdef") // 32 bytes
	aesKeyBad = []byte("short")
)

// TestGCM covers EncryptGCM and DecryptGCM.
func TestGCM(t *testing.T) {
	tests := []struct {
		name           string
		plaintext      []byte
		encryptKey     []byte
		decryptKey     []byte
		wantEncryptErr bool
		wantDecryptErr bool
		check          func(t *testing.T, pt, ct []byte)
	}{
		{
			name:      "AES-128 round-trip",
			plaintext: []byte("hello world"), encryptKey: aesKey128, decryptKey: aesKey128,
		},
		{
			name:      "AES-192 round-trip",
			plaintext: []byte("hello world"), encryptKey: aesKey192, decryptKey: aesKey192,
		},
		{
			name:      "AES-256 round-trip",
			plaintext: []byte("hello world"), encryptKey: aesKey256, decryptKey: aesKey256,
		},
		{
			name:      "empty plaintext",
			plaintext: []byte{}, encryptKey: aesKey256, decryptKey: aesKey256,
		},
		{
			name:      "nil plaintext",
			plaintext: nil, encryptKey: aesKey256, decryptKey: aesKey256,
		},
		{
			name:      "invalid encrypt key",
			plaintext: []byte("hello"), encryptKey: aesKeyBad,
			wantEncryptErr: true,
		},
		{
			name:      "invalid decrypt key",
			plaintext: []byte("hello"), encryptKey: aesKey256, decryptKey: aesKeyBad,
			wantDecryptErr: true,
		},
		{
			name:      "ciphertext too short",
			plaintext: []byte("hello"), encryptKey: aesKey256, decryptKey: aesKey256,
			wantDecryptErr: true,
			check: func(t *testing.T, _, _ []byte) {
				_, err := lxcrypto.DecryptGCM([]byte("short"), aesKey256)
				if err == nil {
					t.Error("expected error for short ciphertext, got nil")
				}
			},
		},
		{
			name:      "tampered ciphertext",
			plaintext: []byte("hello"), encryptKey: aesKey256, decryptKey: aesKey256,
			check: func(t *testing.T, pt, ct []byte) {
				tampered := make([]byte, len(ct))
				copy(tampered, ct)
				tampered[len(tampered)-1] ^= 0xff
				_, err := lxcrypto.DecryptGCM(tampered, aesKey256)
				if err == nil {
					t.Error("expected error for tampered ciphertext, got nil")
				}
			},
		},
		{
			name:      "uniqueness",
			plaintext: []byte("same plaintext"), encryptKey: aesKey256, decryptKey: aesKey256,
			check: func(t *testing.T, pt, ct []byte) {
				ct2, err := lxcrypto.EncryptGCM(pt, aesKey256)
				if err != nil {
					t.Fatalf("second EncryptGCM failed: %v", err)
				}
				if bytes.Equal(ct, ct2) {
					t.Error("two encryptions returned identical ciphertext")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ct, err := lxcrypto.EncryptGCM(tt.plaintext, tt.encryptKey)
			if tt.wantEncryptErr {
				if err == nil {
					t.Error("expected encrypt error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("EncryptGCM: %v", err)
			}

			if tt.check != nil {
				tt.check(t, tt.plaintext, ct)
			}
			if tt.wantDecryptErr || tt.decryptKey == nil {
				return
			}

			got, err := lxcrypto.DecryptGCM(ct, tt.decryptKey)
			if tt.wantDecryptErr {
				if err == nil {
					t.Error("expected decrypt error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("DecryptGCM: %v", err)
			}
			want := tt.plaintext
			if want == nil {
				want = []byte{}
			}
			if !bytes.Equal(got, want) {
				t.Errorf("round-trip mismatch: got %q, want %q", got, want)
			}
		})
	}
}

// TestGCMString covers EncryptGCMString and DecryptGCMString.
func TestGCMString(t *testing.T) {
	tests := []struct {
		name           string
		plaintext      string
		key            []byte
		wantErr        bool
		decryptInput   string // override ciphertext for decrypt-only error cases
		wantDecryptErr bool
	}{
		{name: "round-trip", plaintext: "hello world", key: aesKey256},
		{name: "empty string", plaintext: "", key: aesKey256},
		{name: "invalid key", plaintext: "hello", key: aesKeyBad, wantErr: true},
		{
			name:           "invalid base64 on decrypt",
			plaintext:      "unused",
			key:            aesKey256,
			decryptInput:   "not!!valid@@base64",
			wantDecryptErr: true,
		},
		{
			name:           "valid base64 but bad GCM data",
			plaintext:      "unused",
			key:            aesKey256,
			decryptInput:   base64.URLEncoding.EncodeToString([]byte("short")),
			wantDecryptErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.decryptInput != "" {
				_, err := lxcrypto.DecryptGCMString(tt.decryptInput, tt.key)
				if !tt.wantDecryptErr {
					if err != nil {
						t.Fatalf("unexpected error: %v", err)
					}
				} else if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}

			ct, err := lxcrypto.EncryptGCMString(tt.plaintext, tt.key)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("EncryptGCMString: %v", err)
			}
			got, err := lxcrypto.DecryptGCMString(ct, tt.key)
			if err != nil {
				t.Fatalf("DecryptGCMString: %v", err)
			}
			if got != tt.plaintext {
				t.Errorf("round-trip mismatch: got %q, want %q", got, tt.plaintext)
			}
		})
	}
}

// TestCBC covers EncryptCBC and DecryptCBC, including all pkcs7Unpad branches.
func TestCBC(t *testing.T) {
	tests := []struct {
		name           string
		plaintext      []byte
		encryptKey     []byte
		decryptKey     []byte
		wantEncryptErr bool
		// mutate is called on the ciphertext before decrypting (nil = no mutation).
		mutate         func(ct []byte)
		wantDecryptErr bool
		// directCT bypasses encrypt and directly calls DecryptCBC (for custom ciphertext cases).
		directCT func() []byte
	}{
		{name: "AES-128 round-trip", plaintext: []byte("hello world"), encryptKey: aesKey128, decryptKey: aesKey128},
		{name: "AES-192 round-trip", plaintext: []byte("hello world"), encryptKey: aesKey192, decryptKey: aesKey192},
		{name: "AES-256 round-trip", plaintext: []byte("hello world"), encryptKey: aesKey256, decryptKey: aesKey256},
		{name: "empty plaintext", plaintext: []byte{}, encryptKey: aesKey256, decryptKey: aesKey256},
		{name: "nil plaintext", plaintext: nil, encryptKey: aesKey256, decryptKey: aesKey256},
		{name: "exact block size", plaintext: bytes.Repeat([]byte("a"), 16), encryptKey: aesKey256, decryptKey: aesKey256},
		{name: "invalid encrypt key", plaintext: []byte("hello"), encryptKey: aesKeyBad, wantEncryptErr: true},
		// --- DecryptCBC error paths ---
		{
			name:       "too short ciphertext",
			directCT:   func() []byte { return []byte("short") },
			decryptKey: aesKey256, wantDecryptErr: true,
		},
		{
			name:       "not a multiple of block size",
			directCT:   func() []byte { return bytes.Repeat([]byte("a"), 17) },
			decryptKey: aesKey256, wantDecryptErr: true,
		},
		{
			name:       "invalid decrypt key",
			directCT:   func() []byte { return bytes.Repeat([]byte("a"), 32) },
			decryptKey: aesKeyBad, wantDecryptErr: true,
		},
		// --- pkcs7Unpad error paths ---
		{
			// Empty body  → pkcs7Unpad receives []byte{} → len(src)==0 branch.
			name:           "pkcs7Unpad: empty body",
			directCT:       func() []byte { return make([]byte, 16) },
			decryptKey:     aesKey256,
			wantDecryptErr: true,
		},
		{
			// Flip IV[15]: last decrypted byte becomes 0x11 (17 > aes.BlockSize)
			// → pkcs7Unpad's "padding > aes.BlockSize" branch.
			name:       "pkcs7Unpad: padding byte > block size",
			plaintext:  []byte("hello"), // 5 bytes → 11 padding bytes of 0x0b
			encryptKey: aesKey256, decryptKey: aesKey256,
			mutate: func(ct []byte) {
				ct[15] ^= 0x0b ^ 0x11 // 0x0b→0x11 (17 > 16)
			},
			wantDecryptErr: true,
		},
		{
			// Flip IV[10]: plaintext[10] (a mid-padding byte) becomes ≠ 0x0b
			// while plaintext[15] stays 0x0b (padding value 11) → inconsistent padding.
			name:       "pkcs7Unpad: inconsistent padding bytes",
			plaintext:  []byte("hello"), // 5 bytes → 11 padding bytes of 0x0b
			encryptKey: aesKey256, decryptKey: aesKey256,
			mutate: func(ct []byte) {
				ct[10] ^= 0xff // corrupt one mid-padding byte
			},
			wantDecryptErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ct []byte

			if tt.directCT != nil {
				ct = tt.directCT()
			} else {
				var err error
				ct, err = lxcrypto.EncryptCBC(tt.plaintext, tt.encryptKey)
				if tt.wantEncryptErr {
					if err == nil {
						t.Error("expected encrypt error, got nil")
					}
					return
				}
				if err != nil {
					t.Fatalf("EncryptCBC: %v", err)
				}
				if tt.mutate != nil {
					tt.mutate(ct)
				}
			}

			got, err := lxcrypto.DecryptCBC(ct, tt.decryptKey)
			if tt.wantDecryptErr {
				if err == nil {
					t.Error("expected decrypt error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("DecryptCBC: %v", err)
			}
			want := tt.plaintext
			if want == nil {
				want = []byte{}
			}
			if !bytes.Equal(got, want) {
				t.Errorf("round-trip mismatch: got %q, want %q", got, want)
			}
		})
	}
}

// TestCBCString covers EncryptCBCString and DecryptCBCString.
func TestCBCString(t *testing.T) {
	tests := []struct {
		name           string
		plaintext      string
		key            []byte
		wantErr        bool
		decryptInput   string
		wantDecryptErr bool
	}{
		{name: "round-trip", plaintext: "hello world", key: aesKey256},
		{name: "empty string", plaintext: "", key: aesKey256},
		{name: "invalid key", plaintext: "hello", key: aesKeyBad, wantErr: true},
		{
			name:           "invalid base64 on decrypt",
			plaintext:      "unused",
			key:            aesKey256,
			decryptInput:   "not!!valid@@base64",
			wantDecryptErr: true,
		},
		{
			name:           "valid base64 but bad CBC data",
			plaintext:      "unused",
			key:            aesKey256,
			decryptInput:   base64.URLEncoding.EncodeToString([]byte("short")),
			wantDecryptErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.decryptInput != "" {
				_, err := lxcrypto.DecryptCBCString(tt.decryptInput, tt.key)
				if !tt.wantDecryptErr {
					if err != nil {
						t.Fatalf("unexpected error: %v", err)
					}
				} else if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}

			ct, err := lxcrypto.EncryptCBCString(tt.plaintext, tt.key)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("EncryptCBCString: %v", err)
			}
			got, err := lxcrypto.DecryptCBCString(ct, tt.key)
			if err != nil {
				t.Fatalf("DecryptCBCString: %v", err)
			}
			if got != tt.plaintext {
				t.Errorf("round-trip mismatch: got %q, want %q", got, tt.plaintext)
			}
		})
	}
}
