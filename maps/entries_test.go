package lxmaps_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/lxslices"
	"github.com/nthanhhai2909/lx/lxtypes"
	"github.com/nthanhhai2909/lx/maps"
)

func TestEntries_String(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[string]int
		expected []lxtypes.Pair[string, int]
	}{
		{"no args", nil, nil},
		{"single nil map", []map[string]int{nil}, []lxtypes.Pair[string, int]{}},
		{"single", []map[string]int{{"a": 1}}, []lxtypes.Pair[string, int]{{First: "a", Second: 1}}},
		{"multi maps duplicated pair", []map[string]int{{"a": 1, "b": 2}, {"b": 3, "c": 4}}, []lxtypes.Pair[string, int]{{First: "a", Second: 1}, {First: "b", Second: 2}, {First: "b", Second: 3}, {First: "c", Second: 4}}},
		{"empty + map", []map[string]int{{}, {"x": 1}}, []lxtypes.Pair[string, int]{{First: "x", Second: 1}}},
		{"unicode pairs", []map[string]int{{"こんにちは": 10}, {"世界": 20}}, []lxtypes.Pair[string, int]{{First: "こんにちは", Second: 10}, {First: "世界", Second: 20}}},
		{"emoji pairs", []map[string]int{{"😊": 5}, {"🚀": 6}}, []lxtypes.Pair[string, int]{{First: "😊", Second: 5}, {First: "🚀", Second: 6}}},
		{"many small", []map[string]int{{"k1": 1, "k2": 2, "k3": 3}}, []lxtypes.Pair[string, int]{{First: "k1", Second: 1}, {First: "k2", Second: 2}, {First: "k3", Second: 3}}},
		{"special chars", []map[string]int{{"!@#": 7}, {"$%": 8}}, []lxtypes.Pair[string, int]{{First: "!@#", Second: 7}, {First: "$%", Second: 8}}},
		{"duplicate keys across maps", []map[string]int{{"dup": 1}, {"dup": 2}}, []lxtypes.Pair[string, int]{{First: "dup", Second: 1}, {First: "dup", Second: 2}}},
		{"mixed many", []map[string]int{{"p": 100}, {"q": -1}, {"r": 0}}, []lxtypes.Pair[string, int]{{First: "p", Second: 100}, {First: "q", Second: -1}, {First: "r", Second: 0}}},
		{"final small", []map[string]int{{"end": 3}, {"stop": 4}}, []lxtypes.Pair[string, int]{{First: "end", Second: 3}, {First: "stop", Second: 4}}},
		{"many maps combined", []map[string]int{{"m1": 1}, {"m2": 2}, {"m3": 3}}, []lxtypes.Pair[string, int]{{First: "m1", Second: 1}, {First: "m2", Second: 2}, {First: "m3", Second: 3}}},
		{"long keys", []map[string]int{{"long_string_example_0123456789": 1}}, []lxtypes.Pair[string, int]{{First: "long_string_example_0123456789", Second: 1}}},
		{"mixed unicode ascii", []map[string]int{{"a": 1}, {"ä": 2}}, []lxtypes.Pair[string, int]{{First: "a", Second: 1}, {First: "ä", Second: 2}}},
		{"random small", []map[string]int{{"x": 42}, {"y": 7}}, []lxtypes.Pair[string, int]{{First: "x", Second: 42}, {First: "y", Second: 7}}},
		{"extra1", []map[string]int{{"e1": 11}}, []lxtypes.Pair[string, int]{{First: "e1", Second: 11}}},
		{"extra2", []map[string]int{{"e2": 22}}, []lxtypes.Pair[string, int]{{First: "e2", Second: 22}}},
		{"extra3", []map[string]int{{"e3": 33}}, []lxtypes.Pair[string, int]{{First: "e3", Second: 33}}},
		{"final extras", []map[string]int{{"z1": 1}, {"z2": 2}}, []lxtypes.Pair[string, int]{{First: "z1", Second: 1}, {First: "z2", Second: 2}}},
		// Extra cases
		{"many duplicates", []map[string]int{{"d": 1}, {"d": 2}, {"d": 3}}, []lxtypes.Pair[string, int]{{First: "d", Second: 1}, {First: "d", Second: 2}, {First: "d", Second: 3}}},
		{"empty nil mix", []map[string]int{nil, {}, {"k": 9}}, []lxtypes.Pair[string, int]{{First: "k", Second: 9}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Entries(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("Entries(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("Entries(%v) = nil; want non-nil with length %d", tt.input, len(tt.expected))
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("Entries(%v) length = %d; want %d", tt.input, len(got), len(tt.expected))
			}

			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Fatalf("Entries(%v) missing expected pairs; got %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestEntries_Int(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[int]string
		expected []lxtypes.Pair[int, string]
	}{
		{"no args", nil, nil},
		{"single nil map", []map[int]string{nil}, []lxtypes.Pair[int, string]{}},
		{"single", []map[int]string{{1: "one"}}, []lxtypes.Pair[int, string]{{First: 1, Second: "one"}}},
		{"multi maps duplicated pair", []map[int]string{{1: "a", 2: "b"}, {2: "c", 3: "d"}}, []lxtypes.Pair[int, string]{{First: 1, Second: "a"}, {First: 2, Second: "b"}, {First: 2, Second: "c"}, {First: 3, Second: "d"}}},
		{"empty + map", []map[int]string{{}, {7: "y"}}, []lxtypes.Pair[int, string]{{First: 7, Second: "y"}}},
		{"many small", []map[int]string{{1: "k1", 2: "k2", 3: "k3"}}, []lxtypes.Pair[int, string]{{First: 1, Second: "k1"}, {First: 2, Second: "k2"}, {First: 3, Second: "k3"}}},
		{"duplicate keys across maps", []map[int]string{{2: "b"}, {2: "c"}}, []lxtypes.Pair[int, string]{{First: 2, Second: "b"}, {First: 2, Second: "c"}}},
		{"mixed many", []map[int]string{{100: "x"}, {-1: "y"}, {0: "z"}}, []lxtypes.Pair[int, string]{{First: 100, Second: "x"}, {First: -1, Second: "y"}, {First: 0, Second: "z"}}},
		{"final extra", []map[int]string{{11: "k"}, {22: "l"}}, []lxtypes.Pair[int, string]{{First: 11, Second: "k"}, {First: 22, Second: "l"}}},
		// Extra cases
		{"many duplicates", []map[int]string{{5: "a"}, {5: "b"}, {5: "c"}}, []lxtypes.Pair[int, string]{{First: 5, Second: "a"}, {First: 5, Second: "b"}, {First: 5, Second: "c"}}},
		{"empty nil mix", []map[int]string{nil, {}, {9: "k"}}, []lxtypes.Pair[int, string]{{First: 9, Second: "k"}}},
		{"sparse+single", []map[int]string{{2: "a"}, {999: "x"}}, []lxtypes.Pair[int, string]{{First: 2, Second: "a"}, {First: 999, Second: "x"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Entries(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("Entries(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("Entries(%v) = nil; want non-nil with length %d", tt.input, len(tt.expected))
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("Entries(%v) length = %d; want %d", tt.input, len(got), len(tt.expected))
			}

			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Fatalf("Entries(%v) missing expected pairs; got %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestEntries_Struct(t *testing.T) {
	// comparable key and value structs
	type K struct {
		I int
		S string
	}
	type V struct {
		N string
		M int
	}

	tests := []struct {
		name     string
		input    []map[K]V
		expected []lxtypes.Pair[K, V]
	}{
		{"no args", nil, nil},
		{"single nil map", []map[K]V{nil}, []lxtypes.Pair[K, V]{}},
		{"single", []map[K]V{{{I: 1, S: "a"}: {N: "one", M: 1}}}, []lxtypes.Pair[K, V]{{First: K{I: 1, S: "a"}, Second: V{N: "one", M: 1}}}},
		{"multi maps duplicated pair", []map[K]V{{{I: 1, S: "a"}: {N: "one", M: 1}}, {{I: 1, S: "a"}: {N: "uno", M: 11}, {I: 2, S: "b"}: {N: "two", M: 2}}}, []lxtypes.Pair[K, V]{{First: K{I: 1, S: "a"}, Second: V{N: "one", M: 1}}, {First: K{I: 1, S: "a"}, Second: V{N: "uno", M: 11}}, {First: K{I: 2, S: "b"}, Second: V{N: "two", M: 2}}}},
		{"empty + map", []map[K]V{{}, {{I: 3, S: "x"}: {N: "x", M: 3}}}, []lxtypes.Pair[K, V]{{First: K{I: 3, S: "x"}, Second: V{N: "x", M: 3}}}},
		{"many small", []map[K]V{{{I: 1, S: "a"}: {N: "a", M: 1}, {I: 2, S: "b"}: {N: "b", M: 2}}}, []lxtypes.Pair[K, V]{{First: K{I: 1, S: "a"}, Second: V{N: "a", M: 1}}, {First: K{I: 2, S: "b"}, Second: V{N: "b", M: 2}}}},
		{"duplicate pairs", []map[K]V{{{I: 1, S: "dup"}: {N: "a", M: 1}}, {{I: 1, S: "dup"}: {N: "b", M: 2}}}, []lxtypes.Pair[K, V]{{First: K{I: 1, S: "dup"}, Second: V{N: "a", M: 1}}, {First: K{I: 1, S: "dup"}, Second: V{N: "b", M: 2}}}},
		{"unicode pairs", []map[K]V{{{I: 10, S: "こんにちは"}: {N: "u", M: 10}}, {{I: 11, S: "世界"}: {N: "v", M: 11}}}, []lxtypes.Pair[K, V]{{First: K{I: 10, S: "こんにちは"}, Second: V{N: "u", M: 10}}, {First: K{I: 11, S: "世界"}, Second: V{N: "v", M: 11}}}},
		{"emoji pairs", []map[K]V{{{I: 2, S: "😊"}: {N: "e", M: 2}}, {{I: 3, S: "👋"}: {N: "w", M: 3}}}, []lxtypes.Pair[K, V]{{First: K{I: 2, S: "😊"}, Second: V{N: "e", M: 2}}, {First: K{I: 3, S: "👋"}, Second: V{N: "w", M: 3}}}},
		{"mixed many", []map[K]V{{{I: 5, S: "X"}: {N: "X", M: 5}}, {{I: -5, S: "Y"}: {N: "Y", M: -5}}}, []lxtypes.Pair[K, V]{{First: K{I: 5, S: "X"}, Second: V{N: "X", M: 5}}, {First: K{I: -5, S: "Y"}, Second: V{N: "Y", M: -5}}}},
		{"many maps combined", []map[K]V{{{I: 1, S: "a"}: {N: "a", M: 1}}, {{I: 2, S: "b"}: {N: "b", M: 2}}, {{I: 3, S: "c"}: {N: "c", M: 3}}}, []lxtypes.Pair[K, V]{{First: K{I: 1, S: "a"}, Second: V{N: "a", M: 1}}, {First: K{I: 2, S: "b"}, Second: V{N: "b", M: 2}}, {First: K{I: 3, S: "c"}, Second: V{N: "c", M: 3}}}},
		{"random small", []map[K]V{{{I: 42, S: "x"}: {N: "x", M: 42}}, {{I: 7, S: "y"}: {N: "y", M: 7}}}, []lxtypes.Pair[K, V]{{First: K{I: 42, S: "x"}, Second: V{N: "x", M: 42}}, {First: K{I: 7, S: "y"}, Second: V{N: "y", M: 7}}}},
		{"empty nil mix", []map[K]V{nil, {}, {{I: 8, S: "v"}: {N: "v", M: 8}}}, []lxtypes.Pair[K, V]{{First: K{I: 8, S: "v"}, Second: V{N: "v", M: 8}}}},
		{"final small", []map[K]V{{{I: 3, S: "end"}: {N: "end", M: 3}}}, []lxtypes.Pair[K, V]{{First: K{I: 3, S: "end"}, Second: V{N: "end", M: 3}}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Entries(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("Entries(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("Entries(%v) = nil; want non-nil with length %d", tt.input, len(tt.expected))
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("Entries(%v) length = %d; want %d", tt.input, len(got), len(tt.expected))
			}

			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Fatalf("Entries(%v) missing expected pairs; got %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
