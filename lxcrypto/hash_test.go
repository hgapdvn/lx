package lxcrypto_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxcrypto"
)

func TestMD5(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  string
		outputLen int
	}{
		{
			name:      "nil input",
			input:     nil,
			expected:  "d41d8cd98f00b204e9800998ecf8427e",
			outputLen: 32,
		},
		{
			name:      "empty bytes",
			input:     []byte{},
			expected:  "d41d8cd98f00b204e9800998ecf8427e",
			outputLen: 32,
		},
		{
			name:      "hello",
			input:     []byte("hello"),
			expected:  "5d41402abc4b2a76b9719d911017c592",
			outputLen: 32,
		},
		{
			name:      "hello world",
			input:     []byte("hello world"),
			expected:  "5eb63bbbe01eeed093cb22bb8f5acdc3",
			outputLen: 32,
		},
		{
			name:      "Go",
			input:     []byte("Go"),
			expected:  "5f075ae3e1f9d0382bb8c4632991f96f",
			outputLen: 32,
		},
		{
			name:      "single zero byte",
			input:     []byte{0x00},
			expected:  "93b885adfe0da089cdf634904fd59f71",
			outputLen: 32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxcrypto.MD5(tt.input)
			if result != tt.expected {
				t.Errorf("MD5() = %q; want %q", result, tt.expected)
			}
			if len(result) != tt.outputLen {
				t.Errorf("MD5() length = %d; want %d", len(result), tt.outputLen)
			}
			// determinism: same input must always produce the same digest
			if lxcrypto.MD5(tt.input) != result {
				t.Error("MD5() is not deterministic")
			}
		})
	}

	// different inputs must not collide
	if lxcrypto.MD5([]byte("foo")) == lxcrypto.MD5([]byte("bar")) {
		t.Error("MD5() collision for different inputs")
	}
}

func TestMD5String(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			name:     "hello",
			input:    "hello",
			expected: "5d41402abc4b2a76b9719d911017c592",
		},
		{
			name:     "hello world",
			input:    "hello world",
			expected: "5eb63bbbe01eeed093cb22bb8f5acdc3",
		},
		{
			name:     "Go",
			input:    "Go",
			expected: "5f075ae3e1f9d0382bb8c4632991f96f",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxcrypto.MD5String(tt.input)
			if result != tt.expected {
				t.Errorf("MD5String() = %q; want %q", result, tt.expected)
			}
			// must be consistent with MD5([]byte(s))
			if result != lxcrypto.MD5([]byte(tt.input)) {
				t.Errorf("MD5String() != MD5([]byte(s)) for input %q", tt.input)
			}
		})
	}
}

func TestSHA1(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  string
		outputLen int
	}{
		{
			name:      "nil input",
			input:     nil,
			expected:  "da39a3ee5e6b4b0d3255bfef95601890afd80709",
			outputLen: 40,
		},
		{
			name:      "empty bytes",
			input:     []byte{},
			expected:  "da39a3ee5e6b4b0d3255bfef95601890afd80709",
			outputLen: 40,
		},
		{
			name:      "hello",
			input:     []byte("hello"),
			expected:  "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d",
			outputLen: 40,
		},
		{
			name:      "hello world",
			input:     []byte("hello world"),
			expected:  "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed",
			outputLen: 40,
		},
		{
			name:      "Go",
			input:     []byte("Go"),
			expected:  "2e0b45f2a456e8db55f08d7b65e87593a3e9a140",
			outputLen: 40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxcrypto.SHA1(tt.input)
			if result != tt.expected {
				t.Errorf("SHA1() = %q; want %q", result, tt.expected)
			}
			if len(result) != tt.outputLen {
				t.Errorf("SHA1() length = %d; want %d", len(result), tt.outputLen)
			}
			if lxcrypto.SHA1(tt.input) != result {
				t.Error("SHA1() is not deterministic")
			}
		})
	}

	if lxcrypto.SHA1([]byte("foo")) == lxcrypto.SHA1([]byte("bar")) {
		t.Error("SHA1() collision for different inputs")
	}
}

func TestSHA1String(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "da39a3ee5e6b4b0d3255bfef95601890afd80709",
		},
		{
			name:     "hello",
			input:    "hello",
			expected: "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d",
		},
		{
			name:     "hello world",
			input:    "hello world",
			expected: "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed",
		},
		{
			name:     "Go",
			input:    "Go",
			expected: "2e0b45f2a456e8db55f08d7b65e87593a3e9a140",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxcrypto.SHA1String(tt.input)
			if result != tt.expected {
				t.Errorf("SHA1String() = %q; want %q", result, tt.expected)
			}
			if result != lxcrypto.SHA1([]byte(tt.input)) {
				t.Errorf("SHA1String() != SHA1([]byte(s)) for input %q", tt.input)
			}
		})
	}
}

func TestSHA256(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  string
		outputLen int
	}{
		{
			name:      "nil input",
			input:     nil,
			expected:  "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			outputLen: 64,
		},
		{
			name:      "empty bytes",
			input:     []byte{},
			expected:  "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			outputLen: 64,
		},
		{
			name:      "hello",
			input:     []byte("hello"),
			expected:  "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
			outputLen: 64,
		},
		{
			name:      "hello world",
			input:     []byte("hello world"),
			expected:  "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9",
			outputLen: 64,
		},
		{
			name:      "Go",
			input:     []byte("Go"),
			expected:  "6cc8519b91584e8bd435d63341e0838a99721948718b1c9c1e9c358c64ba992a",
			outputLen: 64,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxcrypto.SHA256(tt.input)
			if result != tt.expected {
				t.Errorf("SHA256() = %q; want %q", result, tt.expected)
			}
			if len(result) != tt.outputLen {
				t.Errorf("SHA256() length = %d; want %d", len(result), tt.outputLen)
			}
			if lxcrypto.SHA256(tt.input) != result {
				t.Error("SHA256() is not deterministic")
			}
		})
	}

	if lxcrypto.SHA256([]byte("foo")) == lxcrypto.SHA256([]byte("bar")) {
		t.Error("SHA256() collision for different inputs")
	}
}

func TestSHA256String(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name:     "hello",
			input:    "hello",
			expected: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
		{
			name:     "hello world",
			input:    "hello world",
			expected: "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9",
		},
		{
			name:     "Go",
			input:    "Go",
			expected: "6cc8519b91584e8bd435d63341e0838a99721948718b1c9c1e9c358c64ba992a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxcrypto.SHA256String(tt.input)
			if result != tt.expected {
				t.Errorf("SHA256String() = %q; want %q", result, tt.expected)
			}
			if result != lxcrypto.SHA256([]byte(tt.input)) {
				t.Errorf("SHA256String() != SHA256([]byte(s)) for input %q", tt.input)
			}
		})
	}
}

func TestSHA512(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  string
		outputLen int
	}{
		{
			name:      "nil input",
			input:     nil,
			expected:  "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e",
			outputLen: 128,
		},
		{
			name:      "empty bytes",
			input:     []byte{},
			expected:  "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e",
			outputLen: 128,
		},
		{
			name:      "hello",
			input:     []byte("hello"),
			expected:  "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043",
			outputLen: 128,
		},
		{
			name:      "hello world",
			input:     []byte("hello world"),
			expected:  "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f",
			outputLen: 128,
		},
		{
			name:      "Go",
			input:     []byte("Go"),
			expected:  "9d2c296205aa517e7fd58412b0b1d8c0e03bdb7904ff820be7b4b08b219af79c1fcc9bfaff686f9fa41b279a98e49e723f4ac3a8f21ab9d8da0077ccb70d7e99",
			outputLen: 128,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxcrypto.SHA512(tt.input)
			if result != tt.expected {
				t.Errorf("SHA512() = %q; want %q", result, tt.expected)
			}
			if len(result) != tt.outputLen {
				t.Errorf("SHA512() length = %d; want %d", len(result), tt.outputLen)
			}
			if lxcrypto.SHA512(tt.input) != result {
				t.Error("SHA512() is not deterministic")
			}
		})
	}

	if lxcrypto.SHA512([]byte("foo")) == lxcrypto.SHA512([]byte("bar")) {
		t.Error("SHA512() collision for different inputs")
	}
}

func TestSHA512String(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e",
		},
		{
			name:     "hello",
			input:    "hello",
			expected: "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043",
		},
		{
			name:     "hello world",
			input:    "hello world",
			expected: "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f",
		},
		{
			name:     "Go",
			input:    "Go",
			expected: "9d2c296205aa517e7fd58412b0b1d8c0e03bdb7904ff820be7b4b08b219af79c1fcc9bfaff686f9fa41b279a98e49e723f4ac3a8f21ab9d8da0077ccb70d7e99",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxcrypto.SHA512String(tt.input)
			if result != tt.expected {
				t.Errorf("SHA512String() = %q; want %q", result, tt.expected)
			}
			if result != lxcrypto.SHA512([]byte(tt.input)) {
				t.Errorf("SHA512String() != SHA512([]byte(s)) for input %q", tt.input)
			}
		})
	}
}
