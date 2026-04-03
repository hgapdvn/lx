package lxmaps_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hgapdvn/lx/maps"
)

// TestForEach_StringInt_BasicCases tests basic scenarios with string keys and int values
func TestForEach_StringInt_BasicCases(t *testing.T) {
	tests := []struct {
		name         string
		m            map[string]int
		expectedSize int
	}{
		{
			name:         "nil map",
			m:            nil,
			expectedSize: 0,
		},
		{
			name:         "empty map",
			m:            map[string]int{},
			expectedSize: 0,
		},
		{
			name:         "single element",
			m:            map[string]int{"a": 1},
			expectedSize: 1,
		},
		{
			name:         "multiple elements",
			m:            map[string]int{"a": 1, "b": 2, "c": 3},
			expectedSize: 3,
		},
		{
			name:         "five elements",
			m:            map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			expectedSize: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callCount := 0
			lxmaps.ForEach(tt.m, func(k string, v int) {
				callCount++
			})

			if callCount != tt.expectedSize {
				t.Errorf("ForEach() called %d times, want %d", callCount, tt.expectedSize)
			}
		})
	}
}

// TestForEach_StringInt_ValueVerification tests that correct values are passed to callback
func TestForEach_StringInt_ValueVerification(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	collected := make(map[string]int)
	lxmaps.ForEach(m, func(k string, v int) {
		collected[k] = v
	})

	if len(collected) != len(m) {
		t.Errorf("ForEach() collected %d items, want %d", len(collected), len(m))
	}

	for k, v := range m {
		if collectedVal, exists := collected[k]; !exists || collectedVal != v {
			t.Errorf("ForEach() did not correctly collect key=%q, value=%d", k, v)
		}
	}
}

// TestForEach_StringInt_ZeroValues tests handling of zero values
func TestForEach_StringInt_ZeroValues(t *testing.T) {
	m := map[string]int{
		"zero": 0,
		"pos":  1,
		"neg":  -1,
	}

	collected := make(map[string]int)
	lxmaps.ForEach(m, func(k string, v int) {
		collected[k] = v
	})

	if collected["zero"] != 0 {
		t.Errorf("ForEach() zero value not correctly handled")
	}
}

// TestForEach_StringInt_NegativeValues tests negative integer values
func TestForEach_StringInt_NegativeValues(t *testing.T) {
	m := map[string]int{
		"neg1": -100,
		"neg2": -5,
		"pos":  10,
	}

	minVal := 0
	hasNegative := false

	lxmaps.ForEach(m, func(k string, v int) {
		if v < 0 {
			hasNegative = true
		}
		if v < minVal {
			minVal = v
		}
	})

	if !hasNegative {
		t.Errorf("ForEach() did not find negative values")
	}
	if minVal >= 0 {
		t.Errorf("ForEach() min value should be negative, got %d", minVal)
	}
}

// TestForEach_StringInt_LargeValues tests large integer values
func TestForEach_StringInt_LargeValues(t *testing.T) {
	m := map[string]int{
		"max_int": 2147483647,
		"min_int": -2147483648,
		"zero":    0,
	}

	maxFound := false
	minFound := false

	lxmaps.ForEach(m, func(k string, v int) {
		if v == 2147483647 {
			maxFound = true
		}
		if v == -2147483648 {
			minFound = true
		}
	})

	if !maxFound || !minFound {
		t.Errorf("ForEach() did not find extreme values")
	}
}

// TestForEach_IntString tests with int keys and string values
func TestForEach_IntString(t *testing.T) {
	tests := []struct {
		name         string
		m            map[int]string
		expectedSize int
	}{
		{
			name:         "nil map",
			m:            nil,
			expectedSize: 0,
		},
		{
			name:         "empty map",
			m:            map[int]string{},
			expectedSize: 0,
		},
		{
			name:         "single element",
			m:            map[int]string{1: "one"},
			expectedSize: 1,
		},
		{
			name:         "multiple elements",
			m:            map[int]string{1: "one", 2: "two", 3: "three"},
			expectedSize: 3,
		},
		{
			name:         "negative keys",
			m:            map[int]string{-1: "neg_one", -2: "neg_two", 0: "zero"},
			expectedSize: 3,
		},
		{
			name:         "zero key",
			m:            map[int]string{0: "zero_value"},
			expectedSize: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callCount := 0
			lxmaps.ForEach(tt.m, func(k int, v string) {
				callCount++
			})

			if callCount != tt.expectedSize {
				t.Errorf("ForEach() called %d times, want %d", callCount, tt.expectedSize)
			}
		})
	}
}

// TestForEach_IntString_EmptyStringValues tests empty strings as values
func TestForEach_IntString_EmptyStringValues(t *testing.T) {
	m := map[int]string{
		1: "value",
		2: "",
		3: "another",
	}

	collected := make(map[int]string)
	lxmaps.ForEach(m, func(k int, v string) {
		collected[k] = v
	})

	if v, exists := collected[2]; !exists || v != "" {
		t.Errorf("ForEach() did not correctly handle empty string value")
	}
}

// TestForEach_StringString tests string keys and string values
func TestForEach_StringString(t *testing.T) {
	tests := []struct {
		name         string
		m            map[string]string
		expectedSize int
	}{
		{
			name:         "basic strings",
			m:            map[string]string{"a": "apple", "b": "banana"},
			expectedSize: 2,
		},
		{
			name:         "empty key with non-empty value",
			m:            map[string]string{"": "empty_key"},
			expectedSize: 1,
		},
		{
			name:         "non-empty key with empty value",
			m:            map[string]string{"key": ""},
			expectedSize: 1,
		},
		{
			name:         "both key and value empty",
			m:            map[string]string{"": ""},
			expectedSize: 1,
		},
		{
			name:         "unicode keys and values",
			m:            map[string]string{"hello": "世界", "foo": "бар"},
			expectedSize: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callCount := 0
			lxmaps.ForEach(tt.m, func(k string, v string) {
				callCount++
			})

			if callCount != tt.expectedSize {
				t.Errorf("ForEach() called %d times, want %d", callCount, tt.expectedSize)
			}
		})
	}
}

// TestForEach_StringBool tests string keys and boolean values
func TestForEach_StringBool(t *testing.T) {
	m := map[string]bool{
		"true_val":  true,
		"false_val": false,
		"another":   true,
	}

	trueCount := 0
	falseCount := 0

	lxmaps.ForEach(m, func(k string, v bool) {
		if v {
			trueCount++
		} else {
			falseCount++
		}
	})

	if trueCount != 2 || falseCount != 1 {
		t.Errorf("ForEach() counted true=%d, false=%d; want true=2, false=1", trueCount, falseCount)
	}
}

// TestForEach_StringFloat tests string keys and float values
func TestForEach_StringFloat(t *testing.T) {
	m := map[string]float64{
		"pi":       3.14159,
		"e":        2.71828,
		"sqrt_two": 1.41421,
		"negative": -2.5,
		"zero":     0.0,
	}

	sum := 0.0
	count := 0

	lxmaps.ForEach(m, func(k string, v float64) {
		sum += v
		count++
	})

	if count != 5 {
		t.Errorf("ForEach() iterated %d times, want 5", count)
	}

	// Check if sum is approximately correct (accounting for float precision)
	expectedSum := 3.14159 + 2.71828 + 1.41421 - 2.5
	if sum < expectedSum-0.01 || sum > expectedSum+0.01 {
		t.Errorf("ForEach() sum %.5f not close to expected %.5f", sum, expectedSum)
	}
}

// TestForEach_StringSlice tests string keys and slice values
func TestForEach_StringSlice(t *testing.T) {
	m := map[string][]int{
		"a": {1, 2, 3},
		"b": {4, 5},
		"c": {},
		"d": nil,
	}

	callCount := 0
	totalElements := 0

	lxmaps.ForEach(m, func(k string, v []int) {
		callCount++
		totalElements += len(v)
	})

	if callCount != 4 {
		t.Errorf("ForEach() called %d times, want 4", callCount)
	}

	if totalElements != 5 {
		t.Errorf("ForEach() total slice elements %d, want 5", totalElements)
	}
}

// TestForEach_CustomStruct tests custom struct values
func TestForEach_CustomStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	m := map[string]Person{
		"alice":   {Name: "Alice", Age: 30},
		"bob":     {Name: "Bob", Age: 25},
		"charlie": {Name: "Charlie", Age: 35},
	}

	totalAge := 0
	names := make([]string, 0)

	lxmaps.ForEach(m, func(k string, v Person) {
		totalAge += v.Age
		names = append(names, v.Name)
	})

	if totalAge != 90 {
		t.Errorf("ForEach() total age %d, want 90", totalAge)
	}

	if len(names) != 3 {
		t.Errorf("ForEach() collected %d names, want 3", len(names))
	}
}

// TestForEach_PointerValues tests pointer values
func TestForEach_StringPointerInt(t *testing.T) {
	val1 := 10
	val2 := 20
	val3 := 30

	m := map[string]*int{
		"a": &val1,
		"b": &val2,
		"c": &val3,
	}

	sum := 0
	nullCount := 0

	lxmaps.ForEach(m, func(k string, v *int) {
		if v != nil {
			sum += *v
		} else {
			nullCount++
		}
	})

	if sum != 60 {
		t.Errorf("ForEach() sum %d, want 60", sum)
	}

	if nullCount != 0 {
		t.Errorf("ForEach() found %d nil pointers, want 0", nullCount)
	}
}

// TestForEach_SpecialCharacterKeys tests keys with special characters
func TestForEach_SpecialCharacterKeys(t *testing.T) {
	m := map[string]string{
		"!@#$%":               "symbols",
		"key with spaces":     "spaced",
		"key\nwith\nnewlines": "newlined",
		"key\twith\ttabs":     "tabbed",
		"":                    "empty",
	}

	callCount := 0
	lxmaps.ForEach(m, func(k string, v string) {
		callCount++
	})

	if callCount != 5 {
		t.Errorf("ForEach() called %d times, want 5", callCount)
	}
}

// TestForEach_UnicodeCharacters tests unicode keys and values
func TestForEach_UnicodeCharacters(t *testing.T) {
	m := map[string]string{
		"日本語":        "Japanese",
		"العربية":    "Arabic",
		"русский":    "Russian",
		"हिन्दी":     "Hindi",
		"emoji😀test": "emoji_value",
	}

	callCount := 0
	lxmaps.ForEach(m, func(k string, v string) {
		callCount++
	})

	if callCount != 5 {
		t.Errorf("ForEach() called %d times, want 5", callCount)
	}
}

// TestForEach_LargeMap tests performance with large number of entries
func TestForEach_LargeMap(t *testing.T) {
	m := make(map[int]string)
	for i := 0; i < 10000; i++ {
		m[i] = fmt.Sprintf("value_%d", i)
	}

	callCount := 0
	lxmaps.ForEach(m, func(k int, v string) {
		callCount++
	})

	if callCount != 10000 {
		t.Errorf("ForEach() called %d times, want 10000", callCount)
	}
}

// TestForEach_CallbackModifiesExternalState tests side effects in callback
func TestForEach_CallbackModifiesExternalState(t *testing.T) {
	m := map[string]int{
		"a": 10,
		"b": 20,
		"c": 30,
	}

	sum := 0
	product := 1
	keys := make([]string, 0)

	lxmaps.ForEach(m, func(k string, v int) {
		sum += v
		product *= v
		keys = append(keys, k)
	})

	if sum != 60 {
		t.Errorf("ForEach() sum = %d, want 60", sum)
	}

	if product != 6000 {
		t.Errorf("ForEach() product = %d, want 6000", product)
	}

	if len(keys) != 3 {
		t.Errorf("ForEach() collected %d keys, want 3", len(keys))
	}
}

// TestForEach_StringBuilding tests string concatenation in callback
func TestForEach_StringBuilding(t *testing.T) {
	m := map[string]string{
		"hello": "world",
		"foo":   "bar",
		"test":  "case",
	}

	var result strings.Builder

	lxmaps.ForEach(m, func(k string, v string) {
		result.WriteString(k)
		result.WriteString("=")
		result.WriteString(v)
		result.WriteString(";")
	})

	resultStr := result.String()

	// Check all key-value pairs are in result
	if !strings.Contains(resultStr, "hello=world") {
		t.Errorf("ForEach() result missing hello=world")
	}
	if !strings.Contains(resultStr, "foo=bar") {
		t.Errorf("ForEach() result missing foo=bar")
	}
	if !strings.Contains(resultStr, "test=case") {
		t.Errorf("ForEach() result missing test=case")
	}
}

// TestForEach_ConditionalLogic tests conditional logic in callback
func TestForEach_ConditionalLogic(t *testing.T) {
	m := map[string]int{
		"a": 5,
		"b": 15,
		"c": 25,
		"d": 10,
		"e": 20,
	}

	above10 := 0
	below10 := 0
	equal10 := 0

	lxmaps.ForEach(m, func(k string, v int) {
		if v > 10 {
			above10++
		} else if v < 10 {
			below10++
		} else {
			equal10++
		}
	})

	if above10 != 3 {
		t.Errorf("ForEach() above 10: %d, want 3", above10)
	}
	if below10 != 1 {
		t.Errorf("ForEach() below 10: %d, want 1", below10)
	}
	if equal10 != 1 {
		t.Errorf("ForEach() equal 10: %d, want 1", equal10)
	}
}

// TestForEach_IntToString tests converting int keys to strings during iteration
func TestForEach_IntToString(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	stringKeys := make([]string, 0)
	lxmaps.ForEach(m, func(k int, v string) {
		stringKeys = append(stringKeys, strconv.Itoa(k))
	})

	if len(stringKeys) != 3 {
		t.Errorf("ForEach() collected %d string keys, want 3", len(stringKeys))
	}
}

// TestForEach_MapToMap tests collecting data from one map into another
func TestForEach_MapToMap(t *testing.T) {
	input := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	output := make(map[string]int)
	lxmaps.ForEach(input, func(k string, v int) {
		// Store doubled value
		output[k] = v * 2
	})

	if len(output) != 3 {
		t.Errorf("ForEach() output map size %d, want 3", len(output))
	}

	for k, v := range input {
		if output[k] != v*2 {
			t.Errorf("ForEach() output[%s] = %d, want %d", k, output[k], v*2)
		}
	}
}

// TestForEach_AllKeysVisited tests that all keys are visited
func TestForEach_AllKeysVisited(t *testing.T) {
	m := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
		"key4": 4,
		"key5": 5,
	}

	visitedKeys := make(map[string]bool)
	lxmaps.ForEach(m, func(k string, v int) {
		visitedKeys[k] = true
	})

	for k := range m {
		if !visitedKeys[k] {
			t.Errorf("ForEach() did not visit key %q", k)
		}
	}
}

// TestForEach_StringIntValueMatching tests that values match their keys
func TestForEach_StringIntValueMatching(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	keyValuePairs := make(map[string]int)
	lxmaps.ForEach(m, func(k string, v int) {
		keyValuePairs[k] = v
	})

	// Verify all key-value pairs match
	for originalKey, originalValue := range m {
		if collectedValue, exists := keyValuePairs[originalKey]; !exists || collectedValue != originalValue {
			t.Errorf("ForEach() key=%q: collected value %d, want %d", originalKey, collectedValue, originalValue)
		}
	}
}

// TestForEach_MapWithDuplicateValues tests maps with duplicate values
func TestForEach_MapWithDuplicateValues(t *testing.T) {
	m := map[string]int{
		"a": 10,
		"b": 10,
		"c": 10,
		"d": 20,
	}

	duplicateCount := 0
	lxmaps.ForEach(m, func(k string, v int) {
		if v == 10 {
			duplicateCount++
		}
	})

	if duplicateCount != 3 {
		t.Errorf("ForEach() found %d duplicate values, want 3", duplicateCount)
	}
}
