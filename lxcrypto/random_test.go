package lxcrypto_test

import (
	"encoding/base64"
	"strings"
	"testing"
	"unicode"

	"github.com/hgapdvn/lx/lxcrypto"
)

const alphanumericCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func TestRandomBytes(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		wantLen int
		wantErr bool
		check   func(t *testing.T, b []byte)
	}{
		{name: "1 byte", n: 1, wantLen: 1},
		{name: "16 bytes", n: 16, wantLen: 16},
		{name: "64 bytes", n: 64, wantLen: 64},
		{
			name:    "uniqueness",
			n:       32,
			wantLen: 32,
			check: func(t *testing.T, b []byte) {
				b2, err := lxcrypto.RandomBytes(32)
				if err != nil {
					t.Fatalf("second call failed: %v", err)
				}
				if string(b) == string(b2) {
					t.Error("two consecutive calls returned identical values")
				}
			},
		},
		{name: "zero", n: 0, wantErr: true},
		{name: "negative", n: -1, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := lxcrypto.RandomBytes(tt.n)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				if b != nil {
					t.Errorf("expected nil on error, got %v", b)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(b) != tt.wantLen {
				t.Errorf("expected length %d, got %d", tt.wantLen, len(b))
			}
			if tt.check != nil {
				tt.check(t, b)
			}
		})
	}
}

func TestRandomString(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		wantErr bool
		check   func(t *testing.T, s string)
	}{
		{name: "1 char", n: 1},
		{name: "16 chars", n: 16},
		{
			name: "charset only alphanumeric",
			n:    256,
			check: func(t *testing.T, s string) {
				for _, c := range s {
					if !strings.ContainsRune(alphanumericCharset, c) {
						t.Errorf("character %q is not in the allowed charset", c)
					}
				}
			},
		},
		{
			name: "only letters and digits",
			n:    128,
			check: func(t *testing.T, s string) {
				for _, c := range s {
					if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
						t.Errorf("non-alphanumeric character found: %q", c)
					}
				}
			},
		},
		{
			name: "uniqueness",
			n:    32,
			check: func(t *testing.T, s string) {
				s2, err := lxcrypto.RandomString(32)
				if err != nil {
					t.Fatalf("second call failed: %v", err)
				}
				if s == s2 {
					t.Error("two consecutive calls returned identical values")
				}
			},
		},
		{name: "zero", n: 0, wantErr: true},
		{name: "negative", n: -5, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := lxcrypto.RandomString(tt.n)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				if s != "" {
					t.Errorf("expected empty string on error, got %q", s)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(s) != tt.n {
				t.Errorf("expected length %d, got %d", tt.n, len(s))
			}
			if tt.check != nil {
				tt.check(t, s)
			}
		})
	}
}

func TestSecureToken(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		wantErr bool
		check   func(t *testing.T, token string)
	}{
		{name: "1 byte", n: 1},
		{name: "16 bytes", n: 16},
		{name: "32 bytes", n: 32},
		{
			name: "valid url-safe base64",
			n:    64,
			check: func(t *testing.T, token string) {
				decoded, err := base64.URLEncoding.DecodeString(token)
				if err != nil {
					t.Errorf("token is not valid URL-safe base64: %v", err)
				}
				if len(decoded) != 64 {
					t.Errorf("expected decoded length 64, got %d", len(decoded))
				}
				for _, c := range token {
					if !unicode.IsLetter(c) && !unicode.IsDigit(c) && c != '-' && c != '_' && c != '=' {
						t.Errorf("unexpected character in token: %q", c)
					}
				}
			},
		},
		{
			name: "uniqueness",
			n:    32,
			check: func(t *testing.T, token string) {
				token2, err := lxcrypto.SecureToken(32)
				if err != nil {
					t.Fatalf("second call failed: %v", err)
				}
				if token == token2 {
					t.Error("two consecutive calls returned identical values")
				}
			},
		},
		{name: "zero", n: 0, wantErr: true},
		{name: "negative", n: -3, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := lxcrypto.SecureToken(tt.n)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				if token != "" {
					t.Errorf("expected empty token on error, got %q", token)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			decoded, decErr := base64.URLEncoding.DecodeString(token)
			if decErr != nil {
				t.Errorf("token is not valid URL-safe base64: %v", decErr)
			}
			if len(decoded) != tt.n {
				t.Errorf("expected decoded length %d, got %d", tt.n, len(decoded))
			}
			if tt.check != nil {
				tt.check(t, token)
			}
		})
	}
}
