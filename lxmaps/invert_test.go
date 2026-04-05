package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxmaps"
)

func TestInvert_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		wantNil  bool
		expected map[int]string
	}{
		{
			name:     "nil map",
			input:    nil,
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			wantNil:  false,
			expected: map[int]string{},
		},
		{
			name:     "single entry",
			input:    map[string]int{"a": 1},
			wantNil:  false,
			expected: map[int]string{1: "a"},
		},
		{
			name:     "two entries",
			input:    map[string]int{"a": 1, "b": 2},
			wantNil:  false,
			expected: map[int]string{1: "a", 2: "b"},
		},
		{
			name: "many entries",
			input: map[string]int{
				"k1": 1, "k2": 2, "k3": 3, "k4": 4, "k5": 5,
			},
			wantNil: false,
			expected: map[int]string{
				1: "k1", 2: "k2", 3: "k3", 4: "k4", 5: "k5",
			},
		},
		{
			name:     "zero value",
			input:    map[string]int{"z": 0},
			wantNil:  false,
			expected: map[int]string{0: "z"},
		},
		{
			name:     "negative value",
			input:    map[string]int{"neg": -1},
			wantNil:  false,
			expected: map[int]string{-1: "neg"},
		},
		{
			name:     "large int value",
			input:    map[string]int{"big": 1_000_000},
			wantNil:  false,
			expected: map[int]string{1_000_000: "big"},
		},
		{
			name: "max int value",
			input: map[string]int{
				"x": int(^uint(0) >> 1),
			},
			wantNil: false,
			expected: map[int]string{
				int(^uint(0) >> 1): "x",
			},
		},
		{
			name:     "empty string key",
			input:    map[string]int{"": 42, "a": 7},
			wantNil:  false,
			expected: map[int]string{42: "", 7: "a"},
		},
		{
			name:     "unicode keys",
			input:    map[string]int{"こんにちは": 1, "世界": 2},
			wantNil:  false,
			expected: map[int]string{1: "こんにちは", 2: "世界"},
		},
		{
			name:     "special character keys",
			input:    map[string]int{"!@#": 10, "a:b": 20},
			wantNil:  false,
			expected: map[int]string{10: "!@#", 20: "a:b"},
		},
		{
			name:     "emoji keys",
			input:    map[string]int{"🚀": 1, "✓": 2},
			wantNil:  false,
			expected: map[int]string{1: "🚀", 2: "✓"},
		},
		{
			name:     "case sensitive keys",
			input:    map[string]int{"A": 1, "a": 2},
			wantNil:  false,
			expected: map[int]string{1: "A", 2: "a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Invert(tt.input)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("Invert() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("Invert() = nil, want non-nil map")
			}
			if len(got) != len(tt.expected) {
				t.Fatalf("len(Invert()) = %d, want %d", len(got), len(tt.expected))
			}
			for k, want := range tt.expected {
				if g, ok := got[k]; !ok || g != want {
					t.Fatalf("Invert()[%d] = %q, ok=%v; want %q", k, g, ok, want)
				}
			}
			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("Invert() unexpected key %d", k)
				}
			}
		})
	}
}

func TestInvert_IntString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		wantNil  bool
		expected map[string]int
	}{
		{
			name:     "nil",
			input:    nil,
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "empty",
			input:    map[int]string{},
			wantNil:  false,
			expected: map[string]int{},
		},
		{
			name:     "swap int key string value",
			input:    map[int]string{1: "one", 2: "two"},
			wantNil:  false,
			expected: map[string]int{"one": 1, "two": 2},
		},
		{
			name: "many entries",
			input: map[int]string{
				10: "a", 20: "b", 30: "c", 40: "d",
			},
			wantNil: false,
			expected: map[string]int{
				"a": 10, "b": 20, "c": 30, "d": 40,
			},
		},
		{
			name:     "zero key",
			input:    map[int]string{0: "zero"},
			wantNil:  false,
			expected: map[string]int{"zero": 0},
		},
		{
			name:     "negative key",
			input:    map[int]string{-1: "neg", -2: "neg2"},
			wantNil:  false,
			expected: map[string]int{"neg": -1, "neg2": -2},
		},
		{
			name:     "large key",
			input:    map[int]string{1_000_000: "big"},
			wantNil:  false,
			expected: map[string]int{"big": 1_000_000},
		},
		{
			name:     "empty string value",
			input:    map[int]string{1: "", 2: "x"},
			wantNil:  false,
			expected: map[string]int{"": 1, "x": 2},
		},
		{
			name:     "unicode values",
			input:    map[int]string{1: "こんにちは", 2: "world"},
			wantNil:  false,
			expected: map[string]int{"こんにちは": 1, "world": 2},
		},
		{
			name:     "emoji values",
			input:    map[int]string{1: "😊", 2: "🚀"},
			wantNil:  false,
			expected: map[string]int{"😊": 1, "🚀": 2},
		},
		{
			name:     "special character values",
			input:    map[int]string{1: "a:b", 2: "x|y"},
			wantNil:  false,
			expected: map[string]int{"a:b": 1, "x|y": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Invert(tt.input)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("Invert() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("Invert() = nil, want non-nil map")
			}
			if len(got) != len(tt.expected) {
				t.Fatalf("len(Invert()) = %d, want %d", len(got), len(tt.expected))
			}
			for k, want := range tt.expected {
				if g, ok := got[k]; !ok || g != want {
					t.Fatalf("Invert()[%q] = %d, ok=%v; want %d", k, g, ok, want)
				}
			}
			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("Invert() unexpected key %q", k)
				}
			}
		})
	}
}

func TestInvert_StringBool(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]bool
		wantNil  bool
		expected map[bool]string
	}{
		{
			name:     "nil",
			input:    nil,
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "empty",
			input:    map[string]bool{},
			wantNil:  false,
			expected: map[bool]string{},
		},
		{
			name:     "true and false",
			input:    map[string]bool{"t": true, "f": false},
			wantNil:  false,
			expected: map[bool]string{true: "t", false: "f"},
		},
		{
			name:     "single true",
			input:    map[string]bool{"yes": true},
			wantNil:  false,
			expected: map[bool]string{true: "yes"},
		},
		{
			name:     "single false",
			input:    map[string]bool{"no": false},
			wantNil:  false,
			expected: map[bool]string{false: "no"},
		},
		{
			name:     "empty string key",
			input:    map[string]bool{"": true, "a": false},
			wantNil:  false,
			expected: map[bool]string{true: "", false: "a"},
		},
		{
			name:     "unicode keys",
			input:    map[string]bool{"はい": true, "いいえ": false},
			wantNil:  false,
			expected: map[bool]string{true: "はい", false: "いいえ"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Invert(tt.input)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("Invert() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("Invert() = nil, want non-nil map")
			}
			if len(got) != len(tt.expected) {
				t.Fatalf("len(Invert()) = %d, want %d", len(got), len(tt.expected))
			}
			for k, want := range tt.expected {
				if g, ok := got[k]; !ok || g != want {
					t.Fatalf("Invert()[%v] = %q, ok=%v; want %q", k, g, ok, want)
				}
			}
			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("Invert() unexpected key %v", k)
				}
			}
		})
	}
}

// TestInvert_StringBool_duplicateTrueKeys: two keys map to true -> one bool key in result.
func TestInvert_StringBool_duplicateTrueKeys(t *testing.T) {
	got := lxmaps.Invert(map[string]bool{"a": true, "b": true})
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
	k := got[true]
	if k != "a" && k != "b" {
		t.Fatalf("Invert()[true] = %q, want %q or %q", k, "a", "b")
	}
}

func TestInvert_StringStruct(t *testing.T) {
	type ID string

	tests := []struct {
		name     string
		input    map[ID]int
		wantNil  bool
		expected map[int]ID
	}{
		{
			name:     "nil",
			input:    nil,
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "empty",
			input:    map[ID]int{},
			wantNil:  false,
			expected: map[int]ID{},
		},
		{
			name:     "two ids",
			input:    map[ID]int{"a": 1, "b": 2},
			wantNil:  false,
			expected: map[int]ID{1: "a", 2: "b"},
		},
		{
			name: "many ids",
			input: map[ID]int{
				"id1": 10, "id2": 20, "id3": 30,
			},
			wantNil: false,
			expected: map[int]ID{
				10: "id1", 20: "id2", 30: "id3",
			},
		},
		{
			name:     "empty string id",
			input:    map[ID]int{"": 0, "x": 1},
			wantNil:  false,
			expected: map[int]ID{0: "", 1: "x"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Invert(tt.input)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("Invert() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("Invert() = nil, want non-nil map")
			}
			if len(got) != len(tt.expected) {
				t.Fatalf("len(Invert()) = %d, want %d", len(got), len(tt.expected))
			}
			for k, want := range tt.expected {
				if g, ok := got[k]; !ok || g != want {
					t.Fatalf("Invert()[%d] = %q, ok=%v; want %q", k, g, ok, want)
				}
			}
			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("Invert() unexpected key %d", k)
				}
			}
		})
	}
}

func TestInvert_ComparableStructKey(t *testing.T) {
	type Key struct {
		X int
		Y int
	}
	input := map[Key]string{
		{1, 2}: "a",
		{3, 4}: "b",
	}
	got := lxmaps.Invert(input)
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	if got["a"] != (Key{1, 2}) || got["b"] != (Key{3, 4}) {
		t.Fatalf("got = %v", got)
	}
}

// TestInvert_DuplicateValues verifies collision behavior when several keys share a value.
func TestInvert_DuplicateValues(t *testing.T) {
	t.Run("string int two keys same value", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 1, "c": 2}
		got := lxmaps.Invert(m)
		if len(got) != 2 {
			t.Fatalf("len = %d, want 2", len(got))
		}
		if got[2] != "c" {
			t.Fatalf("Invert()[2] = %q, want %q", got[2], "c")
		}
		k := got[1]
		if k != "a" && k != "b" {
			t.Fatalf("Invert()[1] = %q, want %q or %q", k, "a", "b")
		}
	})

	t.Run("all keys same int value", func(t *testing.T) {
		got := lxmaps.Invert(map[string]int{"a": 5, "b": 5, "c": 5})
		if len(got) != 1 {
			t.Fatalf("len = %d, want 1", len(got))
		}
		k := got[5]
		if k != "a" && k != "b" && k != "c" {
			t.Fatalf("Invert()[5] = %q, want one of a,b,c", k)
		}
	})

	t.Run("string string duplicate values", func(t *testing.T) {
		got := lxmaps.Invert(map[string]string{
			"k1": "same",
			"k2": "same",
			"k3": "other",
		})
		if len(got) != 2 {
			t.Fatalf("len = %d, want 2", len(got))
		}
		if got["other"] != "k3" {
			t.Fatalf(`Invert()["other"] = %q, want %q`, got["other"], "k3")
		}
		v := got["same"]
		if v != "k1" && v != "k2" {
			t.Fatalf(`Invert()["same"] = %q, want %q or %q`, v, "k1", "k2")
		}
	})

	t.Run("int keys duplicate string values", func(t *testing.T) {
		got := lxmaps.Invert(map[int]string{
			1: "dup",
			2: "dup",
			3: "x",
		})
		if len(got) != 2 {
			t.Fatalf("len = %d, want 2", len(got))
		}
		if got["x"] != 3 {
			t.Fatalf(`Invert()["x"] = %d, want 3`, got["x"])
		}
		n := got["dup"]
		if n != 1 && n != 2 {
			t.Fatalf(`Invert()["dup"] = %d, want 1 or 2`, n)
		}
	})

	t.Run("rune to int duplicate rune values", func(t *testing.T) {
		got := lxmaps.Invert(map[rune]int{
			'a': 1,
			'b': 1,
			'c': 2,
		})
		if len(got) != 2 {
			t.Fatalf("len = %d, want 2", len(got))
		}
		if got[2] != 'c' {
			t.Fatalf("Invert()[2] = %q, want %q", got[2], 'c')
		}
		r := got[1]
		if r != 'a' && r != 'b' {
			t.Fatalf("Invert()[1] = %q, want %q or %q", r, 'a', 'b')
		}
	})
}

func BenchmarkInvert(b *testing.B) {
	m := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lxmaps.Invert(m)
	}
}
