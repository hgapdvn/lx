package lxmaps_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/maps"
)

func TestOmit_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		keys     []string
		wantNil  bool
		expected map[string]int
	}{
		{
			name:     "nil map",
			input:    nil,
			keys:     []string{"a"},
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "nil map empty keys",
			input:    nil,
			keys:     nil,
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "empty map no keys",
			input:    map[string]int{},
			keys:     nil,
			wantNil:  false,
			expected: map[string]int{},
		},
		{
			name:     "empty map omit keys",
			input:    map[string]int{},
			keys:     []string{"a"},
			wantNil:  false,
			expected: map[string]int{},
		},
		{
			name:     "no keys to omit is full clone",
			input:    map[string]int{"a": 1, "b": 2},
			keys:     nil,
			wantNil:  false,
			expected: map[string]int{"a": 1, "b": 2},
		},
		{
			name:     "omit subset",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			keys:     []string{"b"},
			wantNil:  false,
			expected: map[string]int{"a": 1, "c": 3},
		},
		{
			name:     "omit missing keys no op",
			input:    map[string]int{"a": 1, "b": 2},
			keys:     []string{"x", "y"},
			wantNil:  false,
			expected: map[string]int{"a": 1, "b": 2},
		},
		{
			name:     "omit all keys",
			input:    map[string]int{"a": 1, "b": 2},
			keys:     []string{"a", "b"},
			wantNil:  false,
			expected: map[string]int{},
		},
		{
			name:     "duplicate omit args",
			input:    map[string]int{"a": 1, "b": 2},
			keys:     []string{"a", "a"},
			wantNil:  false,
			expected: map[string]int{"b": 2},
		},
		{
			name:     "omit single",
			input:    map[string]int{"only": 42},
			keys:     []string{"only"},
			wantNil:  false,
			expected: map[string]int{},
		},
		{
			name:     "omit multiple keys",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			keys:     []string{"b", "d"},
			wantNil:  false,
			expected: map[string]int{"a": 1, "c": 3},
		},
		{
			name:     "negative values",
			input:    map[string]int{"a": -1, "b": -2, "c": 3},
			keys:     []string{"c"},
			wantNil:  false,
			expected: map[string]int{"a": -1, "b": -2},
		},
		{
			name:     "empty string key",
			input:    map[string]int{"": 42, "a": 1},
			keys:     []string{""},
			wantNil:  false,
			expected: map[string]int{"a": 1},
		},
		{
			name:     "unicode keys",
			input:    map[string]int{"こんにちは": 1, "世界": 2, "en": 3},
			keys:     []string{"世界"},
			wantNil:  false,
			expected: map[string]int{"こんにちは": 1, "en": 3},
		},
		{
			name:     "special character keys",
			input:    map[string]int{"!@#": 1, "$%": 2, "normal": 3},
			keys:     []string{"$%"},
			wantNil:  false,
			expected: map[string]int{"!@#": 1, "normal": 3},
		},
		{
			name:     "emoji keys",
			input:    map[string]int{"😊": 1, "🚀": 2, "plain": 3},
			keys:     []string{"🚀"},
			wantNil:  false,
			expected: map[string]int{"😊": 1, "plain": 3},
		},
		{
			name: "many entries omit several",
			input: map[string]int{
				"k1": 1, "k2": 2, "k3": 3, "k4": 4, "k5": 5,
				"k6": 6, "k7": 7,
			},
			keys:    []string{"k2", "k4", "k6", "missing"},
			wantNil: false,
			expected: map[string]int{
				"k1": 1, "k3": 3, "k5": 5, "k7": 7,
			},
		},
		{
			name:     "case sensitive keys",
			input:    map[string]int{"A": 1, "a": 2},
			keys:     []string{"a"},
			wantNil:  false,
			expected: map[string]int{"A": 1},
		},
		{
			name:     "large int value",
			input:    map[string]int{"big": 1_000_000, "small": 1},
			keys:     []string{"small"},
			wantNil:  false,
			expected: map[string]int{"big": 1_000_000},
		},
		{
			name: "max int value",
			input: map[string]int{
				"x": int(^uint(0) >> 1),
				"y": 1,
			},
			keys:    []string{"y"},
			wantNil: false,
			expected: map[string]int{
				"x": int(^uint(0) >> 1),
			},
		},
		{
			name:     "zero value preserved after omit",
			input:    map[string]int{"z": 0, "a": 1},
			keys:     []string{"a"},
			wantNil:  false,
			expected: map[string]int{"z": 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Omit(tt.input, tt.keys...)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("Omit() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("Omit() = nil, want non-nil map")
			}
			if len(got) != len(tt.expected) {
				t.Fatalf("len(Omit()) = %d, want %d", len(got), len(tt.expected))
			}
			for k, want := range tt.expected {
				if g, ok := got[k]; !ok || g != want {
					t.Fatalf("Omit()[%q] = %v, ok=%v; want %v", k, g, ok, want)
				}
			}
			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("Omit() unexpected key %q", k)
				}
			}
		})
	}
}

func TestOmit_IntString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		keys     []int
		wantNil  bool
		expected map[int]string
	}{
		{
			name:     "nil map",
			input:    nil,
			keys:     []int{1},
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "nil map empty keys",
			input:    nil,
			keys:     nil,
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "empty map",
			input:    map[int]string{},
			keys:     []int{1},
			wantNil:  false,
			expected: map[int]string{},
		},
		{
			name:     "no keys full copy",
			input:    map[int]string{1: "a", 2: "b"},
			keys:     nil,
			wantNil:  false,
			expected: map[int]string{1: "a", 2: "b"},
		},
		{
			name:     "omit one key",
			input:    map[int]string{1: "a", 2: "b", 3: "c"},
			keys:     []int{2, 99},
			wantNil:  false,
			expected: map[int]string{1: "a", 3: "c"},
		},
		{
			name:     "omit all",
			input:    map[int]string{1: "a", 2: "b"},
			keys:     []int{1, 2},
			wantNil:  false,
			expected: map[int]string{},
		},
		{
			name:     "zero key",
			input:    map[int]string{0: "zero", 1: "one"},
			keys:     []int{0},
			wantNil:  false,
			expected: map[int]string{1: "one"},
		},
		{
			name:     "negative keys",
			input:    map[int]string{-2: "a", -1: "b", 1: "c"},
			keys:     []int{-1},
			wantNil:  false,
			expected: map[int]string{-2: "a", 1: "c"},
		},
		{
			name:     "large keys",
			input:    map[int]string{1_000_000: "big", 2: "small"},
			keys:     []int{2},
			wantNil:  false,
			expected: map[int]string{1_000_000: "big"},
		},
		{
			name:     "duplicate omit args",
			input:    map[int]string{1: "a", 2: "b"},
			keys:     []int{1, 1},
			wantNil:  false,
			expected: map[int]string{2: "b"},
		},
		{
			name:     "unicode values",
			input:    map[int]string{1: "hello", 2: "こんにちは"},
			keys:     []int{1},
			wantNil:  false,
			expected: map[int]string{2: "こんにちは"},
		},
		{
			name:     "emoji values",
			input:    map[int]string{1: "😊", 2: "🚀", 3: "x"},
			keys:     []int{2},
			wantNil:  false,
			expected: map[int]string{1: "😊", 3: "x"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Omit(tt.input, tt.keys...)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("Omit() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("Omit() = nil, want non-nil map")
			}
			if len(got) != len(tt.expected) {
				t.Fatalf("len(Omit()) = %d, want %d", len(got), len(tt.expected))
			}
			for k, want := range tt.expected {
				if g, ok := got[k]; !ok || g != want {
					t.Fatalf("Omit()[%d] = %q, ok=%v; want %q", k, g, ok, want)
				}
			}
			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("Omit() unexpected key %d", k)
				}
			}
		})
	}
}

func TestOmit_StringBool(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]bool
		keys     []string
		wantNil  bool
		expected map[string]bool
	}{
		{
			name:     "nil map with keys",
			input:    nil,
			keys:     []string{"a"},
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "nil map empty keys",
			input:    nil,
			keys:     nil,
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "empty map no keys",
			input:    map[string]bool{},
			keys:     nil,
			wantNil:  false,
			expected: map[string]bool{},
		},
		{
			name:     "empty map omit keys",
			input:    map[string]bool{},
			keys:     []string{"x"},
			wantNil:  false,
			expected: map[string]bool{},
		},
		{
			name:     "no keys full clone",
			input:    map[string]bool{"t": true, "f": false},
			keys:     nil,
			wantNil:  false,
			expected: map[string]bool{"t": true, "f": false},
		},
		{
			name:     "omit false key keeps true",
			input:    map[string]bool{"t": true, "f": false},
			keys:     []string{"f"},
			wantNil:  false,
			expected: map[string]bool{"t": true},
		},
		{
			name:     "omit leaves false value",
			input:    map[string]bool{"a": false, "b": true},
			keys:     []string{"b"},
			wantNil:  false,
			expected: map[string]bool{"a": false},
		},
		{
			name:     "omit missing keys no op",
			input:    map[string]bool{"a": true, "b": false},
			keys:     []string{"ghost", "missing"},
			wantNil:  false,
			expected: map[string]bool{"a": true, "b": false},
		},
		{
			name:     "omit all keys",
			input:    map[string]bool{"a": true, "b": false},
			keys:     []string{"a", "b"},
			wantNil:  false,
			expected: map[string]bool{},
		},
		{
			name:     "duplicate omit args",
			input:    map[string]bool{"a": true, "b": false},
			keys:     []string{"a", "a"},
			wantNil:  false,
			expected: map[string]bool{"b": false},
		},
		{
			name:     "omit multiple",
			input:    map[string]bool{"a": true, "b": false, "c": true, "d": false},
			keys:     []string{"b", "d"},
			wantNil:  false,
			expected: map[string]bool{"a": true, "c": true},
		},
		{
			name:     "empty string key",
			input:    map[string]bool{"": true, "a": false},
			keys:     []string{""},
			wantNil:  false,
			expected: map[string]bool{"a": false},
		},
		{
			name:     "unicode keys",
			input:    map[string]bool{"はい": true, "いいえ": false, "en": true},
			keys:     []string{"いいえ"},
			wantNil:  false,
			expected: map[string]bool{"はい": true, "en": true},
		},
		{
			name:     "special character keys",
			input:    map[string]bool{"!@#": true, "$%": false},
			keys:     []string{"!@#"},
			wantNil:  false,
			expected: map[string]bool{"$%": false},
		},
		{
			name:     "emoji keys",
			input:    map[string]bool{"✓": true, "✗": false, "plain": true},
			keys:     []string{"✗"},
			wantNil:  false,
			expected: map[string]bool{"✓": true, "plain": true},
		},
		{
			name:     "case sensitive keys",
			input:    map[string]bool{"T": true, "t": false},
			keys:     []string{"t"},
			wantNil:  false,
			expected: map[string]bool{"T": true},
		},
		{
			name: "many entries omit several",
			input: map[string]bool{
				"k1": true, "k2": false, "k3": true, "k4": false, "k5": true,
			},
			keys:    []string{"k2", "k4", "absent"},
			wantNil: false,
			expected: map[string]bool{
				"k1": true, "k3": true, "k5": true,
			},
		},
		{
			name:     "omit single entry map",
			input:    map[string]bool{"only": true},
			keys:     []string{"only"},
			wantNil:  false,
			expected: map[string]bool{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Omit(tt.input, tt.keys...)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("Omit() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("Omit() = nil, want non-nil map")
			}
			if len(got) != len(tt.expected) {
				t.Fatalf("len(Omit()) = %d, want %d", len(got), len(tt.expected))
			}
			for k, want := range tt.expected {
				if g, ok := got[k]; !ok || g != want {
					t.Fatalf("Omit()[%q] = %v, ok=%v; want %v", k, g, ok, want)
				}
			}
			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("Omit() unexpected key %q", k)
				}
			}
		})
	}
}

func TestOmit_StringStruct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	tests := []struct {
		name     string
		input    map[string]User
		keys     []string
		wantNil  bool
		expected map[string]User
	}{
		{
			name:     "nil map with keys",
			input:    nil,
			keys:     []string{"alice"},
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "nil map empty keys",
			input:    nil,
			keys:     nil,
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "empty map no keys",
			input:    map[string]User{},
			keys:     nil,
			wantNil:  false,
			expected: map[string]User{},
		},
		{
			name:     "empty map omit keys",
			input:    map[string]User{},
			keys:     []string{"a"},
			wantNil:  false,
			expected: map[string]User{},
		},
		{
			name: "omit one user",
			input: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
			keys:    []string{"alice"},
			wantNil: false,
			expected: map[string]User{
				"bob": {Name: "Bob", Age: 30},
			},
		},
		{
			name: "empty keys clones all",
			input: map[string]User{
				"a": {Name: "A", Age: 1},
			},
			keys:    nil,
			wantNil: false,
			expected: map[string]User{
				"a": {Name: "A", Age: 1},
			},
		},
		{
			name: "omit missing keys no op",
			input: map[string]User{
				"a": {Name: "A", Age: 1},
				"b": {Name: "B", Age: 2},
			},
			keys:    []string{"ghost", "phantom"},
			wantNil: false,
			expected: map[string]User{
				"a": {Name: "A", Age: 1},
				"b": {Name: "B", Age: 2},
			},
		},
		{
			name: "omit all keys",
			input: map[string]User{
				"x": {Name: "X", Age: 1},
				"y": {Name: "Y", Age: 2},
			},
			keys:     []string{"x", "y"},
			wantNil:  false,
			expected: map[string]User{},
		},
		{
			name: "duplicate omit args",
			input: map[string]User{
				"a": {Name: "A", Age: 1},
				"b": {Name: "B", Age: 2},
			},
			keys:    []string{"a", "a"},
			wantNil: false,
			expected: map[string]User{
				"b": {Name: "B", Age: 2},
			},
		},
		{
			name: "omit multiple",
			input: map[string]User{
				"k1": {Name: "N1", Age: 1},
				"k2": {Name: "N2", Age: 2},
				"k3": {Name: "N3", Age: 3},
				"k4": {Name: "N4", Age: 4},
			},
			keys:    []string{"k2", "k4"},
			wantNil: false,
			expected: map[string]User{
				"k1": {Name: "N1", Age: 1},
				"k3": {Name: "N3", Age: 3},
			},
		},
		{
			name: "unicode name in value",
			input: map[string]User{
				"jp": {Name: "太郎", Age: 20},
				"en": {Name: "Sam", Age: 21},
			},
			keys:    []string{"en"},
			wantNil: false,
			expected: map[string]User{
				"jp": {Name: "太郎", Age: 20},
			},
		},
		{
			name: "unicode key",
			input: map[string]User{
				"ユーザー":  {Name: "U", Age: 1},
				"admin": {Name: "A", Age: 2},
			},
			keys:    []string{"ユーザー"},
			wantNil: false,
			expected: map[string]User{
				"admin": {Name: "A", Age: 2},
			},
		},
		{
			name: "empty string key",
			input: map[string]User{
				"":  {Name: "Blank", Age: 0},
				"x": {Name: "X", Age: 1},
			},
			keys:    []string{""},
			wantNil: false,
			expected: map[string]User{
				"x": {Name: "X", Age: 1},
			},
		},
		{
			name: "special character keys",
			input: map[string]User{
				"id:1": {Name: "One", Age: 1},
				"id:2": {Name: "Two", Age: 2},
			},
			keys:    []string{"id:2"},
			wantNil: false,
			expected: map[string]User{
				"id:1": {Name: "One", Age: 1},
			},
		},
		{
			name: "emoji key",
			input: map[string]User{
				"👤": {Name: "E", Age: 3},
				"a": {Name: "P", Age: 4},
			},
			keys:    []string{"👤"},
			wantNil: false,
			expected: map[string]User{
				"a": {Name: "P", Age: 4},
			},
		},
		{
			name: "case sensitive keys",
			input: map[string]User{
				"User": {Name: "Upper", Age: 1},
				"user": {Name: "Lower", Age: 2},
			},
			keys:    []string{"User"},
			wantNil: false,
			expected: map[string]User{
				"user": {Name: "Lower", Age: 2},
			},
		},
		{
			name: "zero-value struct remains after omit",
			input: map[string]User{
				"empty": {},
				"full":  {Name: "F", Age: 1},
			},
			keys:    []string{"full"},
			wantNil: false,
			expected: map[string]User{
				"empty": {},
			},
		},
		{
			name: "empty keys clones multiple",
			input: map[string]User{
				"p": {Name: "P", Age: 1},
				"q": {Name: "Q", Age: 2},
			},
			keys:    nil,
			wantNil: false,
			expected: map[string]User{
				"p": {Name: "P", Age: 1},
				"q": {Name: "Q", Age: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Omit(tt.input, tt.keys...)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("Omit() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("Omit() = nil, want non-nil map")
			}
			if len(got) != len(tt.expected) {
				t.Fatalf("len(Omit()) = %d, want %d", len(got), len(tt.expected))
			}
			for k, want := range tt.expected {
				g, ok := got[k]
				if !ok || g != want {
					t.Fatalf("Omit()[%q] = %+v, ok=%v; want %+v", k, g, ok, want)
				}
			}
			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("Omit() unexpected key %q", k)
				}
			}
		})
	}
}

func TestOmit_EmptyKeysMatchesClone(t *testing.T) {
	t.Run("nil map", func(t *testing.T) {
		var m map[string]int
		if lxmaps.Omit(m) != nil {
			t.Fatal("Omit(nil) with no keys to omit = non-nil, want nil")
		}
		if lxmaps.Clone(m) != nil {
			t.Fatal("Clone(nil) = non-nil, want nil")
		}
	})
	t.Run("non-nil map", func(t *testing.T) {
		orig := map[string]int{"a": 1, "b": 2}
		omitOut := lxmaps.Omit(orig)
		cloneOut := lxmaps.Clone(orig)
		if len(omitOut) != len(cloneOut) {
			t.Fatalf("len Omit = %d, len Clone = %d", len(omitOut), len(cloneOut))
		}
		for k, v := range cloneOut {
			if omitOut[k] != v {
				t.Fatalf("Omit()[%q] = %d, Clone()[%q] = %d", k, omitOut[k], k, v)
			}
		}
		for k := range omitOut {
			if _, ok := cloneOut[k]; !ok {
				t.Fatalf("Omit() unexpected key %q", k)
			}
		}
	})
}

func BenchmarkOmit(b *testing.B) {
	m := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
	}
	keys := []string{"b", "d"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lxmaps.Omit(m, keys...)
	}
}
