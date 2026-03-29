package lxmaps_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/lxtypes"
	"github.com/nthanhhai2909/lx/maps"
)

func TestFromEntries_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		entries  []lxtypes.Pair[string, int]
		wantNil  bool
		expected map[string]int
	}{
		{
			name:     "nil slice",
			entries:  nil,
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "empty non-nil slice",
			entries:  []lxtypes.Pair[string, int]{},
			wantNil:  false,
			expected: map[string]int{},
		},
		{
			name: "single pair",
			entries: []lxtypes.Pair[string, int]{
				{First: "a", Second: 1},
			},
			wantNil:  false,
			expected: map[string]int{"a": 1},
		},
		{
			name: "multiple pairs",
			entries: []lxtypes.Pair[string, int]{
				{First: "a", Second: 1},
				{First: "b", Second: 2},
				{First: "c", Second: 3},
			},
			wantNil:  false,
			expected: map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name: "duplicate key last wins",
			entries: []lxtypes.Pair[string, int]{
				{First: "a", Second: 1},
				{First: "a", Second: 99},
			},
			wantNil:  false,
			expected: map[string]int{"a": 99},
		},
		{
			name: "zero value",
			entries: []lxtypes.Pair[string, int]{
				{First: "z", Second: 0},
			},
			wantNil:  false,
			expected: map[string]int{"z": 0},
		},
		{
			name: "negative value",
			entries: []lxtypes.Pair[string, int]{
				{First: "n", Second: -1},
			},
			wantNil:  false,
			expected: map[string]int{"n": -1},
		},
		{
			name: "empty string key",
			entries: []lxtypes.Pair[string, int]{
				{First: "", Second: 42},
				{First: "a", Second: 7},
			},
			wantNil:  false,
			expected: map[string]int{"": 42, "a": 7},
		},
		{
			name: "unicode keys",
			entries: []lxtypes.Pair[string, int]{
				{First: "こんにちは", Second: 1},
				{First: "世界", Second: 2},
			},
			wantNil:  false,
			expected: map[string]int{"こんにちは": 1, "世界": 2},
		},
		{
			name: "special character keys",
			entries: []lxtypes.Pair[string, int]{
				{First: "!@#", Second: 10},
				{First: "a:b", Second: 20},
			},
			wantNil:  false,
			expected: map[string]int{"!@#": 10, "a:b": 20},
		},
		{
			name: "emoji keys",
			entries: []lxtypes.Pair[string, int]{
				{First: "🚀", Second: 1},
				{First: "✓", Second: 2},
			},
			wantNil:  false,
			expected: map[string]int{"🚀": 1, "✓": 2},
		},
		{
			name: "large int value",
			entries: []lxtypes.Pair[string, int]{
				{First: "big", Second: 1_000_000},
			},
			wantNil:  false,
			expected: map[string]int{"big": 1_000_000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.FromEntries(tt.entries)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("FromEntries() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("FromEntries() = nil, want non-nil map")
			}
			if !lxmaps.Equal(got, tt.expected) {
				t.Fatalf("FromEntries() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFromEntries_IntString(t *testing.T) {
	tests := []struct {
		name     string
		entries  []lxtypes.Pair[int, string]
		wantNil  bool
		expected map[int]string
	}{
		{
			name:     "nil slice",
			entries:  nil,
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "empty slice",
			entries:  []lxtypes.Pair[int, string]{},
			wantNil:  false,
			expected: map[int]string{},
		},
		{
			name: "single pair",
			entries: []lxtypes.Pair[int, string]{
				{First: 1, Second: "one"},
			},
			wantNil:  false,
			expected: map[int]string{1: "one"},
		},
		{
			name: "multiple pairs",
			entries: []lxtypes.Pair[int, string]{
				{First: 1, Second: "a"},
				{First: 2, Second: "b"},
				{First: 3, Second: "c"},
			},
			wantNil:  false,
			expected: map[int]string{1: "a", 2: "b", 3: "c"},
		},
		{
			name: "duplicate int key last wins",
			entries: []lxtypes.Pair[int, string]{
				{First: 0, Second: "first"},
				{First: 0, Second: "second"},
			},
			wantNil:  false,
			expected: map[int]string{0: "second"},
		},
		{
			name: "zero key",
			entries: []lxtypes.Pair[int, string]{
				{First: 0, Second: "zero"},
				{First: 1, Second: "one"},
			},
			wantNil:  false,
			expected: map[int]string{0: "zero", 1: "one"},
		},
		{
			name: "negative keys",
			entries: []lxtypes.Pair[int, string]{
				{First: -1, Second: "neg"},
				{First: -2, Second: "neg2"},
			},
			wantNil:  false,
			expected: map[int]string{-1: "neg", -2: "neg2"},
		},
		{
			name: "large key",
			entries: []lxtypes.Pair[int, string]{
				{First: 1_000_000, Second: "big"},
			},
			wantNil:  false,
			expected: map[int]string{1_000_000: "big"},
		},
		{
			name: "empty string value",
			entries: []lxtypes.Pair[int, string]{
				{First: 1, Second: ""},
				{First: 2, Second: "x"},
			},
			wantNil:  false,
			expected: map[int]string{1: "", 2: "x"},
		},
		{
			name: "unicode values",
			entries: []lxtypes.Pair[int, string]{
				{First: 1, Second: "こんにちは"},
				{First: 2, Second: "world"},
			},
			wantNil:  false,
			expected: map[int]string{1: "こんにちは", 2: "world"},
		},
		{
			name: "emoji values",
			entries: []lxtypes.Pair[int, string]{
				{First: 1, Second: "😊"},
				{First: 2, Second: "🚀"},
			},
			wantNil:  false,
			expected: map[int]string{1: "😊", 2: "🚀"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.FromEntries(tt.entries)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("FromEntries() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("FromEntries() = nil, want non-nil map")
			}
			if !lxmaps.Equal(got, tt.expected) {
				t.Fatalf("FromEntries() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFromEntries_StringBool(t *testing.T) {
	tests := []struct {
		name     string
		entries  []lxtypes.Pair[string, bool]
		wantNil  bool
		expected map[string]bool
	}{
		{
			name:     "nil slice",
			entries:  nil,
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "empty slice",
			entries:  []lxtypes.Pair[string, bool]{},
			wantNil:  false,
			expected: map[string]bool{},
		},
		{
			name: "true and false via NewPair",
			entries: []lxtypes.Pair[string, bool]{
				lxtypes.NewPair("t", true),
				lxtypes.NewPair("f", false),
			},
			wantNil:  false,
			expected: map[string]bool{"t": true, "f": false},
		},
		{
			name: "single true",
			entries: []lxtypes.Pair[string, bool]{
				{First: "yes", Second: true},
			},
			wantNil:  false,
			expected: map[string]bool{"yes": true},
		},
		{
			name: "single false",
			entries: []lxtypes.Pair[string, bool]{
				{First: "no", Second: false},
			},
			wantNil:  false,
			expected: map[string]bool{"no": false},
		},
		{
			name: "duplicate key last wins",
			entries: []lxtypes.Pair[string, bool]{
				{First: "x", Second: false},
				{First: "x", Second: true},
			},
			wantNil:  false,
			expected: map[string]bool{"x": true},
		},
		{
			name: "empty string key",
			entries: []lxtypes.Pair[string, bool]{
				{First: "", Second: true},
				{First: "a", Second: false},
			},
			wantNil:  false,
			expected: map[string]bool{"": true, "a": false},
		},
		{
			name: "unicode keys",
			entries: []lxtypes.Pair[string, bool]{
				{First: "はい", Second: true},
				{First: "いいえ", Second: false},
			},
			wantNil:  false,
			expected: map[string]bool{"はい": true, "いいえ": false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.FromEntries(tt.entries)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("FromEntries() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("FromEntries() = nil, want non-nil map")
			}
			if !lxmaps.Equal(got, tt.expected) {
				t.Fatalf("FromEntries() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFromEntries_StringStruct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	tests := []struct {
		name     string
		entries  []lxtypes.Pair[string, User]
		wantNil  bool
		expected map[string]User
	}{
		{
			name:     "nil slice",
			entries:  nil,
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "empty slice",
			entries:  []lxtypes.Pair[string, User]{},
			wantNil:  false,
			expected: map[string]User{},
		},
		{
			name: "single user",
			entries: []lxtypes.Pair[string, User]{
				{First: "a", Second: User{Name: "Ann", Age: 1}},
			},
			wantNil:  false,
			expected: map[string]User{"a": {Name: "Ann", Age: 1}},
		},
		{
			name: "multiple users",
			entries: []lxtypes.Pair[string, User]{
				{First: "alice", Second: User{Name: "Alice", Age: 25}},
				{First: "bob", Second: User{Name: "Bob", Age: 30}},
			},
			wantNil: false,
			expected: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
		},
		{
			name: "zero value struct",
			entries: []lxtypes.Pair[string, User]{
				{First: "empty", Second: User{}},
			},
			wantNil:  false,
			expected: map[string]User{"empty": {}},
		},
		{
			name: "duplicate key last wins",
			entries: []lxtypes.Pair[string, User]{
				{First: "k", Second: User{Name: "First", Age: 1}},
				{First: "k", Second: User{Name: "Last", Age: 2}},
			},
			wantNil:  false,
			expected: map[string]User{"k": {Name: "Last", Age: 2}},
		},
		{
			name: "unicode name",
			entries: []lxtypes.Pair[string, User]{
				{First: "jp", Second: User{Name: "太郎", Age: 20}},
			},
			wantNil:  false,
			expected: map[string]User{"jp": {Name: "太郎", Age: 20}},
		},
		{
			name: "negative age",
			entries: []lxtypes.Pair[string, User]{
				{First: "x", Second: User{Name: "X", Age: -1}},
			},
			wantNil:  false,
			expected: map[string]User{"x": {Name: "X", Age: -1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.FromEntries(tt.entries)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("FromEntries() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("FromEntries() = nil, want non-nil map")
			}
			if !lxmaps.Equal(got, tt.expected) {
				t.Fatalf("FromEntries() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFromEntries_RoundTrip(t *testing.T) {
	t.Run("string int typical", func(t *testing.T) {
		orig := map[string]int{"a": 1, "b": 2, "c": 3}
		got := lxmaps.FromEntries(lxmaps.Entries(orig))
		if !lxmaps.Equal(orig, got) {
			t.Fatalf("got %v, want %v", got, orig)
		}
	})

	t.Run("string int empty map", func(t *testing.T) {
		orig := map[string]int{}
		got := lxmaps.FromEntries(lxmaps.Entries(orig))
		if !lxmaps.Equal(orig, got) {
			t.Fatalf("got %v, want %v", got, orig)
		}
	})

	t.Run("string int single entry", func(t *testing.T) {
		orig := map[string]int{"only": 42}
		got := lxmaps.FromEntries(lxmaps.Entries(orig))
		if !lxmaps.Equal(orig, got) {
			t.Fatalf("got %v, want %v", got, orig)
		}
	})

	t.Run("string int zero values", func(t *testing.T) {
		orig := map[string]int{"z": 0, "a": 1}
		got := lxmaps.FromEntries(lxmaps.Entries(orig))
		if !lxmaps.Equal(orig, got) {
			t.Fatalf("got %v, want %v", got, orig)
		}
	})

	t.Run("string int unicode keys", func(t *testing.T) {
		orig := map[string]int{"こんにちは": 1, "世界": 2}
		got := lxmaps.FromEntries(lxmaps.Entries(orig))
		if !lxmaps.Equal(orig, got) {
			t.Fatalf("got %v, want %v", got, orig)
		}
	})

	t.Run("string int many keys", func(t *testing.T) {
		orig := map[string]int{
			"k1": 1, "k2": 2, "k3": 3, "k4": 4, "k5": 5,
			"k6": 6, "k7": 7,
		}
		got := lxmaps.FromEntries(lxmaps.Entries(orig))
		if !lxmaps.Equal(orig, got) {
			t.Fatalf("got %v, want %v", got, orig)
		}
	})

	t.Run("int string", func(t *testing.T) {
		orig := map[int]string{1: "a", 0: "zero", -1: "neg"}
		got := lxmaps.FromEntries(lxmaps.Entries(orig))
		if !lxmaps.Equal(orig, got) {
			t.Fatalf("got %v, want %v", got, orig)
		}
	})

	t.Run("string bool", func(t *testing.T) {
		orig := map[string]bool{"t": true, "f": false}
		got := lxmaps.FromEntries(lxmaps.Entries(orig))
		if !lxmaps.Equal(orig, got) {
			t.Fatalf("got %v, want %v", got, orig)
		}
	})

	t.Run("nil map entries yields empty slice then empty map not nil orig", func(t *testing.T) {
		var m map[string]int
		entries := lxmaps.Entries(m)
		if entries == nil {
			t.Fatal("Entries(nil map) got nil slice, want empty non-nil slice")
		}
		if len(entries) != 0 {
			t.Fatalf("len(entries) = %d, want 0", len(entries))
		}
		got := lxmaps.FromEntries(entries)
		want := map[string]int{}
		if !lxmaps.Equal(got, want) {
			t.Fatalf("FromEntries after Entries(nil) = %v, want empty map %v", got, want)
		}
		if lxmaps.Equal(m, got) {
			t.Fatal("Equal(nil map, empty non-nil map) should be false")
		}
	})
}

func TestFromEntries_RoundTripMultipleMaps(t *testing.T) {
	t.Run("disjoint keys", func(t *testing.T) {
		m1 := map[string]int{"a": 1}
		m2 := map[string]int{"b": 2}
		got := lxmaps.FromEntries(lxmaps.Entries(m1, m2))
		want := map[string]int{"a": 1, "b": 2}
		if !lxmaps.Equal(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("three maps", func(t *testing.T) {
		got := lxmaps.FromEntries(lxmaps.Entries(
			map[string]int{"a": 1},
			map[string]int{"b": 2},
			map[string]int{"c": 3},
		))
		want := map[string]int{"a": 1, "b": 2, "c": 3}
		if !lxmaps.Equal(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("empty map among inputs", func(t *testing.T) {
		got := lxmaps.FromEntries(lxmaps.Entries(
			map[string]int{"a": 1},
			map[string]int{},
			map[string]int{"b": 2},
		))
		want := map[string]int{"a": 1, "b": 2}
		if !lxmaps.Equal(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("overlapping keys later map wins in slice order", func(t *testing.T) {
		// Entries walks each map in argument order; within each map iteration is random.
		// For same key in m1 and m2, both appear in slice; FromEntries last in slice wins.
		m1 := map[string]int{"x": 1}
		m2 := map[string]int{"x": 2}
		got := lxmaps.FromEntries(lxmaps.Entries(m1, m2))
		if len(got) != 1 {
			t.Fatalf("len = %d, want 1", len(got))
		}
		if got["x"] != 2 {
			t.Fatalf("got[x] = %d, want 2 (value from second map must follow first in Entries order)", got["x"])
		}
	})
}

func BenchmarkFromEntries(b *testing.B) {
	entries := []lxtypes.Pair[string, int]{
		{First: "a", Second: 1},
		{First: "b", Second: 2},
		{First: "c", Second: 3},
		{First: "d", Second: 4},
		{First: "e", Second: 5},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lxmaps.FromEntries(entries)
	}
}
