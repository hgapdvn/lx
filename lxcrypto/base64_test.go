package lxcrypto_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxcrypto"
)

// ── Base64Encode ─────────────────────────────────────────────────────────────

func TestBase64Encode(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{name: "nil input", input: nil, expected: ""},
		{name: "empty", input: []byte{}, expected: ""},
		{name: "hello", input: []byte("hello"), expected: "aGVsbG8="},
		{name: "hello world", input: []byte("hello world"), expected: "aGVsbG8gd29ybGQ="},
		{name: "Go", input: []byte("Go"), expected: "R28="},
		{name: "binary", input: []byte{0x00, 0x01, 0x02}, expected: "AAEC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxcrypto.Base64Encode(tt.input)
			if result != tt.expected {
				t.Errorf("Base64Encode() = %q; want %q", result, tt.expected)
			}
		})
	}
}

// ── Base64EncodeString ───────────────────────────────────────────────────────

func TestBase64EncodeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty string", input: "", expected: ""},
		{name: "hello", input: "hello", expected: "aGVsbG8="},
		{name: "hello world", input: "hello world", expected: "aGVsbG8gd29ybGQ="},
		{name: "Go", input: "Go", expected: "R28="},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxcrypto.Base64EncodeString(tt.input)
			if result != tt.expected {
				t.Errorf("Base64EncodeString() = %q; want %q", result, tt.expected)
			}
			// must match the []byte variant
			if result != lxcrypto.Base64Encode([]byte(tt.input)) {
				t.Errorf("Base64EncodeString() != Base64Encode([]byte) for input %q", tt.input)
			}
		})
	}
}

// ── Base64Decode ─────────────────────────────────────────────────────────────

func TestBase64Decode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []byte
		wantErr  bool
	}{
		{name: "empty string", input: "", expected: []byte{}},
		{name: "hello", input: "aGVsbG8=", expected: []byte("hello")},
		{name: "hello world", input: "aGVsbG8gd29ybGQ=", expected: []byte("hello world")},
		{name: "Go", input: "R28=", expected: []byte("Go")},
		{name: "binary", input: "AAEC", expected: []byte{0x00, 0x01, 0x02}},
		{name: "invalid base64", input: "not!valid@@", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := lxcrypto.Base64Decode(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Error("Base64Decode() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("Base64Decode() unexpected error: %v", err)
			}
			if string(result) != string(tt.expected) {
				t.Errorf("Base64Decode() = %q; want %q", result, tt.expected)
			}
		})
	}
}

func TestBase64DecodeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{name: "empty string", input: "", expected: ""},
		{name: "hello", input: "aGVsbG8=", expected: "hello"},
		{name: "hello world", input: "aGVsbG8gd29ybGQ=", expected: "hello world"},
		{name: "Go", input: "R28=", expected: "Go"},
		{name: "invalid base64", input: "not!valid@@", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := lxcrypto.Base64DecodeString(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Error("Base64DecodeString() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("Base64DecodeString() unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Base64DecodeString() = %q; want %q", result, tt.expected)
			}
		})
	}
}

// ── Base64URLEncode ───────────────────────────────────────────────────────────

func TestBase64URLEncode(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{name: "nil input", input: nil, expected: ""},
		{name: "empty", input: []byte{}, expected: ""},
		{name: "hello", input: []byte("hello"), expected: "aGVsbG8="},
		{name: "hello world", input: []byte("hello world"), expected: "aGVsbG8gd29ybGQ="},
		{name: "Go", input: []byte("Go"), expected: "R28="},
		// 0xfb 0xff encodes to +/8= in standard but -_8= in URL-safe
		{name: "url-safe chars", input: []byte{0xfb, 0xff}, expected: "-_8="},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxcrypto.Base64URLEncode(tt.input)
			if result != tt.expected {
				t.Errorf("Base64URLEncode() = %q; want %q", result, tt.expected)
			}
		})
	}
}

// ── Base64URLEncodeString ────────────────────────────────────────────────────

func TestBase64URLEncodeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty string", input: "", expected: ""},
		{name: "hello", input: "hello", expected: "aGVsbG8="},
		{name: "hello world", input: "hello world", expected: "aGVsbG8gd29ybGQ="},
		{name: "Go", input: "Go", expected: "R28="},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxcrypto.Base64URLEncodeString(tt.input)
			if result != tt.expected {
				t.Errorf("Base64URLEncodeString() = %q; want %q", result, tt.expected)
			}
			if result != lxcrypto.Base64URLEncode([]byte(tt.input)) {
				t.Errorf("Base64URLEncodeString() != Base64URLEncode([]byte) for input %q", tt.input)
			}
		})
	}
}

// ── Base64URLDecode ───────────────────────────────────────────────────────────

func TestBase64URLDecode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []byte
		wantErr  bool
	}{
		{name: "empty string", input: "", expected: []byte{}},
		{name: "hello", input: "aGVsbG8=", expected: []byte("hello")},
		{name: "hello world", input: "aGVsbG8gd29ybGQ=", expected: []byte("hello world")},
		{name: "Go", input: "R28=", expected: []byte("Go")},
		{name: "url-safe chars", input: "-_8=", expected: []byte{0xfb, 0xff}},
		{name: "invalid base64", input: "not!valid@@", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := lxcrypto.Base64URLDecode(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Error("Base64URLDecode() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("Base64URLDecode() unexpected error: %v", err)
			}
			if string(result) != string(tt.expected) {
				t.Errorf("Base64URLDecode() = %q; want %q", result, tt.expected)
			}
		})
	}
}

func TestBase64URLDecodeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{name: "empty string", input: "", expected: ""},
		{name: "hello", input: "aGVsbG8=", expected: "hello"},
		{name: "hello world", input: "aGVsbG8gd29ybGQ=", expected: "hello world"},
		{name: "Go", input: "R28=", expected: "Go"},
		{name: "invalid base64", input: "not!valid@@", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := lxcrypto.Base64URLDecodeString(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Error("Base64URLDecodeString() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("Base64URLDecodeString() unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Base64URLDecodeString() = %q; want %q", result, tt.expected)
			}
		})
	}
}
