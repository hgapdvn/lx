package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/maps"
)

func TestGet_StringInt(t *testing.T) {
	tests := []struct {
		name           string
		input          map[string]int
		key            string
		expectedValue  int
		expectedExists bool
	}{
		{
			name:           "nil map",
			input:          nil,
			key:            "a",
			expectedValue:  0,
			expectedExists: false,
		},
		{
			name:           "empty map",
			input:          map[string]int{},
			key:            "a",
			expectedValue:  0,
			expectedExists: false,
		},
		{
			name:           "single entry key exists",
			input:          map[string]int{"a": 1},
			key:            "a",
			expectedValue:  1,
			expectedExists: true,
		},
		{
			name:           "single entry key does not exist",
			input:          map[string]int{"a": 1},
			key:            "b",
			expectedValue:  0,
			expectedExists: false,
		},
		{
			name:           "multiple entries key exists",
			input:          map[string]int{"a": 1, "b": 2, "c": 3},
			key:            "b",
			expectedValue:  2,
			expectedExists: true,
		},
		{
			name:           "multiple entries key does not exist",
			input:          map[string]int{"a": 1, "b": 2, "c": 3},
			key:            "d",
			expectedValue:  0,
			expectedExists: false,
		},
		{
			name:           "zero value exists",
			input:          map[string]int{"a": 0, "b": 1},
			key:            "a",
			expectedValue:  0,
			expectedExists: true,
		},
		{
			name:           "negative value exists",
			input:          map[string]int{"a": -1, "b": -2},
			key:            "a",
			expectedValue:  -1,
			expectedExists: true,
		},
		{
			name:           "empty string key",
			input:          map[string]int{"": 42, "a": 1},
			key:            "",
			expectedValue:  42,
			expectedExists: true,
		},
		{
			name:           "many entries get first",
			input:          map[string]int{"k1": 10, "k2": 20, "k3": 30, "k4": 40, "k5": 50},
			key:            "k1",
			expectedValue:  10,
			expectedExists: true,
		},
		{
			name:           "many entries get middle",
			input:          map[string]int{"k1": 10, "k2": 20, "k3": 30, "k4": 40, "k5": 50},
			key:            "k3",
			expectedValue:  30,
			expectedExists: true,
		},
		{
			name:           "many entries get last",
			input:          map[string]int{"k1": 10, "k2": 20, "k3": 30, "k4": 40, "k5": 50},
			key:            "k5",
			expectedValue:  50,
			expectedExists: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, exists := lxmaps.Get(tt.input, tt.key)
			if exists != tt.expectedExists {
				t.Errorf("Get() exists = %v, expected %v", exists, tt.expectedExists)
			}
			if value != tt.expectedValue {
				t.Errorf("Get() value = %v, expected %v", value, tt.expectedValue)
			}
		})
	}
}

func TestGet_IntString(t *testing.T) {
	tests := []struct {
		name           string
		input          map[int]string
		key            int
		expectedValue  string
		expectedExists bool
	}{
		{
			name:           "nil map",
			input:          nil,
			key:            1,
			expectedValue:  "",
			expectedExists: false,
		},
		{
			name:           "empty map",
			input:          map[int]string{},
			key:            1,
			expectedValue:  "",
			expectedExists: false,
		},
		{
			name:           "single entry exists",
			input:          map[int]string{1: "apple"},
			key:            1,
			expectedValue:  "apple",
			expectedExists: true,
		},
		{
			name:           "single entry does not exist",
			input:          map[int]string{1: "apple"},
			key:            2,
			expectedValue:  "",
			expectedExists: false,
		},
		{
			name:           "multiple entries exists",
			input:          map[int]string{1: "a", 2: "b", 3: "c"},
			key:            2,
			expectedValue:  "b",
			expectedExists: true,
		},
		{
			name:           "multiple entries does not exist",
			input:          map[int]string{1: "a", 2: "b", 3: "c"},
			key:            4,
			expectedValue:  "",
			expectedExists: false,
		},
		{
			name:           "empty string value",
			input:          map[int]string{1: "", 2: "b"},
			key:            1,
			expectedValue:  "",
			expectedExists: true,
		},
		{
			name:           "zero key",
			input:          map[int]string{0: "zero", 1: "one"},
			key:            0,
			expectedValue:  "zero",
			expectedExists: true,
		},
		{
			name:           "negative key",
			input:          map[int]string{-1: "minus", 1: "plus"},
			key:            -1,
			expectedValue:  "minus",
			expectedExists: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, exists := lxmaps.Get(tt.input, tt.key)
			if exists != tt.expectedExists {
				t.Errorf("Get() exists = %v, expected %v", exists, tt.expectedExists)
			}
			if value != tt.expectedValue {
				t.Errorf("Get() value = %v, expected %v", value, tt.expectedValue)
			}
		})
	}
}

func TestGet_StringBool(t *testing.T) {
	tests := []struct {
		name           string
		input          map[string]bool
		key            string
		expectedValue  bool
		expectedExists bool
	}{
		{
			name:           "nil map",
			input:          nil,
			key:            "flag",
			expectedValue:  false,
			expectedExists: false,
		},
		{
			name:           "empty map",
			input:          map[string]bool{},
			key:            "flag",
			expectedValue:  false,
			expectedExists: false,
		},
		{
			name:           "true value exists",
			input:          map[string]bool{"a": true},
			key:            "a",
			expectedValue:  true,
			expectedExists: true,
		},
		{
			name:           "false value exists",
			input:          map[string]bool{"a": false},
			key:            "a",
			expectedValue:  false,
			expectedExists: true,
		},
		{
			name:           "key does not exist",
			input:          map[string]bool{"a": true},
			key:            "b",
			expectedValue:  false,
			expectedExists: false,
		},
		{
			name:           "multiple entries true",
			input:          map[string]bool{"a": true, "b": false},
			key:            "a",
			expectedValue:  true,
			expectedExists: true,
		},
		{
			name:           "multiple entries false",
			input:          map[string]bool{"a": true, "b": false},
			key:            "b",
			expectedValue:  false,
			expectedExists: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, exists := lxmaps.Get(tt.input, tt.key)
			if exists != tt.expectedExists {
				t.Errorf("Get() exists = %v, expected %v", exists, tt.expectedExists)
			}
			if value != tt.expectedValue {
				t.Errorf("Get() value = %v, expected %v", value, tt.expectedValue)
			}
		})
	}
}

func TestGet_StringInterface(t *testing.T) {
	tests := []struct {
		name           string
		input          map[string]interface{}
		key            string
		expectedValue  interface{}
		expectedExists bool
	}{
		{
			name:           "nil map",
			input:          nil,
			key:            "a",
			expectedValue:  nil,
			expectedExists: false,
		},
		{
			name:           "empty map",
			input:          map[string]interface{}{},
			key:            "a",
			expectedValue:  nil,
			expectedExists: false,
		},
		{
			name:           "string value",
			input:          map[string]interface{}{"a": "hello"},
			key:            "a",
			expectedValue:  "hello",
			expectedExists: true,
		},
		{
			name:           "int value",
			input:          map[string]interface{}{"a": 42},
			key:            "a",
			expectedValue:  42,
			expectedExists: true,
		},
		{
			name:           "nil value exists",
			input:          map[string]interface{}{"a": nil},
			key:            "a",
			expectedValue:  nil,
			expectedExists: true,
		},
		{
			name:           "bool value",
			input:          map[string]interface{}{"a": true},
			key:            "a",
			expectedValue:  true,
			expectedExists: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, exists := lxmaps.Get(tt.input, tt.key)
			if exists != tt.expectedExists {
				t.Errorf("Get() exists = %v, expected %v", exists, tt.expectedExists)
			}
			if value != tt.expectedValue {
				t.Errorf("Get() value = %v, expected %v", value, tt.expectedValue)
			}
		})
	}
}

func BenchmarkGet(b *testing.B) {
	m := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
		"f": 6, "g": 7, "h": 8, "i": 9, "j": 10,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lxmaps.Get(m, "d")
	}
}
