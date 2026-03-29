package lxmaps_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/lxslices"
	"github.com/nthanhhai2909/lx/maps"
)

func TestKeys_String(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[string]int // support multiple maps
		expected []string
	}{
		{"no args", nil, nil},
		{"single nil map", []map[string]int{nil}, []string{}},
		{"empty map", []map[string]int{{}}, []string{}},
		{"single key", []map[string]int{{"a": 1}}, []string{"a"}},
		{"multiple keys small", []map[string]int{{"a": 1, "b": 2, "c": 3}}, []string{"a", "b", "c"}},
		{"multi maps duplicated key", []map[string]int{{"a": 1, "b": 2}, {"b": 3, "c": 4}}, []string{"a", "b", "b", "c"}},
		{"empty + map", []map[string]int{{}, {"x": 1}}, []string{"x"}},
		{"unicode keys", []map[string]int{{"こんにちは": 1}, {"世界": 2}}, []string{"こんにちは", "世界"}},
		{"emoji keys", []map[string]int{{"😊": 1}, {"emoji": 2}}, []string{"😊", "emoji"}},
		{"long strings", []map[string]int{{"long_string_example_0123456789": 1}, {"other_long_string_abcdefgh": 2}}, []string{"long_string_example_0123456789", "other_long_string_abcdefgh"}},
		{"separator-like keys", []map[string]int{{"a,b": 1}, {"c;d": 2}}, []string{"a,b", "c;d"}},
		{"numeric-like strings", []map[string]int{{"1": 1}, {"2": 2}}, []string{"1", "2"}},
		{"special chars", []map[string]int{{"!@#": 1}, {"$%": 2}}, []string{"!@#", "$%"}},
		{"space keys", []map[string]int{{" ": 1}, {"  ": 2}}, []string{" ", "  "}},
		{"newline key", []map[string]int{{"\n": 1}, {"ok": 2}}, []string{"\n", "ok"}},
		{"tab key", []map[string]int{{"\t": 1}, {"ok": 2}}, []string{"\t", "ok"}},
		{"case sensitive", []map[string]int{{"Test": 1}, {"test": 2}}, []string{"Test", "test"}},
		{"mixed ascii unicode", []map[string]int{{"a": 1}, {"ä": 2}, {"ö": 3}, {"ü": 4}}, []string{"a", "ä", "ö", "ü"}},
		{"prefix suffix", []map[string]int{{"pre_": 1}, {"_suf": 2}}, []string{"pre_", "_suf"}},
		{"zwj emoji", []map[string]int{{"👩‍💻": 1}, {"👨‍🚀": 2}}, []string{"👩‍💻", "👨‍🚀"}},
		{"many keys small", []map[string]int{{"k1": 1, "k2": 2, "k3": 3, "k4": 4, "k5": 5}}, []string{"k1", "k2", "k3", "k4", "k5"}},
		{"mixed small", []map[string]int{{"X": 1}, {"y": 2}, {"Z": 3}}, []string{"X", "y", "Z"}},
		{"numeric words", []map[string]int{{"one": 1}, {"two": 2}, {"three": 3}}, []string{"one", "two", "three"}},
		{"punctuated", []map[string]int{{"end.": 1}, {"start,": 2}}, []string{"end.", "start,"}},
		{"mixed lengths", []map[string]int{{"s": 1}, {"medium": 2}, {"a very long key indeed": 3}}, []string{"s", "medium", "a very long key indeed"}},
		{"repeated pattern keys", []map[string]int{{"p1": 1}, {"p2": 2}, {"p3": 3}, {"p4": 4}}, []string{"p1", "p2", "p3", "p4"}},
		{"underscore keys", []map[string]int{{"a_b": 1}, {"c_d": 2}}, []string{"a_b", "c_d"}},
		{"hyphen keys", []map[string]int{{"a-b": 1}, {"c-d": 2}}, []string{"a-b", "c-d"}},
		{"mixed symbols", []map[string]int{{"$a": 1}, {"#b": 2}, {"&c": 3}}, []string{"$a", "#b", "&c"}},
		{"final extra", []map[string]int{{"z1": 1}, {"z2": 2}}, []string{"z1", "z2"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Keys(tt.input...)

			// expected nil is only for true no-args; tables use single-nil map -> empty slice
			if tt.expected == nil {
				if got != nil {
					t.Fatalf("Keys(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("Keys(%v) = nil; want non-nil with length %d", tt.input, len(tt.expected))
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("Keys(%v) length = %d; want %d", tt.input, len(got), len(tt.expected))
			}

			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Fatalf("Keys(%v) missing expected keys; got %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestKeys_Int(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[int]string
		expected []int
	}{
		{"no args", nil, nil},
		{"single nil map", []map[int]string{nil}, []int{}},
		{"empty map", []map[int]string{{}}, []int{}},
		{"single key", []map[int]string{{1: "one"}}, []int{1}},
		{"multi maps duplicated key", []map[int]string{{1: "a", 2: "b"}, {2: "c", 3: "d"}}, []int{1, 2, 2, 3}},
		{"empty + map", []map[int]string{{}, {7: "y"}}, []int{7}},
		{"negative keys", []map[int]string{{-1: "neg"}, {-2: "neg2"}}, []int{-1, -2}},
		{"mixed sign", []map[int]string{{-5: "a"}, {5: "b"}}, []int{-5, 5}},
		{"large numbers", []map[int]string{{1000000: "m"}, {999999: "n"}}, []int{1000000, 999999}},
		{"sequential 5", []map[int]string{{10: "a", 11: "b", 12: "c", 13: "d", 14: "e"}}, []int{10, 11, 12, 13, 14}},
		{"sparse keys", []map[int]string{{2: "a"}, {100: "b"}}, []int{2, 100}},
		{"single large", []map[int]string{{999: "x"}}, []int{999}},
		{"many keys 6", []map[int]string{{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f"}}, []int{1, 2, 3, 4, 5, 6}},
		{"mixed small", []map[int]string{{7: "g"}, {8: "h"}, {9: "i"}}, []int{7, 8, 9}},
		{"even odd mix", []map[int]string{{2: "even"}, {3: "odd"}, {4: "even"}}, []int{2, 3, 4}},
		{"small negatives", []map[int]string{{-10: "a"}, {-11: "b"}}, []int{-10, -11}},
		{"zero only", []map[int]string{{0: "zero"}}, []int{0}},
		{"range 3", []map[int]string{{20: "a", 21: "b", 22: "c"}}, []int{20, 21, 22}},
		{"random small", []map[int]string{{42: "x"}, {7: "y"}}, []int{42, 7}},
		{"mixed many", []map[int]string{{100: "x"}, {-1: "y"}, {0: "z"}}, []int{100, -1, 0}},
		{"negative range", []map[int]string{{-20: "a"}, {-19: "b"}}, []int{-20, -19}},
		{"high low", []map[int]string{{214: "a"}, {-214: "b"}}, []int{214, -214}},
		{"step keys", []map[int]string{{1: "a"}, {3: "b"}, {5: "c"}}, []int{1, 3, 5}},
		{"modulus keys", []map[int]string{{2: "r2"}, {4: "r4"}, {6: "r6"}}, []int{2, 4, 6}},
		{"final extras", []map[int]string{{11: "k"}, {22: "l"}}, []int{11, 22}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Keys(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("Keys(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("Keys(%v) = nil; want non-nil with length %d", tt.input, len(tt.expected))
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("Keys(%v) length = %d; want %d", tt.input, len(got), len(tt.expected))
			}

			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Fatalf("Keys(%v) missing expected keys; got %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestKeys_Struct(t *testing.T) {
	// define a comparable struct type for keys
	type KS struct {
		I int
		S string
	}

	tests := []struct {
		name     string
		input    []map[KS]int
		expected []KS
	}{
		{"no args", nil, nil},
		{"single nil map", []map[KS]int{nil}, []KS{}},
		{"single key", []map[KS]int{{{I: 1, S: "one"}: 1}}, []KS{{I: 1, S: "one"}}},
		{"multi maps duplicated key", []map[KS]int{{{I: 1, S: "a"}: 1}, {{I: 1, S: "a"}: 3, {I: 2, S: "b"}: 2}}, []KS{{I: 1, S: "a"}, {I: 1, S: "a"}, {I: 2, S: "b"}}},
		{"many small maps", []map[KS]int{{{I: 1, S: "a"}: 1}, {{I: 2, S: "b"}: 2}, {{I: 3, S: "c"}: 3}}, []KS{{I: 1, S: "a"}, {I: 2, S: "b"}, {I: 3, S: "c"}}},
		{"unicode fields", []map[KS]int{{{I: 10, S: "こんにちは"}: 1}, {{I: 11, S: "世界"}: 2}}, []KS{{I: 10, S: "こんにちは"}, {I: 11, S: "世界"}}},
		{"emoji fields", []map[KS]int{{{I: 2, S: "😊"}: 1}, {{I: 3, S: "👋"}: 1}}, []KS{{I: 2, S: "😊"}, {I: 3, S: "👋"}}},
		{"empty then map", []map[KS]int{{}, {{I: 5, S: "X"}: 1}}, []KS{{I: 5, S: "X"}}},
		{"negatives and zero", []map[KS]int{{{I: 0, S: "z"}: 0}, {{I: -1, S: "n"}: -1}}, []KS{{I: 0, S: "z"}, {I: -1, S: "n"}}},
		{"prefix/suffix fields", []map[KS]int{{{I: 7, S: "pre_"}: 1}, {{I: 8, S: "_suf"}: 2}}, []KS{{I: 7, S: "pre_"}, {I: 8, S: "_suf"}}},
		{"long string field", []map[KS]int{{{I: 100, S: "long_string_example_012345"}: 1}}, []KS{{I: 100, S: "long_string_example_012345"}}},
		{"mixed unicode ascii", []map[KS]int{{{I: 9, S: "a"}: 1}, {{I: 10, S: "ä"}: 2}}, []KS{{I: 9, S: "a"}, {I: 10, S: "ä"}}},
		{"random small", []map[KS]int{{{I: 42, S: "x"}: 1}, {{I: 7, S: "y"}: 2}}, []KS{{I: 42, S: "x"}, {I: 7, S: "y"}}},
		{"two similar", []map[KS]int{{{I: 1, S: "same"}: 1}, {{I: 1, S: "same2"}: 2}}, []KS{{I: 1, S: "same"}, {I: 1, S: "same2"}}},
		{"mixed many", []map[KS]int{{{I: 0, S: "0"}: 1}, {{I: 1, S: "1"}: 2}, {{I: 2, S: "2"}: 3}}, []KS{{I: 0, S: "0"}, {I: 1, S: "1"}, {I: 2, S: "2"}}},
		{"final small", []map[KS]int{{{I: 3, S: "end"}: 1}, {{I: 4, S: "stop"}: 2}}, []KS{{I: 3, S: "end"}, {I: 4, S: "stop"}}},
		{"many duplicates", []map[KS]int{{{I: 1, S: "d"}: 1}, {{I: 1, S: "d"}: 2}, {{I: 1, S: "d"}: 3}}, []KS{{I: 1, S: "d"}, {I: 1, S: "d"}, {I: 1, S: "d"}}},
		{"mixed empty/nil", []map[KS]int{nil, {}, {{I: 8, S: "v"}: 8}}, []KS{{I: 8, S: "v"}}},
		{"extra unicode mix", []map[KS]int{{{I: 21, S: "α"}: 1}, {{I: 22, S: "β"}: 2}}, []KS{{I: 21, S: "α"}, {I: 22, S: "β"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Keys(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("Keys(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("Keys(%v) = nil; want non-nil with length %d", tt.input, len(tt.expected))
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("Keys(%v) length = %d; want %d", tt.input, len(got), len(tt.expected))
			}

			for _, exp := range tt.expected {
				if !lxslices.Contains(got, exp) {
					t.Fatalf("Keys(%v) missing key %v in result %v", tt.input, exp, got)
				}
			}
		})
	}
}
