package lxmaps_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/lxmaps"
)

func TestClone_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		wantNil  bool
		wantLen  int
		wantVals map[string]int
	}{
		{
			name:     "nil map",
			input:    nil,
			wantNil:  true,
			wantLen:  0,
			wantVals: nil,
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			wantNil:  false,
			wantLen:  0,
			wantVals: map[string]int{},
		},
		{
			name:     "single entry",
			input:    map[string]int{"a": 1},
			wantNil:  false,
			wantLen:  1,
			wantVals: map[string]int{"a": 1},
		},
		{
			name:     "multiple entries",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			wantNil:  false,
			wantLen:  3,
			wantVals: map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name:     "zero value",
			input:    map[string]int{"z": 0},
			wantNil:  false,
			wantLen:  1,
			wantVals: map[string]int{"z": 0},
		},
		{
			name:     "unicode keys",
			input:    map[string]int{"こんにちは": 1, "世界": 2},
			wantNil:  false,
			wantLen:  2,
			wantVals: map[string]int{"こんにちは": 1, "世界": 2},
		},
		{
			name:     "empty string key",
			input:    map[string]int{"": 42, "a": 1},
			wantNil:  false,
			wantLen:  2,
			wantVals: map[string]int{"": 42, "a": 1},
		},
		{
			name:     "special character keys",
			input:    map[string]int{"!@#": 1, "$%": 2},
			wantNil:  false,
			wantLen:  2,
			wantVals: map[string]int{"!@#": 1, "$%": 2},
		},
		{
			name:     "negative values",
			input:    map[string]int{"a": -1, "b": -2},
			wantNil:  false,
			wantLen:  2,
			wantVals: map[string]int{"a": -1, "b": -2},
		},
		{
			name:     "large int value",
			input:    map[string]int{"big": 1_000_000},
			wantNil:  false,
			wantLen:  1,
			wantVals: map[string]int{"big": 1_000_000},
		},
		{
			name: "max int value",
			input: map[string]int{
				"x": int(^uint(0) >> 1),
			},
			wantNil: false,
			wantLen: 1,
			wantVals: map[string]int{
				"x": int(^uint(0) >> 1),
			},
		},
		{
			name:     "emoji keys",
			input:    map[string]int{"😊": 1, "🚀": 2},
			wantNil:  false,
			wantLen:  2,
			wantVals: map[string]int{"😊": 1, "🚀": 2},
		},
		{
			name: "many keys",
			input: map[string]int{
				"k1": 1, "k2": 2, "k3": 3, "k4": 4, "k5": 5,
				"k6": 6, "k7": 7,
			},
			wantNil: false,
			wantLen: 7,
			wantVals: map[string]int{
				"k1": 1, "k2": 2, "k3": 3, "k4": 4, "k5": 5,
				"k6": 6, "k7": 7,
			},
		},
		{
			name:     "case sensitive keys",
			input:    map[string]int{"A": 1, "a": 2},
			wantNil:  false,
			wantLen:  2,
			wantVals: map[string]int{"A": 1, "a": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Clone(tt.input)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("Clone() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("Clone() = nil, want non-nil map")
			}
			if len(got) != tt.wantLen {
				t.Fatalf("len(Clone()) = %d, want %d", len(got), tt.wantLen)
			}
			for k, want := range tt.wantVals {
				if g, ok := got[k]; !ok || g != want {
					t.Fatalf("Clone()[%q] = %v, ok=%v; want %v", k, g, ok, want)
				}
			}
			for k := range got {
				if _, ok := tt.wantVals[k]; !ok {
					t.Fatalf("Clone() unexpected key %q", k)
				}
			}
		})
	}
}

func TestClone_IntString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		wantNil  bool
		wantVals map[int]string
	}{
		{
			name:     "nil",
			input:    nil,
			wantNil:  true,
			wantVals: nil,
		},
		{
			name:     "empty",
			input:    map[int]string{},
			wantNil:  false,
			wantVals: map[int]string{},
		},
		{
			name:     "negative and zero keys",
			input:    map[int]string{-1: "neg", 0: "zero", 1: "one"},
			wantNil:  false,
			wantVals: map[int]string{-1: "neg", 0: "zero", 1: "one"},
		},
		{
			name:     "single entry",
			input:    map[int]string{42: "answer"},
			wantNil:  false,
			wantVals: map[int]string{42: "answer"},
		},
		{
			name:     "large key",
			input:    map[int]string{1_000_000: "big", 1: "small"},
			wantNil:  false,
			wantVals: map[int]string{1_000_000: "big", 1: "small"},
		},
		{
			name:     "empty string value",
			input:    map[int]string{1: "", 2: "b"},
			wantNil:  false,
			wantVals: map[int]string{1: "", 2: "b"},
		},
		{
			name:     "unicode values",
			input:    map[int]string{1: "hello", 2: "こんにちは"},
			wantNil:  false,
			wantVals: map[int]string{1: "hello", 2: "こんにちは"},
		},
		{
			name:     "emoji values",
			input:    map[int]string{0: "😊", 1: "text"},
			wantNil:  false,
			wantVals: map[int]string{0: "😊", 1: "text"},
		},
		{
			name: "many entries",
			input: map[int]string{
				1: "a", 2: "b", 3: "c", 4: "d", 5: "e",
			},
			wantNil: false,
			wantVals: map[int]string{
				1: "a", 2: "b", 3: "c", 4: "d", 5: "e",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Clone(tt.input)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("Clone() = %v, want nil", got)
				}
				return
			}
			if len(got) != len(tt.wantVals) {
				t.Fatalf("len = %d, want %d", len(got), len(tt.wantVals))
			}
			for k, want := range tt.wantVals {
				if g, ok := got[k]; !ok || g != want {
					t.Fatalf("Clone()[%d] = %q, ok=%v; want %q", k, g, ok, want)
				}
			}
			for k := range got {
				if _, ok := tt.wantVals[k]; !ok {
					t.Fatalf("Clone() unexpected key %d", k)
				}
			}
		})
	}
}

func TestClone_IndependentFromSource(t *testing.T) {
	t.Run("add key on clone", func(t *testing.T) {
		orig := map[string]int{"a": 1, "b": 2}
		cp := lxmaps.Clone(orig)
		cp["c"] = 3
		if _, ok := orig["c"]; ok {
			t.Fatal("mutating clone added key to original")
		}
	})
	t.Run("change value on clone", func(t *testing.T) {
		orig := map[string]int{"a": 1, "b": 2}
		cp := lxmaps.Clone(orig)
		cp["a"] = 99
		if orig["a"] != 1 {
			t.Fatalf("mutating clone changed original value: orig[a]=%d, want 1", orig["a"])
		}
	})
	t.Run("delete from clone", func(t *testing.T) {
		orig := map[string]int{"a": 1, "b": 2}
		cp := lxmaps.Clone(orig)
		delete(cp, "b")
		if _, ok := orig["b"]; !ok {
			t.Fatal("deleting from clone removed key from original")
		}
	})
}

func TestClone_ShallowCopyPointerValues(t *testing.T) {
	n := 1
	orig := map[string]*int{"x": &n}
	cp := lxmaps.Clone(orig)

	if cp["x"] != orig["x"] {
		t.Fatal("clone should share pointer values (shallow copy)")
	}

	*cp["x"] = 42
	if *orig["x"] != 42 {
		t.Fatal("mutating through clone should affect original *int")
	}
}

func TestClone_EmptyNonNilDistinctFromNil(t *testing.T) {
	empty := map[string]int{}
	got := lxmaps.Clone(empty)
	if got == nil {
		t.Fatal("Clone(empty map) should be non-nil")
	}
	if len(got) != 0 {
		t.Fatalf("len = %d, want 0", len(got))
	}
	got["x"] = 1
	if len(empty) != 0 {
		t.Fatal("clone should not alias original empty map")
	}
}

func TestClone_StringBool(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]bool
		wantNil  bool
		wantVals map[string]bool
	}{
		{
			name:     "nil",
			input:    nil,
			wantNil:  true,
			wantVals: nil,
		},
		{
			name:     "false and true",
			input:    map[string]bool{"f": false, "t": true},
			wantNil:  false,
			wantVals: map[string]bool{"f": false, "t": true},
		},
		{
			name:     "only false values",
			input:    map[string]bool{"a": false, "b": false},
			wantNil:  false,
			wantVals: map[string]bool{"a": false, "b": false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Clone(tt.input)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("Clone() = %v, want nil", got)
				}
				return
			}
			if len(got) != len(tt.wantVals) {
				t.Fatalf("len = %d, want %d", len(got), len(tt.wantVals))
			}
			for k, want := range tt.wantVals {
				if g, ok := got[k]; !ok || g != want {
					t.Fatalf("Clone()[%q] = %v, ok=%v; want %v", k, g, ok, want)
				}
			}
		})
	}
}

func TestClone_StructKey(t *testing.T) {
	type key struct {
		ID   int
		Name string
	}
	tests := []struct {
		name     string
		input    map[key]string
		wantNil  bool
		wantVals map[key]string
	}{
		{
			name:     "nil",
			input:    nil,
			wantNil:  true,
			wantVals: nil,
		},
		{
			name: "two struct keys",
			input: map[key]string{
				{ID: 1, Name: "a"}: "x",
				{ID: 2, Name: "b"}: "y",
			},
			wantNil: false,
			wantVals: map[key]string{
				{ID: 1, Name: "a"}: "x",
				{ID: 2, Name: "b"}: "y",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Clone(tt.input)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("Clone() = %v, want nil", got)
				}
				return
			}
			if len(got) != len(tt.wantVals) {
				t.Fatalf("len = %d, want %d", len(got), len(tt.wantVals))
			}
			for k, want := range tt.wantVals {
				if g, ok := got[k]; !ok || g != want {
					t.Fatalf("Clone()[%+v] = %q, ok=%v; want %q", k, g, ok, want)
				}
			}
		})
	}
}

func TestClone_StructValue(t *testing.T) {
	type item struct {
		N int
		S string
	}
	orig := map[string]item{
		"a": {N: 1, S: "one"},
		"b": {N: 2, S: "two"},
	}
	cp := lxmaps.Clone(orig)
	cp["a"] = item{N: 99, S: "changed"}
	if orig["a"].N != 1 || orig["a"].S != "one" {
		t.Fatalf("mutating struct value in clone changed original: %+v", orig["a"])
	}
}

func TestClone_ShallowCopySliceValues(t *testing.T) {
	orig := map[string][]int{
		"a": {1, 2, 3},
		"b": {4},
	}
	cp := lxmaps.Clone(orig)

	cp["a"][0] = 999
	if orig["a"][0] != 999 {
		t.Fatal("mutating clone slice element should affect original (shared backing array)")
	}

	cp["a"] = []int{7, 8, 9}
	if len(orig["a"]) != 3 || orig["a"][0] != 999 {
		t.Fatal("replacing clone slice value should not replace original slice")
	}
	if len(cp["a"]) != 3 || cp["a"][0] != 7 {
		t.Fatalf("clone got new slice: %+v", cp["a"])
	}
}

func TestClone_ShallowCopyNestedMapValue(t *testing.T) {
	inner := map[string]int{"x": 1}
	orig := map[string]map[string]int{"outer": inner}
	cp := lxmaps.Clone(orig)

	cp["outer"]["x"] = 42
	if orig["outer"]["x"] != 42 {
		t.Fatal("mutating inner map through clone should affect original (shared map value)")
	}
}

func TestClone_MutateOriginalDoesNotAffectClone(t *testing.T) {
	orig := map[string]int{"a": 1, "b": 2}
	cp := lxmaps.Clone(orig)

	orig["c"] = 3
	if _, ok := cp["c"]; ok {
		t.Fatal("adding key to original should not appear in clone")
	}

	delete(orig, "a")
	if _, ok := cp["a"]; !ok {
		t.Fatal("deleting from original should not remove key from clone")
	}

	orig["b"] = 99
	if cp["b"] != 2 {
		t.Fatalf("cp[b] = %d, want 2 (unchanged)", cp["b"])
	}
}

func BenchmarkClone(b *testing.B) {
	m := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
		"f": 6, "g": 7, "h": 8, "i": 9, "j": 10,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lxmaps.Clone(m)
	}
}
