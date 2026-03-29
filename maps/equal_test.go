package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/maps"
)

func TestEqual_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		a        map[string]int
		b        map[string]int
		expected bool
	}{
		{
			name:     "both nil",
			a:        nil,
			b:        nil,
			expected: true,
		},
		{
			name:     "nil vs empty non-nil",
			a:        nil,
			b:        map[string]int{},
			expected: false,
		},
		{
			name:     "empty non-nil vs nil",
			a:        map[string]int{},
			b:        nil,
			expected: false,
		},
		{
			name:     "both empty",
			a:        map[string]int{},
			b:        map[string]int{},
			expected: true,
		},
		{
			name:     "same single entry",
			a:        map[string]int{"a": 1},
			b:        map[string]int{"a": 1},
			expected: true,
		},
		{
			name:     "same entries different map identity",
			a:        map[string]int{"a": 1, "b": 2},
			b:        map[string]int{"b": 2, "a": 1},
			expected: true,
		},
		{
			name:     "different value",
			a:        map[string]int{"a": 1},
			b:        map[string]int{"a": 2},
			expected: false,
		},
		{
			name:     "missing key in b",
			a:        map[string]int{"a": 1, "b": 2},
			b:        map[string]int{"a": 1},
			expected: false,
		},
		{
			name:     "extra key in b",
			a:        map[string]int{"a": 1},
			b:        map[string]int{"a": 1, "b": 2},
			expected: false,
		},
		{
			name:     "zero value equal",
			a:        map[string]int{"z": 0},
			b:        map[string]int{"z": 0},
			expected: true,
		},
		{
			name:     "zero vs missing not equal same len trick",
			a:        map[string]int{"a": 0},
			b:        map[string]int{"b": 0},
			expected: false,
		},
		{
			name:     "negative values",
			a:        map[string]int{"x": -1, "y": -2},
			b:        map[string]int{"y": -2, "x": -1},
			expected: true,
		},
		{
			name:     "empty string key",
			a:        map[string]int{"": 42, "a": 1},
			b:        map[string]int{"a": 1, "": 42},
			expected: true,
		},
		{
			name:     "unicode keys",
			a:        map[string]int{"こんにちは": 1, "世界": 2},
			b:        map[string]int{"世界": 2, "こんにちは": 1},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lxmaps.Equal(tt.a, tt.b); got != tt.expected {
				t.Fatalf("Equal() = %v, want %v", got, tt.expected)
			}
			if got := lxmaps.Equal(tt.b, tt.a); got != tt.expected {
				t.Fatalf("Equal(symmetric) = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestEqual_IntString(t *testing.T) {
	tests := []struct {
		name     string
		a        map[int]string
		b        map[int]string
		expected bool
	}{
		{
			name:     "both nil",
			a:        nil,
			b:        nil,
			expected: true,
		},
		{
			name:     "nil vs empty",
			a:        nil,
			b:        map[int]string{},
			expected: false,
		},
		{
			name:     "both empty",
			a:        map[int]string{},
			b:        map[int]string{},
			expected: true,
		},
		{
			name:     "same",
			a:        map[int]string{1: "a", 2: "b"},
			b:        map[int]string{2: "b", 1: "a"},
			expected: true,
		},
		{
			name:     "different string",
			a:        map[int]string{1: "a"},
			b:        map[int]string{1: "b"},
			expected: false,
		},
		{
			name:     "zero key",
			a:        map[int]string{0: "zero"},
			b:        map[int]string{0: "zero"},
			expected: true,
		},
		{
			name:     "negative key",
			a:        map[int]string{-1: "neg"},
			b:        map[int]string{-1: "neg"},
			expected: true,
		},
		{
			name:     "empty string value",
			a:        map[int]string{1: ""},
			b:        map[int]string{1: ""},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lxmaps.Equal(tt.a, tt.b); got != tt.expected {
				t.Fatalf("Equal() = %v, want %v", got, tt.expected)
			}
			if got := lxmaps.Equal(tt.b, tt.a); got != tt.expected {
				t.Fatalf("Equal(symmetric) = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestEqual_StringBool(t *testing.T) {
	tests := []struct {
		name     string
		a        map[string]bool
		b        map[string]bool
		expected bool
	}{
		{
			name:     "both nil",
			a:        nil,
			b:        nil,
			expected: true,
		},
		{
			name:     "nil vs empty non-nil",
			a:        nil,
			b:        map[string]bool{},
			expected: false,
		},
		{
			name:     "both empty",
			a:        map[string]bool{},
			b:        map[string]bool{},
			expected: true,
		},
		{
			name:     "true vs true",
			a:        map[string]bool{"x": true},
			b:        map[string]bool{"x": true},
			expected: true,
		},
		{
			name:     "false vs false",
			a:        map[string]bool{"x": false},
			b:        map[string]bool{"x": false},
			expected: true,
		},
		{
			name:     "false vs true",
			a:        map[string]bool{"x": false},
			b:        map[string]bool{"x": true},
			expected: false,
		},
		{
			name:     "true vs false",
			a:        map[string]bool{"x": true},
			b:        map[string]bool{"x": false},
			expected: false,
		},
		{
			name: "multiple keys same order independent",
			a: map[string]bool{
				"a": true, "b": false, "c": true,
			},
			b: map[string]bool{
				"c": true, "a": true, "b": false,
			},
			expected: true,
		},
		{
			name: "one key differs",
			a: map[string]bool{
				"a": true, "b": false,
			},
			b: map[string]bool{
				"a": true, "b": true,
			},
			expected: false,
		},
		{
			name: "missing key in b",
			a: map[string]bool{
				"a": true, "b": false,
			},
			b: map[string]bool{
				"a": true,
			},
			expected: false,
		},
		{
			name: "extra key in b",
			a: map[string]bool{
				"a": true,
			},
			b: map[string]bool{
				"a": true, "b": false,
			},
			expected: false,
		},
		{
			name:     "empty string key same",
			a:        map[string]bool{"": true, "a": false},
			b:        map[string]bool{"a": false, "": true},
			expected: true,
		},
		{
			name:     "empty string key different value",
			a:        map[string]bool{"": true},
			b:        map[string]bool{"": false},
			expected: false,
		},
		{
			name:     "unicode keys same",
			a:        map[string]bool{"はい": true, "いいえ": false},
			b:        map[string]bool{"いいえ": false, "はい": true},
			expected: true,
		},
		{
			name:     "case sensitive keys different entries",
			a:        map[string]bool{"T": true, "t": false},
			b:        map[string]bool{"T": true, "t": false},
			expected: true,
		},
		{
			name:     "case sensitive mismatch",
			a:        map[string]bool{"T": true},
			b:        map[string]bool{"t": true},
			expected: false,
		},
		{
			name:     "special character keys",
			a:        map[string]bool{"!@#": true, "$%": false},
			b:        map[string]bool{"$%": false, "!@#": true},
			expected: true,
		},
		{
			name:     "all true",
			a:        map[string]bool{"x": true, "y": true},
			b:        map[string]bool{"y": true, "x": true},
			expected: true,
		},
		{
			name:     "all false",
			a:        map[string]bool{"x": false, "y": false},
			b:        map[string]bool{"y": false, "x": false},
			expected: true,
		},
		{
			name:     "single entry true",
			a:        map[string]bool{"only": true},
			b:        map[string]bool{"only": true},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lxmaps.Equal(tt.a, tt.b); got != tt.expected {
				t.Fatalf("Equal() = %v, want %v", got, tt.expected)
			}
			if got := lxmaps.Equal(tt.b, tt.a); got != tt.expected {
				t.Fatalf("Equal(symmetric) = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestEqual_StringStruct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	tests := []struct {
		name     string
		a        map[string]User
		b        map[string]User
		expected bool
	}{
		{
			name:     "both nil",
			a:        nil,
			b:        nil,
			expected: true,
		},
		{
			name:     "nil vs empty",
			a:        nil,
			b:        map[string]User{},
			expected: false,
		},
		{
			name:     "both empty",
			a:        map[string]User{},
			b:        map[string]User{},
			expected: true,
		},
		{
			name: "same structs",
			a: map[string]User{
				"a": {Name: "Ann", Age: 1},
			},
			b: map[string]User{
				"a": {Name: "Ann", Age: 1},
			},
			expected: true,
		},
		{
			name: "different age",
			a: map[string]User{
				"a": {Name: "Ann", Age: 1},
			},
			b: map[string]User{
				"a": {Name: "Ann", Age: 2},
			},
			expected: false,
		},
		{
			name: "different name same age",
			a: map[string]User{
				"a": {Name: "Ann", Age: 1},
			},
			b: map[string]User{
				"a": {Name: "Bob", Age: 1},
			},
			expected: false,
		},
		{
			name: "zero value struct",
			a: map[string]User{
				"k": {},
			},
			b: map[string]User{
				"k": {},
			},
			expected: true,
		},
		{
			name: "zero value vs populated",
			a: map[string]User{
				"k": {},
			},
			b: map[string]User{
				"k": {Name: "X", Age: 0},
			},
			expected: false,
		},
		{
			name: "multiple users order independent",
			a: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
			b: map[string]User{
				"bob":   {Name: "Bob", Age: 30},
				"alice": {Name: "Alice", Age: 25},
			},
			expected: true,
		},
		{
			name: "missing key in b",
			a: map[string]User{
				"a": {Name: "A", Age: 1},
				"b": {Name: "B", Age: 2},
			},
			b: map[string]User{
				"a": {Name: "A", Age: 1},
			},
			expected: false,
		},
		{
			name: "extra key in b",
			a: map[string]User{
				"a": {Name: "A", Age: 1},
			},
			b: map[string]User{
				"a": {Name: "A", Age: 1},
				"b": {Name: "B", Age: 2},
			},
			expected: false,
		},
		{
			name: "unicode name same",
			a: map[string]User{
				"jp": {Name: "太郎", Age: 20},
			},
			b: map[string]User{
				"jp": {Name: "太郎", Age: 20},
			},
			expected: true,
		},
		{
			name: "unicode name different",
			a: map[string]User{
				"jp": {Name: "太郎", Age: 20},
			},
			b: map[string]User{
				"jp": {Name: "花子", Age: 20},
			},
			expected: false,
		},
		{
			name: "unicode key",
			a: map[string]User{
				"ユーザー":  {Name: "U", Age: 1},
				"admin": {Name: "A", Age: 2},
			},
			b: map[string]User{
				"admin": {Name: "A", Age: 2},
				"ユーザー":  {Name: "U", Age: 1},
			},
			expected: true,
		},
		{
			name: "empty string key",
			a: map[string]User{
				"":  {Name: "Blank", Age: 0},
				"x": {Name: "X", Age: 1},
			},
			b: map[string]User{
				"x": {Name: "X", Age: 1},
				"":  {Name: "Blank", Age: 0},
			},
			expected: true,
		},
		{
			name: "negative age same",
			a: map[string]User{
				"a": {Name: "A", Age: -1},
			},
			b: map[string]User{
				"a": {Name: "A", Age: -1},
			},
			expected: true,
		},
		{
			name: "negative age mismatch",
			a: map[string]User{
				"a": {Name: "A", Age: -1},
			},
			b: map[string]User{
				"a": {Name: "A", Age: -2},
			},
			expected: false,
		},
		{
			name: "case sensitive keys same content",
			a: map[string]User{
				"User": {Name: "Upper", Age: 1},
				"user": {Name: "Lower", Age: 2},
			},
			b: map[string]User{
				"user": {Name: "Lower", Age: 2},
				"User": {Name: "Upper", Age: 1},
			},
			expected: true,
		},
		{
			name: "special character keys",
			a: map[string]User{
				"id:1": {Name: "One", Age: 1},
				"id:2": {Name: "Two", Age: 2},
			},
			b: map[string]User{
				"id:2": {Name: "Two", Age: 2},
				"id:1": {Name: "One", Age: 1},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lxmaps.Equal(tt.a, tt.b); got != tt.expected {
				t.Fatalf("Equal() = %v, want %v", got, tt.expected)
			}
			if got := lxmaps.Equal(tt.b, tt.a); got != tt.expected {
				t.Fatalf("Equal(symmetric) = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkEqual(b *testing.B) {
	m1 := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
	}
	m2 := map[string]int{
		"e": 5, "d": 4, "c": 3, "b": 2, "a": 1,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lxmaps.Equal(m1, m2)
	}
}
