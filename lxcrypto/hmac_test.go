package lxcrypto_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxcrypto"
)

// ── HMAC256 ──────────────────────────────────────────────────────────────────

func TestHMAC256_Bytes(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		key      []byte
		expected string
	}{
		{
			name:     "empty data empty key",
			data:     []byte{},
			key:      []byte{},
			expected: "b613679a0814d9ec772f95d778c35fc5ff1697c493715653c6c712144292c5ad",
		},
		{
			name:     "hello / secret",
			data:     []byte("hello"),
			key:      []byte("secret"),
			expected: "88aab3ede8d3adf94d26ab90d3bafd4a2083070c3bcce9c014ee04a443847c0b",
		},
		{
			name:     "message / key",
			data:     []byte("message"),
			key:      []byte("key"),
			expected: "6e9ef29b75fffc5b7abae527d58fdadb2fe42e7219011976917343065f58ed4a",
		},
		{
			name:     "nil data",
			data:     nil,
			key:      []byte("key"),
			expected: "5d5d139563c95b5967b9bd9a8c9b233a9dedb45072794cd232dc1b74832607d0",
		},
		{
			name:     "nil key",
			data:     []byte("data"),
			key:      nil,
			expected: "e528c4d99e6177f5841f712a143b90843299a4aa181a06501422d9ca862bd2a5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxcrypto.HMAC256(tt.data, tt.key)
			if result != tt.expected {
				t.Errorf("HMAC256() = %q; want %q", result, tt.expected)
			}
			if len(result) != 64 {
				t.Errorf("HMAC256() length = %d; want 64", len(result))
			}
			if lxcrypto.HMAC256(tt.data, tt.key) != result {
				t.Error("HMAC256() is not deterministic")
			}
		})
	}
	// different keys must not produce the same tag
	if lxcrypto.HMAC256("data", "key1") == lxcrypto.HMAC256("data", "key2") {
		t.Error("HMAC256() produced same tag for different keys")
	}
}

func TestHMAC256_String(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		key      string
		expected string
	}{
		{
			name:     "empty data empty key",
			data:     "",
			key:      "",
			expected: "b613679a0814d9ec772f95d778c35fc5ff1697c493715653c6c712144292c5ad",
		},
		{
			name:     "hello / secret",
			data:     "hello",
			key:      "secret",
			expected: "88aab3ede8d3adf94d26ab90d3bafd4a2083070c3bcce9c014ee04a443847c0b",
		},
		{
			name:     "message / key",
			data:     "message",
			key:      "key",
			expected: "6e9ef29b75fffc5b7abae527d58fdadb2fe42e7219011976917343065f58ed4a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxcrypto.HMAC256(tt.data, tt.key)
			if result != tt.expected {
				t.Errorf("HMAC256() = %q; want %q", result, tt.expected)
			}
			// string and []byte inputs must produce the same tag
			if result != lxcrypto.HMAC256([]byte(tt.data), []byte(tt.key)) {
				t.Errorf("HMAC256(string) != HMAC256([]byte) for data=%q key=%q", tt.data, tt.key)
			}
		})
	}
}

// ── HMAC512 ──────────────────────────────────────────────────────────────────

func TestHMAC512_Bytes(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		key      []byte
		expected string
	}{
		{
			name:     "empty data empty key",
			data:     []byte{},
			key:      []byte{},
			expected: "b936cee86c9f87aa5d3c6f2e84cb5a4239a5fe50480a6ec66b70ab5b1f4ac6730c6c515421b327ec1d69402e53dfb49ad7381eb067b338fd7b0cb22247225d47",
		},
		{
			name:     "hello / secret",
			data:     []byte("hello"),
			key:      []byte("secret"),
			expected: "db1595ae88a62fd151ec1cba81b98c39df82daae7b4cb9820f446d5bf02f1dcfca6683d88cab3e273f5963ab8ec469a746b5b19086371239f67d1e5f99a79440",
		},
		{
			name:     "message / key",
			data:     []byte("message"),
			key:      []byte("key"),
			expected: "e477384d7ca229dd1426e64b63ebf2d36ebd6d7e669a6735424e72ea6c01d3f8b56eb39c36d8232f5427999b8d1a3f9cd1128fc69f4d75b434216810fa367e98",
		},
		{
			name:     "nil data",
			data:     nil,
			key:      []byte("key"),
			expected: "84fa5aa0279bbc473267d05a53ea03310a987cecc4c1535ff29b6d76b8f1444a728df3aadb89d4a9a6709e1998f373566e8f824a8ca93b1821f0b69bc2a2f65e",
		},
		{
			name:     "nil key",
			data:     []byte("data"),
			key:      nil,
			expected: "768c8b8791fd5a59d1cd1edb860054d746b181926b0551bcff4bd4e135f4bbc89e395f9f250f8b582ebe92d3ff63dd401d3ab2af85790b24ecd92dce7466c16d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxcrypto.HMAC512(tt.data, tt.key)
			if result != tt.expected {
				t.Errorf("HMAC512() = %q; want %q", result, tt.expected)
			}
			if len(result) != 128 {
				t.Errorf("HMAC512() length = %d; want 128", len(result))
			}
			if lxcrypto.HMAC512(tt.data, tt.key) != result {
				t.Error("HMAC512() is not deterministic")
			}
		})
	}
	if lxcrypto.HMAC512("data", "key1") == lxcrypto.HMAC512("data", "key2") {
		t.Error("HMAC512() produced same tag for different keys")
	}
}

func TestHMAC512_String(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		key      string
		expected string
	}{
		{
			name:     "empty data empty key",
			data:     "",
			key:      "",
			expected: "b936cee86c9f87aa5d3c6f2e84cb5a4239a5fe50480a6ec66b70ab5b1f4ac6730c6c515421b327ec1d69402e53dfb49ad7381eb067b338fd7b0cb22247225d47",
		},
		{
			name:     "hello / secret",
			data:     "hello",
			key:      "secret",
			expected: "db1595ae88a62fd151ec1cba81b98c39df82daae7b4cb9820f446d5bf02f1dcfca6683d88cab3e273f5963ab8ec469a746b5b19086371239f67d1e5f99a79440",
		},
		{
			name:     "message / key",
			data:     "message",
			key:      "key",
			expected: "e477384d7ca229dd1426e64b63ebf2d36ebd6d7e669a6735424e72ea6c01d3f8b56eb39c36d8232f5427999b8d1a3f9cd1128fc69f4d75b434216810fa367e98",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxcrypto.HMAC512(tt.data, tt.key)
			if result != tt.expected {
				t.Errorf("HMAC512() = %q; want %q", result, tt.expected)
			}
			if result != lxcrypto.HMAC512([]byte(tt.data), []byte(tt.key)) {
				t.Errorf("HMAC512(string) != HMAC512([]byte) for data=%q key=%q", tt.data, tt.key)
			}
		})
	}
}
