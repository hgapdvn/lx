package lxmaps_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/maps"
)

func TestGetOrDefault_StringInt(t *testing.T) {
	tests := []struct {
		name          string
		input         map[string]int
		key           string
		defaultValue  int
		expectedValue int
	}{
		{
			name:          "nil map returns default",
			input:         nil,
			key:           "a",
			defaultValue:  999,
			expectedValue: 999,
		},
		{
			name:          "empty map returns default",
			input:         map[string]int{},
			key:           "a",
			defaultValue:  999,
			expectedValue: 999,
		},
		{
			name:          "single entry key exists",
			input:         map[string]int{"a": 1},
			key:           "a",
			defaultValue:  999,
			expectedValue: 1,
		},
		{
			name:          "single entry key does not exist",
			input:         map[string]int{"a": 1},
			key:           "b",
			defaultValue:  999,
			expectedValue: 999,
		},
		{
			name:          "multiple entries key exists",
			input:         map[string]int{"a": 1, "b": 2, "c": 3},
			key:           "b",
			defaultValue:  999,
			expectedValue: 2,
		},
		{
			name:          "multiple entries key does not exist",
			input:         map[string]int{"a": 1, "b": 2, "c": 3},
			key:           "d",
			defaultValue:  999,
			expectedValue: 999,
		},
		{
			name:          "zero value exists",
			input:         map[string]int{"a": 0, "b": 1},
			key:           "a",
			defaultValue:  999,
			expectedValue: 0,
		},
		{
			name:          "zero default",
			input:         map[string]int{"a": 1},
			key:           "b",
			defaultValue:  0,
			expectedValue: 0,
		},
		{
			name:          "negative value exists",
			input:         map[string]int{"a": -5},
			key:           "a",
			defaultValue:  999,
			expectedValue: -5,
		},
		{
			name:          "negative value missing",
			input:         map[string]int{"a": 1},
			key:           "b",
			defaultValue:  -999,
			expectedValue: -999,
		},
		{
			name:          "empty string key exists",
			input:         map[string]int{"": 42, "a": 1},
			key:           "",
			defaultValue:  999,
			expectedValue: 42,
		},
		{
			name:          "empty string key missing",
			input:         map[string]int{"a": 1},
			key:           "",
			defaultValue:  999,
			expectedValue: 999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := lxmaps.GetOrDefault(tt.input, tt.key, tt.defaultValue)
			if value != tt.expectedValue {
				t.Errorf("GetOrDefault() = %v, expected %v", value, tt.expectedValue)
			}
		})
	}
}

func TestGetOrDefault_IntString(t *testing.T) {
	tests := []struct {
		name          string
		input         map[int]string
		key           int
		defaultValue  string
		expectedValue string
	}{
		{
			name:          "nil map returns default",
			input:         nil,
			key:           1,
			defaultValue:  "default",
			expectedValue: "default",
		},
		{
			name:          "empty map returns default",
			input:         map[int]string{},
			key:           1,
			defaultValue:  "default",
			expectedValue: "default",
		},
		{
			name:          "single entry exists",
			input:         map[int]string{1: "apple"},
			key:           1,
			defaultValue:  "default",
			expectedValue: "apple",
		},
		{
			name:          "single entry does not exist",
			input:         map[int]string{1: "apple"},
			key:           2,
			defaultValue:  "default",
			expectedValue: "default",
		},
		{
			name:          "multiple entries exists",
			input:         map[int]string{1: "a", 2: "b", 3: "c"},
			key:           2,
			defaultValue:  "default",
			expectedValue: "b",
		},
		{
			name:          "empty string value exists",
			input:         map[int]string{1: "", 2: "b"},
			key:           1,
			defaultValue:  "default",
			expectedValue: "",
		},
		{
			name:          "zero key exists",
			input:         map[int]string{0: "zero", 1: "one"},
			key:           0,
			defaultValue:  "default",
			expectedValue: "zero",
		},
		{
			name:          "negative key exists",
			input:         map[int]string{-1: "minus", 1: "plus"},
			key:           -1,
			defaultValue:  "default",
			expectedValue: "minus",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := lxmaps.GetOrDefault(tt.input, tt.key, tt.defaultValue)
			if value != tt.expectedValue {
				t.Errorf("GetOrDefault() = %v, expected %v", value, tt.expectedValue)
			}
		})
	}
}

func TestGetOrDefault_StringBool(t *testing.T) {
	tests := []struct {
		name          string
		input         map[string]bool
		key           string
		defaultValue  bool
		expectedValue bool
	}{
		{
			name:          "nil map returns default true",
			input:         nil,
			key:           "a",
			defaultValue:  true,
			expectedValue: true,
		},
		{
			name:          "nil map returns default false",
			input:         nil,
			key:           "a",
			defaultValue:  false,
			expectedValue: false,
		},
		{
			name:          "true value exists",
			input:         map[string]bool{"a": true},
			key:           "a",
			defaultValue:  false,
			expectedValue: true,
		},
		{
			name:          "false value exists returns false",
			input:         map[string]bool{"a": false},
			key:           "a",
			defaultValue:  true,
			expectedValue: false,
		},
		{
			name:          "key does not exist returns default true",
			input:         map[string]bool{"a": true},
			key:           "b",
			defaultValue:  true,
			expectedValue: true,
		},
		{
			name:          "key does not exist returns default false",
			input:         map[string]bool{"a": true},
			key:           "b",
			defaultValue:  false,
			expectedValue: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := lxmaps.GetOrDefault(tt.input, tt.key, tt.defaultValue)
			if value != tt.expectedValue {
				t.Errorf("GetOrDefault() = %v, expected %v", value, tt.expectedValue)
			}
		})
	}
}

func TestGetOrDefault_StringInterface(t *testing.T) {
	tests := []struct {
		name          string
		input         map[string]interface{}
		key           string
		defaultValue  interface{}
		expectedValue interface{}
	}{
		{
			name:          "nil map returns default",
			input:         nil,
			key:           "a",
			defaultValue:  "default",
			expectedValue: "default",
		},
		{
			name:          "string value exists",
			input:         map[string]interface{}{"a": "hello"},
			key:           "a",
			defaultValue:  "default",
			expectedValue: "hello",
		},
		{
			name:          "int value exists",
			input:         map[string]interface{}{"a": 42},
			key:           "a",
			defaultValue:  0,
			expectedValue: 42,
		},
		{
			name:          "nil value exists",
			input:         map[string]interface{}{"a": nil},
			key:           "a",
			defaultValue:  "default",
			expectedValue: nil,
		},
		{
			name:          "key missing with default",
			input:         map[string]interface{}{"a": "value"},
			key:           "b",
			defaultValue:  "default",
			expectedValue: "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := lxmaps.GetOrDefault(tt.input, tt.key, tt.defaultValue)
			if value != tt.expectedValue {
				t.Errorf("GetOrDefault() = %v, expected %v", value, tt.expectedValue)
			}
		})
	}
}

func BenchmarkGetOrDefault(b *testing.B) {
	m := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
		"f": 6, "g": 7, "h": 8, "i": 9, "j": 10,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lxmaps.GetOrDefault(m, "d", 999)
	}
}
