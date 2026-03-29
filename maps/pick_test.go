package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/maps"
)

func TestPick_StringInt(t *testing.T) {
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
			name:     "empty map with keys",
			input:    map[string]int{},
			keys:     []string{"a", "b"},
			wantNil:  false,
			expected: map[string]int{},
		},
		{
			name:     "non-empty map no keys",
			input:    map[string]int{"a": 1, "b": 2},
			keys:     nil,
			wantNil:  false,
			expected: map[string]int{},
		},
		{
			name:     "pick subset",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			keys:     []string{"a", "c"},
			wantNil:  false,
			expected: map[string]int{"a": 1, "c": 3},
		},
		{
			name:     "pick missing keys ignored",
			input:    map[string]int{"a": 1, "b": 2},
			keys:     []string{"a", "x", "y"},
			wantNil:  false,
			expected: map[string]int{"a": 1},
		},
		{
			name:     "pick only missing yields empty",
			input:    map[string]int{"a": 1},
			keys:     []string{"z"},
			wantNil:  false,
			expected: map[string]int{},
		},
		{
			name:     "duplicate key args",
			input:    map[string]int{"a": 1, "b": 2},
			keys:     []string{"a", "a", "b"},
			wantNil:  false,
			expected: map[string]int{"a": 1, "b": 2},
		},
		{
			name:     "pick all keys",
			input:    map[string]int{"a": 1, "b": 2},
			keys:     []string{"a", "b"},
			wantNil:  false,
			expected: map[string]int{"a": 1, "b": 2},
		},
		{
			name:     "zero value preserved",
			input:    map[string]int{"z": 0, "a": 1},
			keys:     []string{"z"},
			wantNil:  false,
			expected: map[string]int{"z": 0},
		},
		{
			name:     "negative values",
			input:    map[string]int{"a": -1, "b": -2, "c": 3},
			keys:     []string{"a", "b"},
			wantNil:  false,
			expected: map[string]int{"a": -1, "b": -2},
		},
		{
			name:     "empty string key",
			input:    map[string]int{"": 42, "a": 1},
			keys:     []string{"", "a"},
			wantNil:  false,
			expected: map[string]int{"": 42, "a": 1},
		},
		{
			name:     "unicode keys",
			input:    map[string]int{"こんにちは": 1, "世界": 2, "en": 3},
			keys:     []string{"こんにちは", "en"},
			wantNil:  false,
			expected: map[string]int{"こんにちは": 1, "en": 3},
		},
		{
			name:     "special character keys",
			input:    map[string]int{"!@#": 1, "$%": 2, "normal": 3},
			keys:     []string{"!@#", "normal"},
			wantNil:  false,
			expected: map[string]int{"!@#": 1, "normal": 3},
		},
		{
			name:     "emoji keys",
			input:    map[string]int{"😊": 1, "🚀": 2, "plain": 3},
			keys:     []string{"😊", "plain"},
			wantNil:  false,
			expected: map[string]int{"😊": 1, "plain": 3},
		},
		{
			name: "many keys pick subset",
			input: map[string]int{
				"k1": 1, "k2": 2, "k3": 3, "k4": 4, "k5": 5,
				"k6": 6, "k7": 7,
			},
			keys:    []string{"k1", "k4", "k7", "missing"},
			wantNil: false,
			expected: map[string]int{
				"k1": 1, "k4": 4, "k7": 7,
			},
		},
		{
			name:     "case sensitive keys",
			input:    map[string]int{"A": 1, "a": 2},
			keys:     []string{"A"},
			wantNil:  false,
			expected: map[string]int{"A": 1},
		},
		{
			name:     "large int value",
			input:    map[string]int{"big": 1_000_000, "small": 1},
			keys:     []string{"big"},
			wantNil:  false,
			expected: map[string]int{"big": 1_000_000},
		},
		{
			name: "max int value",
			input: map[string]int{
				"x": int(^uint(0) >> 1),
				"y": 0,
			},
			keys:    []string{"x"},
			wantNil: false,
			expected: map[string]int{
				"x": int(^uint(0) >> 1),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Pick(tt.input, tt.keys...)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("Pick() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("Pick() = nil, want non-nil map")
			}
			if len(got) != len(tt.expected) {
				t.Fatalf("len(Pick()) = %d, want %d", len(got), len(tt.expected))
			}
			for k, want := range tt.expected {
				if g, ok := got[k]; !ok || g != want {
					t.Fatalf("Pick()[%q] = %v, ok=%v; want %v", k, g, ok, want)
				}
			}
			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("Pick() unexpected key %q", k)
				}
			}
		})
	}
}

func TestPick_IntString(t *testing.T) {
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
			keys:     []int{1, 2},
			wantNil:  false,
			expected: map[int]string{},
		},
		{
			name:     "empty map no keys arg",
			input:    map[int]string{1: "a"},
			keys:     nil,
			wantNil:  false,
			expected: map[int]string{},
		},
		{
			name:     "pick subset missing key",
			input:    map[int]string{1: "a", 2: "b", 3: "c"},
			keys:     []int{2, 99, 1},
			wantNil:  false,
			expected: map[int]string{1: "a", 2: "b"},
		},
		{
			name:     "pick only missing",
			input:    map[int]string{1: "a"},
			keys:     []int{0, -1},
			wantNil:  false,
			expected: map[int]string{},
		},
		{
			name:     "zero key",
			input:    map[int]string{0: "zero", 1: "one"},
			keys:     []int{0},
			wantNil:  false,
			expected: map[int]string{0: "zero"},
		},
		{
			name:     "negative keys",
			input:    map[int]string{-2: "neg2", -1: "neg1", 1: "pos"},
			keys:     []int{-2, 1},
			wantNil:  false,
			expected: map[int]string{-2: "neg2", 1: "pos"},
		},
		{
			name:     "large keys",
			input:    map[int]string{1_000_000: "big", 2: "small"},
			keys:     []int{1_000_000, 2},
			wantNil:  false,
			expected: map[int]string{1_000_000: "big", 2: "small"},
		},
		{
			name:     "unicode values",
			input:    map[int]string{1: "hello", 2: "こんにちは"},
			keys:     []int{2},
			wantNil:  false,
			expected: map[int]string{2: "こんにちは"},
		},
		{
			name:     "emoji values",
			input:    map[int]string{1: "😊", 2: "🚀"},
			keys:     []int{1, 2},
			wantNil:  false,
			expected: map[int]string{1: "😊", 2: "🚀"},
		},
		{
			name:     "duplicate key args",
			input:    map[int]string{1: "a", 2: "b"},
			keys:     []int{1, 1, 2},
			wantNil:  false,
			expected: map[int]string{1: "a", 2: "b"},
		},
		{
			name:     "empty string value",
			input:    map[int]string{1: "", 2: "x"},
			keys:     []int{1},
			wantNil:  false,
			expected: map[int]string{1: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Pick(tt.input, tt.keys...)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("Pick() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("Pick() = nil, want non-nil map")
			}
			if len(got) != len(tt.expected) {
				t.Fatalf("len(Pick()) = %d, want %d", len(got), len(tt.expected))
			}
			for k, want := range tt.expected {
				if g, ok := got[k]; !ok || g != want {
					t.Fatalf("Pick()[%d] = %q, ok=%v; want %q", k, g, ok, want)
				}
			}
			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("Pick() unexpected key %d", k)
				}
			}
		})
	}
}

func TestPick_StringBool(t *testing.T) {
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
			keys:     []string{"t"},
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
			name:     "empty map with keys",
			input:    map[string]bool{},
			keys:     []string{"a", "b"},
			wantNil:  false,
			expected: map[string]bool{},
		},
		{
			name:     "non-empty map no keys",
			input:    map[string]bool{"a": true, "b": false},
			keys:     nil,
			wantNil:  false,
			expected: map[string]bool{},
		},
		{
			name:     "pick true and false",
			input:    map[string]bool{"t": true, "f": false, "other": true},
			keys:     []string{"t", "f"},
			wantNil:  false,
			expected: map[string]bool{"t": true, "f": false},
		},
		{
			name:     "false value preserved",
			input:    map[string]bool{"a": false},
			keys:     []string{"a"},
			wantNil:  false,
			expected: map[string]bool{"a": false},
		},
		{
			name:     "pick missing keys ignored",
			input:    map[string]bool{"yes": true, "no": false},
			keys:     []string{"yes", "missing", "nope"},
			wantNil:  false,
			expected: map[string]bool{"yes": true},
		},
		{
			name:     "pick only missing yields empty",
			input:    map[string]bool{"x": true},
			keys:     []string{"y", "z"},
			wantNil:  false,
			expected: map[string]bool{},
		},
		{
			name:     "duplicate key args",
			input:    map[string]bool{"a": true, "b": false},
			keys:     []string{"a", "a", "b"},
			wantNil:  false,
			expected: map[string]bool{"a": true, "b": false},
		},
		{
			name:     "pick all keys",
			input:    map[string]bool{"a": true, "b": false},
			keys:     []string{"a", "b"},
			wantNil:  false,
			expected: map[string]bool{"a": true, "b": false},
		},
		{
			name:     "empty string key",
			input:    map[string]bool{"": true, "a": false},
			keys:     []string{"", "a"},
			wantNil:  false,
			expected: map[string]bool{"": true, "a": false},
		},
		{
			name:     "unicode keys",
			input:    map[string]bool{"はい": true, "いいえ": false, "en": true},
			keys:     []string{"はい", "en"},
			wantNil:  false,
			expected: map[string]bool{"はい": true, "en": true},
		},
		{
			name:     "special character keys",
			input:    map[string]bool{"!@#": true, "$%": false, "ok": true},
			keys:     []string{"!@#", "ok"},
			wantNil:  false,
			expected: map[string]bool{"!@#": true, "ok": true},
		},
		{
			name:     "emoji keys",
			input:    map[string]bool{"✓": true, "✗": false, "plain": true},
			keys:     []string{"✓", "plain"},
			wantNil:  false,
			expected: map[string]bool{"✓": true, "plain": true},
		},
		{
			name:     "case sensitive keys",
			input:    map[string]bool{"T": true, "t": false},
			keys:     []string{"T"},
			wantNil:  false,
			expected: map[string]bool{"T": true},
		},
		{
			name: "many keys pick subset",
			input: map[string]bool{
				"k1": true, "k2": false, "k3": true, "k4": false, "k5": true,
			},
			keys:    []string{"k1", "k4", "k5", "absent"},
			wantNil: false,
			expected: map[string]bool{
				"k1": true, "k4": false, "k5": true,
			},
		},
		{
			name:     "all true values",
			input:    map[string]bool{"a": true, "b": true},
			keys:     []string{"a"},
			wantNil:  false,
			expected: map[string]bool{"a": true},
		},
		{
			name:     "all false values",
			input:    map[string]bool{"a": false, "b": false},
			keys:     []string{"b"},
			wantNil:  false,
			expected: map[string]bool{"b": false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Pick(tt.input, tt.keys...)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("Pick() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("Pick() = nil, want non-nil map")
			}
			if len(got) != len(tt.expected) {
				t.Fatalf("len(Pick()) = %d, want %d", len(got), len(tt.expected))
			}
			for k, want := range tt.expected {
				if g, ok := got[k]; !ok || g != want {
					t.Fatalf("Pick()[%q] = %v, ok=%v; want %v", k, g, ok, want)
				}
			}
			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("Pick() unexpected key %q", k)
				}
			}
		})
	}
}

func TestPick_StringStruct(t *testing.T) {
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
			name:     "empty map with keys",
			input:    map[string]User{},
			keys:     []string{"a", "b"},
			wantNil:  false,
			expected: map[string]User{},
		},
		{
			name: "non-empty no keys",
			input: map[string]User{
				"x": {Name: "X", Age: 1},
			},
			keys:     nil,
			wantNil:  false,
			expected: map[string]User{},
		},
		{
			name: "pick subset",
			input: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
			keys:    []string{"bob"},
			wantNil: false,
			expected: map[string]User{
				"bob": {Name: "Bob", Age: 30},
			},
		},
		{
			name: "pick zero-value struct",
			input: map[string]User{
				"empty": {},
				"full":  {Name: "X", Age: 1},
			},
			keys:    []string{"empty"},
			wantNil: false,
			expected: map[string]User{
				"empty": {},
			},
		},
		{
			name: "pick missing keys ignored",
			input: map[string]User{
				"a": {Name: "A", Age: 1},
				"b": {Name: "B", Age: 2},
			},
			keys:    []string{"a", "ghost"},
			wantNil: false,
			expected: map[string]User{
				"a": {Name: "A", Age: 1},
			},
		},
		{
			name: "pick only missing yields empty",
			input: map[string]User{
				"only": {Name: "O", Age: 0},
			},
			keys:     []string{"nope"},
			wantNil:  false,
			expected: map[string]User{},
		},
		{
			name: "duplicate key args",
			input: map[string]User{
				"a": {Name: "A", Age: 1},
				"b": {Name: "B", Age: 2},
			},
			keys:    []string{"a", "a", "b"},
			wantNil: false,
			expected: map[string]User{
				"a": {Name: "A", Age: 1},
				"b": {Name: "B", Age: 2},
			},
		},
		{
			name: "pick all",
			input: map[string]User{
				"p": {Name: "Pat", Age: 40},
				"q": {Name: "Quinn", Age: 41},
			},
			keys:    []string{"p", "q"},
			wantNil: false,
			expected: map[string]User{
				"p": {Name: "Pat", Age: 40},
				"q": {Name: "Quinn", Age: 41},
			},
		},
		{
			name: "unicode name field",
			input: map[string]User{
				"jp": {Name: "太郎", Age: 20},
				"en": {Name: "Sam", Age: 21},
			},
			keys:    []string{"jp"},
			wantNil: false,
			expected: map[string]User{
				"jp": {Name: "太郎", Age: 20},
			},
		},
		{
			name: "unicode key",
			input: map[string]User{
				"ユーザー":  {Name: "U", Age: 1},
				"other": {Name: "O", Age: 2},
			},
			keys:    []string{"ユーザー", "other"},
			wantNil: false,
			expected: map[string]User{
				"ユーザー":  {Name: "U", Age: 1},
				"other": {Name: "O", Age: 2},
			},
		},
		{
			name: "empty string key",
			input: map[string]User{
				"":  {Name: "BlankKey", Age: 0},
				"x": {Name: "X", Age: 1},
			},
			keys:    []string{"", "x"},
			wantNil: false,
			expected: map[string]User{
				"":  {Name: "BlankKey", Age: 0},
				"x": {Name: "X", Age: 1},
			},
		},
		{
			name: "special character keys",
			input: map[string]User{
				"id:1": {Name: "One", Age: 1},
				"id:2": {Name: "Two", Age: 2},
			},
			keys:    []string{"id:1"},
			wantNil: false,
			expected: map[string]User{
				"id:1": {Name: "One", Age: 1},
			},
		},
		{
			name: "emoji key",
			input: map[string]User{
				"👤": {Name: "Emoji", Age: 3},
				"a": {Name: "Plain", Age: 4},
			},
			keys:    []string{"👤"},
			wantNil: false,
			expected: map[string]User{
				"👤": {Name: "Emoji", Age: 3},
			},
		},
		{
			name: "case sensitive keys",
			input: map[string]User{
				"User": {Name: "Upper", Age: 1},
				"user": {Name: "Lower", Age: 2},
			},
			keys:    []string{"user"},
			wantNil: false,
			expected: map[string]User{
				"user": {Name: "Lower", Age: 2},
			},
		},
		{
			name: "negative age preserved",
			input: map[string]User{
				"a": {Name: "A", Age: -1},
				"b": {Name: "B", Age: 10},
			},
			keys:    []string{"a"},
			wantNil: false,
			expected: map[string]User{
				"a": {Name: "A", Age: -1},
			},
		},
		{
			name: "many entries pick subset",
			input: map[string]User{
				"k1": {Name: "N1", Age: 1},
				"k2": {Name: "N2", Age: 2},
				"k3": {Name: "N3", Age: 3},
				"k4": {Name: "N4", Age: 4},
				"k5": {Name: "N5", Age: 5},
			},
			keys:    []string{"k1", "k4", "missing"},
			wantNil: false,
			expected: map[string]User{
				"k1": {Name: "N1", Age: 1},
				"k4": {Name: "N4", Age: 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Pick(tt.input, tt.keys...)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("Pick() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("Pick() = nil, want non-nil map")
			}
			if len(got) != len(tt.expected) {
				t.Fatalf("len(Pick()) = %d, want %d", len(got), len(tt.expected))
			}
			for k, want := range tt.expected {
				g, ok := got[k]
				if !ok || g != want {
					t.Fatalf("Pick()[%q] = %+v, ok=%v; want %+v", k, g, ok, want)
				}
			}
			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("Pick() unexpected key %q", k)
				}
			}
		})
	}
}

func BenchmarkPick(b *testing.B) {
	m := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
	}
	keys := []string{"a", "c", "e"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lxmaps.Pick(m, keys...)
	}
}
