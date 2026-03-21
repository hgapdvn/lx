package lxmaps_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/lxmaps"
	"github.com/nthanhhai2909/lx/lxslices"
	"github.com/nthanhhai2909/lx/lxtypes"
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
		{"empty map", []map[int]string{map[int]string{}}, []int{}},
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

func TestValues_String(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[string]int
		expected []int
	}{
		{"no args", nil, nil},
		{"single nil map", []map[string]int{nil}, []int{}},
		{"empty map", []map[string]int{{}}, []int{}},
		{"single value", []map[string]int{{"a": 1}}, []int{1}},
		{"multiple values small", []map[string]int{{"a": 1, "b": 2, "c": 3}}, []int{1, 2, 3}},
		{"unicode values", []map[string]int{{"k1": 10}, {"k2": 20}}, []int{10, 20}},
		{"emoji values", []map[string]int{{"e": 100}, {"f": 200}}, []int{100, 200}},
		{"long values", []map[string]int{{"long": 123456}, {"other": 654321}}, []int{123456, 654321}},
		{"many values small", []map[string]int{{"k1": 1, "k2": 2, "k3": 3, "k4": 4, "k5": 5}}, []int{1, 2, 3, 4, 5}},
		{"mixed values", []map[string]int{{"x": 7}, {"y": 8}, {"z": 9}}, []int{7, 8, 9}},
		{"zero value", []map[string]int{{"zero": 0}, {"one": 1}}, []int{0, 1}},
		{"negative values", []map[string]int{{"n1": -1}, {"n2": -2}}, []int{-1, -2}},
		{"sparse values", []map[string]int{{"a": 2}, {"b": 100}}, []int{2, 100}},
		{"single large", []map[string]int{{"big": 999}}, []int{999}},
		{"sequential", []map[string]int{{"a": 10}, {"b": 11}, {"c": 12}}, []int{10, 11, 12}},
		{"mixed many", []map[string]int{{"p": 100}, {"q": -1}, {"r": 0}}, []int{100, -1, 0}},
		{"duplicate key across maps", []map[string]int{{"dup": 1}, {"dup": 2}}, []int{1, 2}},
		{"empty + map", []map[string]int{{}, {"x": 3}}, []int{3}},
		{"many maps combined", []map[string]int{{"m1": 1}, {"m2": 2}, {"m3": 3}, {"m4": 4}}, []int{1, 2, 3, 4}},
		{"special chars", []map[string]int{{"!@#": 7}, {"$%": 8}}, []int{7, 8}},
		{"final small", []map[string]int{{"end": 3}, {"stop": 4}}, []int{3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Values(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("Values(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("Values(%v) = nil; want non-nil with length %d", tt.input, len(tt.expected))
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("Values(%v) length = %d; want %d", tt.input, len(got), len(tt.expected))
			}

			for _, exp := range tt.expected {
				if !lxslices.Contains(got, exp) {
					t.Fatalf("Values(%v) missing value %v in result %v", tt.input, exp, got)
				}
			}
		})
	}
}

func TestValues_Int(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[int]string
		expected []string
	}{
		{"no args", nil, nil},
		{"single nil map", []map[int]string{nil}, []string{}},
		{"empty map", []map[int]string{{}}, []string{}},
		{"single value", []map[int]string{{1: "one"}}, []string{"one"}},
		{"multiple values small", []map[int]string{{1: "a", 2: "b", 3: "c"}}, []string{"a", "b", "c"}},
		{"zero value", []map[int]string{{0: "zero"}, {1: "one"}}, []string{"zero", "one"}},
		{"negative values", []map[int]string{{-1: "neg"}, {-2: "neg2"}}, []string{"neg", "neg2"}},
		{"mixed values", []map[int]string{{-5: "a"}, {5: "b"}}, []string{"a", "b"}},
		{"long strings", []map[int]string{{100: "long"}, {200: "other"}}, []string{"long", "other"}},
		{"many values", []map[int]string{{1: "k1", 2: "k2", 3: "k3", 4: "k4", 5: "k5"}}, []string{"k1", "k2", "k3", "k4", "k5"}},
		{"mixed small", []map[int]string{{7: "g"}, {8: "h"}, {9: "i"}}, []string{"g", "h", "i"}},
		{"sparse", []map[int]string{{2: "a"}, {100: "b"}}, []string{"a", "b"}},
		{"single large", []map[int]string{{999: "x"}}, []string{"x"}},
		{"duplicate across maps", []map[int]string{{2: "b"}, {2: "c"}}, []string{"b", "c"}},
		{"empty + map", []map[int]string{{}, {42: "x"}}, []string{"x"}},
		{"final small", []map[int]string{{3: "end"}, {4: "stop"}}, []string{"end", "stop"}},
		{"random small", []map[int]string{{42: "x"}, {7: "y"}}, []string{"x", "y"}},
		{"mixed many", []map[int]string{{100: "x"}, {-1: "y"}, {0: "z"}}, []string{"x", "y", "z"}},
		{"sequential", []map[int]string{{10: "a"}, {11: "b"}, {12: "c"}}, []string{"a", "b", "c"}},
		{"final extras", []map[int]string{{11: "k"}, {22: "l"}}, []string{"k", "l"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Values(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("Values(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("Values(%v) = nil; want non-nil with length %d", tt.input, len(tt.expected))
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("Values(%v) length = %d; want %d", tt.input, len(got), len(tt.expected))
			}

			for _, exp := range tt.expected {
				if !lxslices.Contains(got, exp) {
					t.Fatalf("Values(%v) missing value %q in result %v", tt.input, exp, got)
				}
			}
		})
	}
}

func TestValues_Struct(t *testing.T) {
	// comparable struct type for values
	type VS struct {
		I int
		S string
	}

	tests := []struct {
		name     string
		input    []map[string]VS
		expected []VS
	}{
		{"no args", nil, nil},
		{"single nil map", []map[string]VS{nil}, []VS{}},
		{"empty map", []map[string]VS{{}}, []VS{}},
		{"single value", []map[string]VS{{"a": {I: 1, S: "one"}}}, []VS{{I: 1, S: "one"}}},
		{"multiple values small", []map[string]VS{{"a": {I: 1, S: "a"}, "b": {I: 2, S: "b"}}}, []VS{{I: 1, S: "a"}, {I: 2, S: "b"}}},
		{"duplicate across maps", []map[string]VS{{"dup": {I: 1, S: "a"}}, {"dup": {I: 2, S: "b"}}}, []VS{{I: 1, S: "a"}, {I: 2, S: "b"}}},
		{"many maps combined", []map[string]VS{{"m1": {I: 1, S: "m1"}}, {"m2": {I: 2, S: "m2"}}, {"m3": {I: 3, S: "m3"}}}, []VS{{I: 1, S: "m1"}, {I: 2, S: "m2"}, {I: 3, S: "m3"}}},
		{"unicode values", []map[string]VS{{"u": {I: 10, S: "こんにちは"}}, {"v": {I: 11, S: "世界"}}}, []VS{{I: 10, S: "こんにちは"}, {I: 11, S: "世界"}}},
		{"emoji values", []map[string]VS{{"e": {I: 2, S: "😊"}}, {"f": {I: 3, S: "👋"}}}, []VS{{I: 2, S: "😊"}, {I: 3, S: "👋"}}},
		{"empty+nil mix", []map[string]VS{nil, {}, {"x": {I: 5, S: "X"}}}, []VS{{I: 5, S: "X"}}},
		{"long strings", []map[string]VS{{"long": {I: 12345, S: "longstring..."}}}, []VS{{I: 12345, S: "longstring..."}}},
		{"random small", []map[string]VS{{"x": {I: 42, S: "x"}}, {"y": {I: 7, S: "y"}}}, []VS{{I: 42, S: "x"}, {I: 7, S: "y"}}},
		{"many duplicates", []map[string]VS{{"d": {I: 1, S: "a"}}, {"d": {I: 2, S: "b"}}, {"d": {I: 3, S: "c"}}}, []VS{{I: 1, S: "a"}, {I: 2, S: "b"}, {I: 3, S: "c"}}},
		{"special chars", []map[string]VS{{"!@#": {I: 7, S: "sym"}}, {"$%": {I: 8, S: "sym2"}}}, []VS{{I: 7, S: "sym"}, {I: 8, S: "sym2"}}},
		{"final small", []map[string]VS{{"end": {I: 3, S: "end"}}, {"stop": {I: 4, S: "stop"}}}, []VS{{I: 3, S: "end"}, {I: 4, S: "stop"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Values(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("Values(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("Values(%v) = nil; want non-nil with length %d", tt.input, len(tt.expected))
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("Values(%v) length = %d; want %d", tt.input, len(got), len(tt.expected))
			}

			for _, exp := range tt.expected {
				if !lxslices.Contains(got, exp) {
					t.Fatalf("Values(%v) missing value %v in result %v", tt.input, exp, got)
				}
			}
		})
	}
}

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

func TestUniqKeys_String(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[string]int
		expected []string // unique expected keys, order not assumed
	}{
		{"no args", nil, nil},
		{"single nil map", []map[string]int{nil}, []string{}},
		{"all nil variadic", []map[string]int{nil, nil}, []string{}},
		{"single map", []map[string]int{{"a": 1, "b": 2}}, []string{"a", "b"}},
		{"multi maps duplicated keys", []map[string]int{{"a": 1, "b": 2}, {"b": 3, "c": 4}}, []string{"a", "b", "c"}},
		{"many duplicates", []map[string]int{{"d": 1}, {"d": 2}, {"d": 3}}, []string{"d"}},
		{"mix empty and nil", []map[string]int{nil, {}, {"x": 9}}, []string{"x"}},
		{"unicode and emoji", []map[string]int{{"こんにちは": 1}, {"😊": 2, "こんにちは": 3}}, []string{"こんにちは", "😊"}},
		{"special chars", []map[string]int{{"!@#": 7}, {"$%": 8}}, []string{"!@#", "$%"}},
		{"many maps combined", []map[string]int{{"m1": 1}, {"m2": 2}, {"m3": 3}}, []string{"m1", "m2", "m3"}},
		{"long keys", []map[string]int{{"this_is_a_very_long_key_example_0001": 1}, {"this_is_a_very_long_key_example_0002": 2}}, []string{"this_is_a_very_long_key_example_0001", "this_is_a_very_long_key_example_0002"}},
		{"comma/semicolon mix", []map[string]int{{"a,b": 1}, {"c;d": 2}}, []string{"a,b", "c;d"}},
		{"numeric string keys", []map[string]int{{"1": 1}, {"2": 2}, {"10": 10}}, []string{"1", "2", "10"}},
		{"whitespace keys", []map[string]int{{" ": 1}, {"\t": 2}, {"\n": 3}}, []string{" ", "\t", "\n"}},
		{"case sensitivity", []map[string]int{{"Go": 1}, {"go": 2}}, []string{"Go", "go"}},
		{"prefix/suffix patterns", []map[string]int{{"pre_": 1}, {"_suf": 2}, {"pre_suf": 3}}, []string{"pre_", "_suf", "pre_suf"}},
		{"punctuation", []map[string]int{{"end.": 1}, {"start,": 2}, {"(paren)": 3}}, []string{"end.", "start,", "(paren)"}},
		{"small randomized mix", []map[string]int{{"a": 1, "b": 2}, {"c": 3}, {"b": 4, "d": 5}}, []string{"a", "b", "c", "d"}},
		{"final extras", []map[string]int{{"z1": 1}, {"z2": 2}, {"z3": 3}}, []string{"z1", "z2", "z3"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.UniqKeys(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("UniqKeys(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("UniqKeys(%v) = nil; want non-nil", tt.input)
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("UniqKeys(%v) length = %d; want %d; got=%v", tt.input, len(got), len(tt.expected), got)
			}

			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Fatalf("UniqKeys(%v) missing expected keys; got %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestUniqKeys_Int(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[int]string
		expected []int
	}{
		{"no args", nil, nil},
		{"single nil map", []map[int]string{nil}, []int{}},
		{"all nil variadic", []map[int]string{nil, nil}, []int{}},
		{"multi maps duplicated", []map[int]string{{1: "a", 2: "b"}, {2: "c", 3: "d"}}, []int{1, 2, 3}},
		{"many duplicates", []map[int]string{{5: "x"}, {5: "y"}, {5: "z"}}, []int{5}},
		{"mix empty and nil", []map[int]string{nil, {}, {9: "k"}}, []int{9}},
		{"large keys", []map[int]string{{1000: "a"}, {2000: "b"}, {3000: "c"}}, []int{1000, 2000, 3000}},
		{"negative and zero", []map[int]string{{0: "z"}, {-1: "n"}, {-2: "m"}}, []int{0, -1, -2}},
		{"sequential keys", []map[int]string{{1: "a", 2: "b", 3: "c"}}, []int{1, 2, 3}},
		{"sparse and dense mix", []map[int]string{{1: "a"}, {100: "b"}, {2: "c"}}, []int{1, 100, 2}},
		{"combined many", []map[int]string{{10: "x"}, {11: "y"}, {12: "z"}, {13: "t"}}, []int{10, 11, 12, 13}},
		{"duplicates across many", []map[int]string{{7: "a"}, {7: "b"}, {7: "c"}, {8: "d"}}, []int{7, 8}},
		{"final extras", []map[int]string{{11: "k"}, {22: "l"}}, []int{11, 22}},
		{"randomized small", []map[int]string{{42: "x", 7: "y"}, {7: "z"}}, []int{42, 7}},
		{"edge small", []map[int]string{{-2147483648: "min"}, {2147483647: "max"}}, []int{-2147483648, 2147483647}},
		{"mixed signed", []map[int]string{{-3: "a"}, {3: "b"}}, []int{-3, 3}},
		{"step keys", []map[int]string{{2: "a"}, {4: "b"}, {6: "c"}}, []int{2, 4, 6}},
		{"modulus keys", []map[int]string{{2: "r2"}, {4: "r4"}, {6: "r6"}}, []int{2, 4, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.UniqKeys(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("UniqKeys(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("UniqKeys(%v) = nil; want non-nil", tt.input)
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("UniqKeys(%v) length = %d; want %d; got=%v", tt.input, len(got), len(tt.expected), got)
			}

			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Fatalf("UniqKeys(%v) missing expected keys; got %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestUniqKeys_Struct(t *testing.T) {
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
		{"all nil variadic", []map[KS]int{nil, nil}, []KS{}},
		{"multi maps struct keys", []map[KS]int{{{I: 1, S: "a"}: 1, {I: 2, S: "b"}: 2}, {{I: 2, S: "b"}: 3, {I: 3, S: "c"}: 4}}, []KS{{I: 1, S: "a"}, {I: 2, S: "b"}, {I: 3, S: "c"}}},
		{"duplicates many", []map[KS]int{{{I: 9, S: "x"}: 1}, {{I: 9, S: "x"}: 2}}, []KS{{I: 9, S: "x"}}},
		{"empty+nil+map", []map[KS]int{nil, {}, {{I: 5, S: "X"}: 5}}, []KS{{I: 5, S: "X"}}},
		{"mixed structs", []map[KS]int{{{I: 0, S: "zero"}: 0}, {{I: -1, S: "neg"}: -1}, {{I: 42, S: "ok"}: 42}}, []KS{{I: 0, S: "zero"}, {I: -1, S: "neg"}, {I: 42, S: "ok"}}},
		{"unicode struct fields", []map[KS]int{{{I: 1, S: "こんにちは"}: 1}, {{I: 2, S: "世界"}: 2}}, []KS{{I: 1, S: "こんにちは"}, {I: 2, S: "世界"}}},
		{"many maps combined", []map[KS]int{{{I: 1, S: "a"}: 1}, {{I: 2, S: "b"}: 2}, {{I: 3, S: "c"}: 3}, {{I: 4, S: "d"}: 4}}, []KS{{I: 1, S: "a"}, {I: 2, S: "b"}, {I: 3, S: "c"}, {I: 4, S: "d"}}},
		{"final extra", []map[KS]int{{{I: 7, S: "f"}: 7}, {{I: 8, S: "g"}: 8}}, []KS{{I: 7, S: "f"}, {I: 8, S: "g"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.UniqKeys(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("UniqKeys(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("UniqKeys(%v) = nil; want non-nil", tt.input)
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("UniqKeys(%v) length = %d; want %d; got=%v", tt.input, len(got), len(tt.expected), got)
			}

			for _, exp := range tt.expected {
				if !lxslices.Contains(got, exp) {
					t.Fatalf("UniqKeys(%v) missing expected key %v; got=%v", tt.input, exp, got)
				}
			}
		})
	}
}
