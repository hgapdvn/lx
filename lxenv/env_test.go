package lxenv_test

import (
	"errors"
	"math"
	"os"
	"strings"
	"testing"

	"github.com/nthanhhai2909/lx/lxenv"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		setValue string
		setVar   bool
		expected string
	}{
		{
			name:     "existing variable",
			key:      "TEST_GET_EXISTING",
			setValue: "hello",
			setVar:   true,
			expected: "hello",
		},
		{
			name:     "non-existent variable",
			key:      "TEST_GET_NONEXISTENT",
			setVar:   false,
			expected: "",
		},
		{
			name:     "empty variable",
			key:      "TEST_GET_EMPTY",
			setValue: "",
			setVar:   true,
			expected: "",
		},
		{
			name:     "value with special characters",
			key:      "TEST_GET_SPECIAL",
			setValue: "hello@world!#$%",
			setVar:   true,
			expected: "hello@world!#$%",
		},
		{
			name:     "value with spaces",
			key:      "TEST_GET_SPACES",
			setValue: "hello world",
			setVar:   true,
			expected: "hello world",
		},
		{
			name:     "whitespace-only value",
			key:      "TEST_GET_WHITESPACE",
			setValue: "   ",
			setVar:   true,
			expected: "   ",
		},
		{
			name:     "numeric string value",
			key:      "TEST_GET_NUMERIC",
			setValue: "12345",
			setVar:   true,
			expected: "12345",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				t.Setenv(tt.key, tt.setValue)
			}

			result := lxenv.Get(tt.key)
			if result != tt.expected {
				t.Errorf("Get(%q) = %q, want %q", tt.key, result, tt.expected)
			}
		})
	}
}

func TestGetOr(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		setValue     string
		setVar       bool
		defaultValue string
		expected     string
	}{
		{
			name:         "existing variable returns value",
			key:          "TEST_GETOR_EXISTING",
			setValue:     "hello",
			setVar:       true,
			defaultValue: "default",
			expected:     "hello",
		},
		{
			name:         "non-existent variable returns default",
			key:          "TEST_GETOR_NONEXISTENT",
			setVar:       false,
			defaultValue: "default",
			expected:     "default",
		},
		{
			name:         "empty variable returns default",
			key:          "TEST_GETOR_EMPTY",
			setValue:     "",
			setVar:       true,
			defaultValue: "default",
			expected:     "default",
		},
		{
			name:         "value with special characters",
			key:          "TEST_GETOR_SPECIAL",
			setValue:     "hello@world!#$%",
			setVar:       true,
			defaultValue: "default",
			expected:     "hello@world!#$%",
		},
		{
			name:         "value with spaces",
			key:          "TEST_GETOR_SPACES",
			setValue:     "hello world",
			setVar:       true,
			defaultValue: "default",
			expected:     "hello world",
		},
		{
			name:         "whitespace-only value returns whitespace",
			key:          "TEST_GETOR_WHITESPACE",
			setValue:     "   ",
			setVar:       true,
			defaultValue: "default",
			expected:     "   ",
		},
		{
			name:         "numeric string value",
			key:          "TEST_GETOR_NUMERIC",
			setValue:     "12345",
			setVar:       true,
			defaultValue: "0",
			expected:     "12345",
		},
		{
			name:         "empty default value",
			key:          "TEST_GETOR_EMPTY_DEFAULT",
			setVar:       false,
			defaultValue: "",
			expected:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				t.Setenv(tt.key, tt.setValue)
			}

			result := lxenv.GetOr(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("GetOr(%q, %q) = %q, want %q", tt.key, tt.defaultValue, result, tt.expected)
			}
		})
	}
}

func TestMustGet(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		value     string
		setEnv    bool
		want      string
		wantPanic bool
	}{
		{
			name:   "env exists",
			key:    "TEST_ENV_EXIST",
			value:  "hello",
			setEnv: true,
			want:   "hello",
		},
		{
			name:      "env not exists",
			key:       "TEST_ENV_NOT_EXIST",
			setEnv:    false,
			wantPanic: true,
		},
		{
			name:   "empty value does not panic",
			key:    "TEST_ENV_EMPTY",
			value:  "",
			setEnv: true,
			want:   "",
		},
		{
			name:   "whitespace value preserved",
			key:    "TEST_ENV_WHITESPACE",
			value:  "   ",
			setEnv: true,
			want:   "   ",
		},
		{
			name:   "special characters preserved",
			key:    "TEST_ENV_SPECIAL",
			value:  "hello@world!#$%",
			setEnv: true,
			want:   "hello@world!#$%",
		},
		{
			name:   "unicode preserved",
			key:    "TEST_ENV_UNICODE",
			value:  "こんにちは🌏",
			setEnv: true,
			want:   "こんにちは🌏",
		},
		{
			name:   "newline preserved",
			key:    "TEST_ENV_NEWLINE",
			value:  "line1\nline2",
			setEnv: true,
			want:   "line1\nline2",
		},
		{
			name:   "tab preserved",
			key:    "TEST_ENV_TAB",
			value:  "\tindented",
			setEnv: true,
			want:   "\tindented",
		},
		{
			name:   "equals sign preserved",
			key:    "TEST_ENV_EQUALS",
			value:  "a=b=c",
			setEnv: true,
			want:   "a=b=c",
		},
		{
			name:   "zero string preserved",
			key:    "TEST_ENV_ZERO",
			value:  "0",
			setEnv: true,
			want:   "0",
		},
		{
			name:   "long value preserved",
			key:    "TEST_ENV_LONG",
			value:  strings.Repeat("x", 2048),
			setEnv: true,
			want:   strings.Repeat("x", 2048),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setEnv {
				t.Setenv(tt.key, tt.value)
			} else {
				os.Unsetenv(tt.key)
			}

			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("MustGet(%q) did not panic", tt.key)
					}
				}()
			}

			got := lxenv.MustGet(tt.key)
			if !tt.wantPanic && got != tt.want {
				t.Errorf("MustGet(%q) = %q, want %q", tt.key, got, tt.want)
			}
		})
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    string
		expected string
	}{
		{
			name:     "set simple value",
			key:      "TEST_SET_SIMPLE",
			value:    "hello",
			expected: "hello",
		},
		{
			name:     "set empty value",
			key:      "TEST_SET_EMPTY",
			value:    "",
			expected: "",
		},
		{
			name:     "set value with special characters",
			key:      "TEST_SET_SPECIAL",
			value:    "hello@world!#$%",
			expected: "hello@world!#$%",
		},
		{
			name:     "set value with spaces",
			key:      "TEST_SET_SPACES",
			value:    "hello world",
			expected: "hello world",
		},
		{
			name:     "set whitespace-only value",
			key:      "TEST_SET_WHITESPACE",
			value:    "   ",
			expected: "   ",
		},
		{
			name:     "set numeric string value",
			key:      "TEST_SET_NUMERIC",
			value:    "12345",
			expected: "12345",
		},
		{
			name:     "overwrite existing value",
			key:      "TEST_SET_OVERWRITE",
			value:    "new_value",
			expected: "new_value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "overwrite existing value" {
				os.Setenv(tt.key, "old_value")
			}
			defer os.Unsetenv(tt.key)

			err := lxenv.Set(tt.key, tt.value)
			if err != nil {
				t.Errorf("Set(%q, %q) returned unexpected error: %v", tt.key, tt.value, err)
			}

			result := os.Getenv(tt.key)
			if result != tt.expected {
				t.Errorf("after Set(%q, %q), Get = %q, want %q", tt.key, tt.value, result, tt.expected)
			}
		})
	}
}

func TestUnset(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		preset string
		setVar bool
	}{
		{
			name:   "unset existing variable",
			key:    "TEST_UNSET_EXISTING",
			preset: "hello",
			setVar: true,
		},
		{
			name:   "unset variable with empty value",
			key:    "TEST_UNSET_EMPTY",
			preset: "",
			setVar: true,
		},
		{
			name:   "unset variable with special characters value",
			key:    "TEST_UNSET_SPECIAL",
			preset: "hello@world!#$%",
			setVar: true,
		},
		{
			name:   "unset variable with spaces value",
			key:    "TEST_UNSET_SPACES",
			preset: "hello world",
			setVar: true,
		},
		{
			name:   "unset variable with numeric value",
			key:    "TEST_UNSET_NUMERIC",
			preset: "12345",
			setVar: true,
		},
		{
			name:   "unset non-existent variable",
			key:    "TEST_UNSET_NONEXISTENT",
			setVar: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				os.Setenv(tt.key, tt.preset)
			}

			err := lxenv.Unset(tt.key)
			if err != nil {
				t.Errorf("Unset(%q) returned unexpected error: %v", tt.key, err)
			}

			_, exists := os.LookupEnv(tt.key)
			if exists {
				t.Errorf("after Unset(%q), variable still exists", tt.key)
			}
		})
	}
}

func TestHas(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		preset   string
		setVar   bool
		expected bool
	}{
		{
			name:     "existing variable with value",
			key:      "TEST_HAS_EXISTING",
			preset:   "hello",
			setVar:   true,
			expected: true,
		},
		{
			name:     "existing variable with empty value",
			key:      "TEST_HAS_EMPTY_VALUE",
			preset:   "",
			setVar:   true,
			expected: true,
		},
		{
			name:     "existing variable with whitespace value",
			key:      "TEST_HAS_WHITESPACE",
			preset:   "   ",
			setVar:   true,
			expected: true,
		},
		{
			name:     "existing variable with special characters",
			key:      "TEST_HAS_SPECIAL",
			preset:   "hello@world!#$%",
			setVar:   true,
			expected: true,
		},
		{
			name:     "existing variable with numeric value",
			key:      "TEST_HAS_NUMERIC",
			preset:   "12345",
			setVar:   true,
			expected: true,
		},
		{
			name:     "non-existent variable",
			key:      "TEST_HAS_NONEXISTENT",
			setVar:   false,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				t.Setenv(tt.key, tt.preset)
			}

			result := lxenv.Has(tt.key)
			if result != tt.expected {
				t.Errorf("Has(%q) = %v, want %v", tt.key, result, tt.expected)
			}
		})
	}
}

func TestNotHas(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		preset   string
		setVar   bool
		expected bool
	}{
		{
			name:     "existing variable returns false",
			key:      "TEST_NOTHAS_EXISTING",
			preset:   "value",
			setVar:   true,
			expected: false,
		},
		{
			name:     "existing empty variable returns false",
			key:      "TEST_NOTHAS_EMPTY",
			preset:   "",
			setVar:   true,
			expected: false,
		},
		{
			name:     "existing whitespace variable returns false",
			key:      "TEST_NOTHAS_SPACE",
			preset:   "   ",
			setVar:   true,
			expected: false,
		},
		{
			name:     "existing special characters variable returns false",
			key:      "TEST_NOTHAS_SPECIAL",
			preset:   "hello@world!#$%",
			setVar:   true,
			expected: false,
		},
		{
			name:     "missing variable returns true",
			key:      "TEST_NOTHAS_MISSING",
			setVar:   false,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				t.Setenv(tt.key, tt.preset)
			} else {
				os.Unsetenv(tt.key)
			}

			result := lxenv.NotHas(tt.key)
			if result != tt.expected {
				t.Errorf("NotHas(%q) = %v, want %v", tt.key, result, tt.expected)
			}
		})
	}
}

func TestExists(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		preset   string
		setVar   bool
		expected bool
	}{
		{
			name:     "existing variable with value",
			key:      "TEST_HAS_EXISTING",
			preset:   "hello",
			setVar:   true,
			expected: true,
		},
		{
			name:     "existing variable with empty value",
			key:      "TEST_HAS_EMPTY_VALUE",
			preset:   "",
			setVar:   true,
			expected: true,
		},
		{
			name:     "existing variable with whitespace value",
			key:      "TEST_HAS_WHITESPACE",
			preset:   "   ",
			setVar:   true,
			expected: true,
		},
		{
			name:     "existing variable with special characters",
			key:      "TEST_HAS_SPECIAL",
			preset:   "hello@world!#$%",
			setVar:   true,
			expected: true,
		},
		{
			name:     "existing variable with numeric value",
			key:      "TEST_HAS_NUMERIC",
			preset:   "12345",
			setVar:   true,
			expected: true,
		},
		{
			name:     "non-existent variable",
			key:      "TEST_HAS_NONEXISTENT",
			setVar:   false,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				t.Setenv(tt.key, tt.preset)
			}

			result := lxenv.Exists(tt.key)
			if result != tt.expected {
				t.Errorf("Has(%q) = %v, want %v", tt.key, result, tt.expected)
			}
		})
	}
}

func TestNotExists(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		preset   string
		setVar   bool
		expected bool
	}{
		{
			name:     "existing variable returns false",
			key:      "TEST_NOTEXISTS_EXISTING",
			preset:   "value",
			setVar:   true,
			expected: false,
		},
		{
			name:     "existing empty variable returns false",
			key:      "TEST_NOTEXISTS_EMPTY",
			preset:   "",
			setVar:   true,
			expected: false,
		},
		{
			name:     "existing whitespace variable returns false",
			key:      "TEST_NOTEXISTS_SPACE",
			preset:   "   ",
			setVar:   true,
			expected: false,
		},
		{
			name:     "existing special characters variable returns false",
			key:      "TEST_NOTEXISTS_SPECIAL",
			preset:   "hello@world!#$%",
			setVar:   true,
			expected: false,
		},
		{
			name:     "missing variable returns true",
			key:      "TEST_NOTEXISTS_MISSING",
			setVar:   false,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				t.Setenv(tt.key, tt.preset)
			} else {
				os.Unsetenv(tt.key)
			}

			result := lxenv.NotExists(tt.key)
			if result != tt.expected {
				t.Errorf("NotExists(%q) = %v, want %v", tt.key, result, tt.expected)
			}
		})
	}
}

func TestLookup(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		preset        string
		setVar        bool
		expectedValue string
		expectedOk    bool
	}{
		{
			name:          "existing variable with value",
			key:           "TEST_LOOKUP_EXISTING",
			preset:        "hello",
			setVar:        true,
			expectedValue: "hello",
			expectedOk:    true,
		},
		{
			name:          "existing variable with empty value",
			key:           "TEST_LOOKUP_EMPTY_VALUE",
			preset:        "",
			setVar:        true,
			expectedValue: "",
			expectedOk:    true,
		},
		{
			name:          "existing variable with whitespace value",
			key:           "TEST_LOOKUP_WHITESPACE",
			preset:        "   ",
			setVar:        true,
			expectedValue: "   ",
			expectedOk:    true,
		},
		{
			name:          "existing variable with special characters",
			key:           "TEST_LOOKUP_SPECIAL",
			preset:        "hello@world!#$%",
			setVar:        true,
			expectedValue: "hello@world!#$%",
			expectedOk:    true,
		},
		{
			name:          "existing variable with numeric value",
			key:           "TEST_LOOKUP_NUMERIC",
			preset:        "12345",
			setVar:        true,
			expectedValue: "12345",
			expectedOk:    true,
		},
		{
			name:          "non-existent variable",
			key:           "TEST_LOOKUP_NONEXISTENT",
			setVar:        false,
			expectedValue: "",
			expectedOk:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				t.Setenv(tt.key, tt.preset)
			}

			value, ok := lxenv.Lookup(tt.key)
			if value != tt.expectedValue || ok != tt.expectedOk {
				t.Errorf("Lookup(%q) = (%q, %v), want (%q, %v)", tt.key, value, ok, tt.expectedValue, tt.expectedOk)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		preset        string
		setVar        bool
		expectedValue int
		expectedOk    bool
	}{
		{
			name:          "valid integer",
			key:           "TEST_GETINT_VALID",
			preset:        "42",
			setVar:        true,
			expectedValue: 42,
			expectedOk:    true,
		},
		{
			name:          "negative integer",
			key:           "TEST_GETINT_NEGATIVE",
			preset:        "-10",
			setVar:        true,
			expectedValue: -10,
			expectedOk:    true,
		},
		{
			name:          "zero",
			key:           "TEST_GETINT_ZERO",
			preset:        "0",
			setVar:        true,
			expectedValue: 0,
			expectedOk:    true,
		},
		{
			name:          "float value returns false",
			key:           "TEST_GETINT_FLOAT",
			preset:        "3.14",
			setVar:        true,
			expectedValue: 0,
			expectedOk:    false,
		},
		{
			name:          "string value returns false",
			key:           "TEST_GETINT_STRING",
			preset:        "not_a_number",
			setVar:        true,
			expectedValue: 0,
			expectedOk:    false,
		},
		{
			name:          "empty value returns false",
			key:           "TEST_GETINT_EMPTY",
			preset:        "",
			setVar:        true,
			expectedValue: 0,
			expectedOk:    false,
		},
		{
			name:          "non-existent variable returns false",
			key:           "TEST_GETINT_NONEXISTENT",
			setVar:        false,
			expectedValue: 0,
			expectedOk:    false,
		},
		{
			name:          "whitespace value returns false",
			key:           "TEST_GETINT_WHITESPACE",
			preset:        "   ",
			setVar:        true,
			expectedValue: 0,
			expectedOk:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				t.Setenv(tt.key, tt.preset)
			}

			value, ok := lxenv.GetInt(tt.key)
			if value != tt.expectedValue || ok != tt.expectedOk {
				t.Errorf("GetInt(%q) = (%d, %v), want (%d, %v)", tt.key, value, ok, tt.expectedValue, tt.expectedOk)
			}
		})
	}
}

func TestGetIntOr(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		preset       string
		setVar       bool
		defaultValue int
		expected     int
	}{
		{
			name:         "valid integer returns parsed value",
			key:          "TEST_GETINTOR_VALID",
			preset:       "42",
			setVar:       true,
			defaultValue: 100,
			expected:     42,
		},
		{
			name:         "negative integer returns parsed value",
			key:          "TEST_GETINTOR_NEGATIVE",
			preset:       "-10",
			setVar:       true,
			defaultValue: 100,
			expected:     -10,
		},
		{
			name:         "zero returns zero",
			key:          "TEST_GETINTOR_ZERO",
			preset:       "0",
			setVar:       true,
			defaultValue: 100,
			expected:     0,
		},
		{
			name:         "float value returns default",
			key:          "TEST_GETINTOR_FLOAT",
			preset:       "3.14",
			setVar:       true,
			defaultValue: 100,
			expected:     100,
		},
		{
			name:         "string value returns default",
			key:          "TEST_GETINTOR_STRING",
			preset:       "not_a_number",
			setVar:       true,
			defaultValue: 100,
			expected:     100,
		},
		{
			name:         "empty value returns default",
			key:          "TEST_GETINTOR_EMPTY",
			preset:       "",
			setVar:       true,
			defaultValue: 100,
			expected:     100,
		},
		{
			name:         "non-existent variable returns default",
			key:          "TEST_GETINTOR_NONEXISTENT",
			setVar:       false,
			defaultValue: 100,
			expected:     100,
		},
		{
			name:         "zero default value",
			key:          "TEST_GETINTOR_ZERO_DEFAULT",
			setVar:       false,
			defaultValue: 0,
			expected:     0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				t.Setenv(tt.key, tt.preset)
			}

			result := lxenv.GetIntOr(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("GetIntOr(%q, %d) = %d, want %d", tt.key, tt.defaultValue, result, tt.expected)
			}
		})
	}
}

func TestMustGetInt(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		value     string
		setEnv    bool
		want      int
		wantPanic bool
	}{
		{
			name:   "valid integer",
			key:    "TEST_MUSTGETINT_VALID",
			value:  "8080",
			setEnv: true,
			want:   8080,
		},
		{
			name:      "missing key panics",
			key:       "TEST_MUSTGETINT_MISSING",
			setEnv:    false,
			wantPanic: true,
		},
		{
			name:      "invalid integer panics",
			key:       "TEST_MUSTGETINT_INVALID",
			value:     "not_an_int",
			setEnv:    true,
			wantPanic: true,
		},
		{
			name:      "empty value treated as missing (panics)",
			key:       "TEST_MUSTGETINT_EMPTY",
			value:     "",
			setEnv:    true,
			wantPanic: true,
		},
		{
			name:      "whitespace value treated as invalid (panics)",
			key:       "TEST_MUSTGETINT_SPACE",
			value:     "   ",
			setEnv:    true,
			wantPanic: true,
		},
		{
			name:   "negative integer",
			key:    "TEST_MUSTGETINT_NEGATIVE",
			value:  "-42",
			setEnv: true,
			want:   -42,
		},
		{
			name:   "plus-sign integer",
			key:    "TEST_MUSTGETINT_PLUS",
			value:  "+123",
			setEnv: true,
			want:   123,
		},
		{
			name:      "leading/trailing spaces (panics)",
			key:       "TEST_MUSTGETINT_SPACED",
			value:     " 123 ",
			setEnv:    true,
			wantPanic: true,
		},
		{
			name:      "overflow value panics",
			key:       "TEST_MUSTGETINT_OVERFLOW",
			value:     "999999999999999999999999999999",
			setEnv:    true,
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setEnv {
				t.Setenv(tt.key, tt.value)
			} else {
				os.Unsetenv(tt.key)
			}

			var (
				got      int
				didPanic bool
			)
			func() {
				defer func() {
					if r := recover(); r != nil {
						didPanic = true
					}
				}()
				got = lxenv.MustGetInt(tt.key)
			}()

			if tt.wantPanic {
				if !didPanic {
					t.Fatalf("MustGetInt(%q) did not panic", tt.key)
				}
				return
			}

			if didPanic {
				t.Fatalf("MustGetInt(%q) panicked unexpectedly", tt.key)
			}

			if got != tt.want {
				t.Fatalf("MustGetInt(%q) = %d, want %d", tt.key, got, tt.want)
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		preset        string
		setVar        bool
		expectedValue bool
		expectedOk    bool
	}{
		// true values
		{
			name:          `"1" returns true`,
			key:           "TEST_GETBOOL_1",
			preset:        "1",
			setVar:        true,
			expectedValue: true,
			expectedOk:    true,
		},
		{
			name:          `"t" returns true`,
			key:           "TEST_GETBOOL_t",
			preset:        "t",
			setVar:        true,
			expectedValue: true,
			expectedOk:    true,
		},
		{
			name:          `"T" returns true`,
			key:           "TEST_GETBOOL_T",
			preset:        "T",
			setVar:        true,
			expectedValue: true,
			expectedOk:    true,
		},
		{
			name:          `"true" returns true`,
			key:           "TEST_GETBOOL_true",
			preset:        "true",
			setVar:        true,
			expectedValue: true,
			expectedOk:    true,
		},
		{
			name:          `"TRUE" returns true`,
			key:           "TEST_GETBOOL_TRUE",
			preset:        "TRUE",
			setVar:        true,
			expectedValue: true,
			expectedOk:    true,
		},
		{
			name:          `"True" returns true`,
			key:           "TEST_GETBOOL_True",
			preset:        "True",
			setVar:        true,
			expectedValue: true,
			expectedOk:    true,
		},
		// false values
		{
			name:          `"0" returns false`,
			key:           "TEST_GETBOOL_0",
			preset:        "0",
			setVar:        true,
			expectedValue: false,
			expectedOk:    true,
		},
		{
			name:          `"f" returns false`,
			key:           "TEST_GETBOOL_f",
			preset:        "f",
			setVar:        true,
			expectedValue: false,
			expectedOk:    true,
		},
		{
			name:          `"F" returns false`,
			key:           "TEST_GETBOOL_F",
			preset:        "F",
			setVar:        true,
			expectedValue: false,
			expectedOk:    true,
		},
		{
			name:          `"false" returns false`,
			key:           "TEST_GETBOOL_false",
			preset:        "false",
			setVar:        true,
			expectedValue: false,
			expectedOk:    true,
		},
		{
			name:          `"FALSE" returns false`,
			key:           "TEST_GETBOOL_FALSE",
			preset:        "FALSE",
			setVar:        true,
			expectedValue: false,
			expectedOk:    true,
		},
		{
			name:          `"False" returns false`,
			key:           "TEST_GETBOOL_False",
			preset:        "False",
			setVar:        true,
			expectedValue: false,
			expectedOk:    true,
		},
		// invalid values
		{
			name:          "invalid string returns false",
			key:           "TEST_GETBOOL_INVALID",
			preset:        "not_a_bool",
			setVar:        true,
			expectedValue: false,
			expectedOk:    false,
		},
		{
			name:          "numeric string other than 0/1 returns false",
			key:           "TEST_GETBOOL_NUMERIC",
			preset:        "42",
			setVar:        true,
			expectedValue: false,
			expectedOk:    false,
		},
		{
			name:          "empty value returns false",
			key:           "TEST_GETBOOL_EMPTY",
			preset:        "",
			setVar:        true,
			expectedValue: false,
			expectedOk:    false,
		},
		{
			name:          "whitespace value returns false",
			key:           "TEST_GETBOOL_WHITESPACE",
			preset:        "   ",
			setVar:        true,
			expectedValue: false,
			expectedOk:    false,
		},
		{
			name:          "non-existent variable returns false",
			key:           "TEST_GETBOOL_NONEXISTENT",
			setVar:        false,
			expectedValue: false,
			expectedOk:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				t.Setenv(tt.key, tt.preset)
			}

			value, ok := lxenv.GetBool(tt.key)
			if value != tt.expectedValue || ok != tt.expectedOk {
				t.Errorf("GetBool(%q) = (%v, %v), want (%v, %v)", tt.key, value, ok, tt.expectedValue, tt.expectedOk)
			}
		})
	}
}

func TestGetBoolOr(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		preset       string
		setVar       bool
		defaultValue bool
		expected     bool
	}{
		{
			name:         "true value returns true",
			key:          "TEST_GETBOOLOR_TRUE",
			preset:       "true",
			setVar:       true,
			defaultValue: false,
			expected:     true,
		},
		{
			name:         "false value returns false",
			key:          "TEST_GETBOOLOR_FALSE",
			preset:       "false",
			setVar:       true,
			defaultValue: true,
			expected:     false,
		},
		{
			name:         `"1" returns true`,
			key:          "TEST_GETBOOLOR_1",
			preset:       "1",
			setVar:       true,
			defaultValue: false,
			expected:     true,
		},
		{
			name:         `"0" returns false`,
			key:          "TEST_GETBOOLOR_0",
			preset:       "0",
			setVar:       true,
			defaultValue: true,
			expected:     false,
		},
		{
			name:         "invalid value returns default",
			key:          "TEST_GETBOOLOR_INVALID",
			preset:       "not_a_bool",
			setVar:       true,
			defaultValue: true,
			expected:     true,
		},
		{
			name:         "empty value returns default",
			key:          "TEST_GETBOOLOR_EMPTY",
			preset:       "",
			setVar:       true,
			defaultValue: true,
			expected:     true,
		},
		{
			name:         "whitespace value returns default",
			key:          "TEST_GETBOOLOR_WHITESPACE",
			preset:       "   ",
			setVar:       true,
			defaultValue: true,
			expected:     true,
		},
		{
			name:         "non-existent variable returns default",
			key:          "TEST_GETBOOLOR_NONEXISTENT",
			setVar:       false,
			defaultValue: true,
			expected:     true,
		},
		{
			name:         "false default value",
			key:          "TEST_GETBOOLOR_FALSE_DEFAULT",
			setVar:       false,
			defaultValue: false,
			expected:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				t.Setenv(tt.key, tt.preset)
			}

			result := lxenv.GetBoolOr(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("GetBoolOr(%q, %v) = %v, want %v", tt.key, tt.defaultValue, result, tt.expected)
			}
		})
	}
}

func TestMustGetBool(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		value     string
		setEnv    bool
		want      bool
		wantPanic bool
	}{
		// true values
		{name: `"1" returns true`, key: "TEST_MUSTGETBOOL_1", value: "1", setEnv: true, want: true},
		{name: `"t" returns true`, key: "TEST_MUSTGETBOOL_t", value: "t", setEnv: true, want: true},
		{name: `"T" returns true`, key: "TEST_MUSTGETBOOL_T", value: "T", setEnv: true, want: true},
		{name: `"true" returns true`, key: "TEST_MUSTGETBOOL_true", value: "true", setEnv: true, want: true},
		{name: `"TRUE" returns true`, key: "TEST_MUSTGETBOOL_TRUE", value: "TRUE", setEnv: true, want: true},
		{name: `"True" returns true`, key: "TEST_MUSTGETBOOL_True", value: "True", setEnv: true, want: true},
		// false values
		{name: `"0" returns false`, key: "TEST_MUSTGETBOOL_0", value: "0", setEnv: true, want: false},
		{name: `"f" returns false`, key: "TEST_MUSTGETBOOL_f", value: "f", setEnv: true, want: false},
		{name: `"F" returns false`, key: "TEST_MUSTGETBOOL_F", value: "F", setEnv: true, want: false},
		{name: `"false" returns false`, key: "TEST_MUSTGETBOOL_false", value: "false", setEnv: true, want: false},
		{name: `"FALSE" returns false`, key: "TEST_MUSTGETBOOL_FALSE", value: "FALSE", setEnv: true, want: false},
		{name: `"False" returns false`, key: "TEST_MUSTGETBOOL_False", value: "False", setEnv: true, want: false},
		// edge / invalid
		{name: "invalid string panics", key: "TEST_MUSTGETBOOL_INVALID", value: "not_a_bool", setEnv: true, wantPanic: true},
		{name: "missing key panics", key: "TEST_MUSTGETBOOL_MISSING", setEnv: false, wantPanic: true},
		{name: "empty value panics", key: "TEST_MUSTGETBOOL_EMPTY", value: "", setEnv: true, wantPanic: true},
		{name: "whitespace value panics", key: "TEST_MUSTGETBOOL_SPACE", value: "   ", setEnv: true, wantPanic: true},
		{name: "trimmed value invalid (panics)", key: "TEST_MUSTGETBOOL_SPACED", value: " true ", setEnv: true, wantPanic: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setEnv {
				t.Setenv(tt.key, tt.value)
			} else {
				os.Unsetenv(tt.key)
			}

			var (
				got      bool
				didPanic bool
			)
			func() {
				defer func() {
					if r := recover(); r != nil {
						didPanic = true
					}
				}()
				got = lxenv.MustGetBool(tt.key)
			}()

			if tt.wantPanic {
				if !didPanic {
					t.Fatalf("MustGetBool(%q) did not panic", tt.key)
				}
				return
			}

			if didPanic {
				t.Fatalf("MustGetBool(%q) panicked unexpectedly", tt.key)
			}

			if got != tt.want {
				t.Fatalf("MustGetBool(%q) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}

func TestGetFloat(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		preset     string
		setVar     bool
		expected   float64
		expectedOk bool
	}{
		{"valid float", "TEST_GETFLOAT_VALID", "3.14", true, 3.14, true},
		{"negative float", "TEST_GETFLOAT_NEG", "-2.5", true, -2.5, true},
		{"plus sign", "TEST_GETFLOAT_PLUS", "+1.5", true, 1.5, true},
		{"integer value", "TEST_GETFLOAT_INT", "42", true, 42.0, true},
		{"scientific notation", "TEST_GETFLOAT_SCI", "1e6", true, 1e6, true},
		{"nan value", "TEST_GETFLOAT_NAN", "NaN", true, math.NaN(), true},
		{"inf value", "TEST_GETFLOAT_INF", "Inf", true, math.Inf(1), true},
		{"negative zero", "TEST_GETFLOAT_NEGZERO", "-0", true, 0, true},
		{"empty value", "TEST_GETFLOAT_EMPTY", "", true, 0, false},
		{"whitespace value", "TEST_GETFLOAT_SPACE", "   ", true, 0, false},
		{"invalid value", "TEST_GETFLOAT_INVALID", "not_float", true, 0, false},
		{"non-existent", "TEST_GETFLOAT_NONEXIST", "", false, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				t.Setenv(tt.key, tt.preset)
			} else {
				os.Unsetenv(tt.key)
			}

			v, ok := lxenv.GetFloat(tt.key)
			if ok != tt.expectedOk {
				t.Fatalf("GetFloat(%q) ok = %v, want %v", tt.key, ok, tt.expectedOk)
			}
			if ok {
				// handle NaN specially
				if math.IsNaN(tt.expected) {
					if !math.IsNaN(v) {
						t.Fatalf("GetFloat(%q) = %v, want NaN", tt.key, v)
					}
					return
				}
				// handle infinities
				if math.IsInf(tt.expected, 0) {
					if !math.IsInf(v, 1) && !math.IsInf(v, -1) {
						t.Fatalf("GetFloat(%q) = %v, want %v", tt.key, v, tt.expected)
					}
					return
				}
				if v != tt.expected {
					t.Fatalf("GetFloat(%q) = %v, want %v", tt.key, v, tt.expected)
				}
			}
		})
	}
}

func TestGetFloatOr(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		preset       string
		setVar       bool
		defaultValue float64
		expected     float64
	}{
		{"existing returns value", "TEST_GETFLOATOR_EXIST", "2.71", true, 0.0, 2.71},
		{"non-existent returns default", "TEST_GETFLOATOR_NONEXIST", "", false, 1.23, 1.23},
		{"invalid returns default", "TEST_GETFLOATOR_INVALID", "abc", true, 9.9, 9.9},
		{"empty returns default", "TEST_GETFLOATOR_EMPTY", "", true, 7.7, 7.7},
		{"scientific returns value", "TEST_GETFLOATOR_SCI", "-1e3", true, 0, -1000.0},
		{"inf returns inf", "TEST_GETFLOATOR_INF", "Inf", true, 0, math.Inf(1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				t.Setenv(tt.key, tt.preset)
			} else {
				os.Unsetenv(tt.key)
			}

			res := lxenv.GetFloatOr(tt.key, tt.defaultValue)
			// handle NaN/Inf if needed
			if math.IsInf(tt.expected, 0) {
				if !math.IsInf(res, 1) {
					t.Fatalf("GetFloatOr(%q) = %v, want %v", tt.key, res, tt.expected)
				}
				return
			}
			if math.IsNaN(tt.expected) {
				if !math.IsNaN(res) {
					t.Fatalf("GetFloatOr(%q) = %v, want NaN", tt.key, res)
				}
				return
			}
			if res != tt.expected {
				t.Fatalf("GetFloatOr(%q, %v) = %v, want %v", tt.key, tt.defaultValue, res, tt.expected)
			}
		})
	}
}

func TestMustGetFloat(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		preset    string
		setVar    bool
		want      float64
		wantPanic bool
	}{
		{"valid float", "TEST_MUSTGETFLOAT_VALID", "6.28", true, 6.28, false},
		{"missing panics", "TEST_MUSTGETFLOAT_MISSING", "", false, 0, true},
		{"invalid panics", "TEST_MUSTGETFLOAT_INVALID", "notfloat", true, 0, true},
		{"empty panics", "TEST_MUSTGETFLOAT_EMPTY", "", true, 0, true},
		{"whitespace panics", "TEST_MUSTGETFLOAT_SPACE", "  ", true, 0, true},
		{"plus sign", "TEST_MUSTGETFLOAT_PLUS", "+4.0", true, 4.0, false},
		{"scientific", "TEST_MUSTGETFLOAT_SCI", "1.2e3", true, 1200.0, false},
		{"NaN handled", "TEST_MUSTGETFLOAT_NAN", "NaN", true, math.NaN(), false},
		{"Inf handled", "TEST_MUSTGETFLOAT_INF", "-Inf", true, math.Inf(-1), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				t.Setenv(tt.key, tt.preset)
			} else {
				os.Unsetenv(tt.key)
			}

			var (
				got      float64
				didPanic bool
			)
			func() {
				defer func() {
					if r := recover(); r != nil {
						didPanic = true
					}
				}()
				got = lxenv.MustGetFloat(tt.key)
			}()

			if tt.wantPanic {
				if !didPanic {
					t.Fatalf("MustGetFloat(%q) did not panic", tt.key)
				}
				return
			}

			if didPanic {
				t.Fatalf("MustGetFloat(%q) panicked unexpectedly", tt.key)
			}

			if math.IsNaN(tt.want) {
				if !math.IsNaN(got) {
					t.Fatalf("MustGetFloat(%q) = %v, want NaN", tt.key, got)
				}
				return
			}
			if math.IsInf(tt.want, 0) {
				if !math.IsInf(got, -1) && !math.IsInf(got, 1) {
					t.Fatalf("MustGetFloat(%q) = %v, want %v", tt.key, got, tt.want)
				}
				return
			}

			if got != tt.want {
				t.Fatalf("MustGetFloat(%q) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}

func TestRequire(t *testing.T) {
	tests := []struct {
		name       string
		keys       []string // keys to pass to Require
		preset     map[string]string
		expectsErr bool
		missing    []string // expected missing keys included in the error text
	}{
		{
			name:       "all keys present",
			keys:       []string{"TEST_REQUIRE_K1", "TEST_REQUIRE_K2"},
			preset:     map[string]string{"TEST_REQUIRE_K1": "v1", "TEST_REQUIRE_K2": "v2"},
			expectsErr: false,
		},
		{
			name:       "single missing key",
			keys:       []string{"TEST_REQUIRE_PRESENT", "TEST_REQUIRE_MISSING"},
			preset:     map[string]string{"TEST_REQUIRE_PRESENT": "v"},
			expectsErr: true,
			missing:    []string{"TEST_REQUIRE_MISSING"},
		},
		{
			name:       "multiple missing keys",
			keys:       []string{"TEST_REQUIRE_A", "TEST_REQUIRE_M1", "TEST_REQUIRE_M2"},
			preset:     map[string]string{"TEST_REQUIRE_A": "v"},
			expectsErr: true,
			missing:    []string{"TEST_REQUIRE_M1", "TEST_REQUIRE_M2"},
		},
		{
			name:       "none present",
			keys:       []string{"TEST_REQUIRE_NONE"},
			preset:     map[string]string{},
			expectsErr: true,
			missing:    []string{"TEST_REQUIRE_NONE"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original values for all relevant keys so we can restore after the case.
			orig := make(map[string]*string)
			allKeys := make(map[string]struct{})
			for _, k := range tt.keys {
				allKeys[k] = struct{}{}
			}
			for k := range tt.preset {
				allKeys[k] = struct{}{}
			}
			for k := range allKeys {
				if v, ok := os.LookupEnv(k); ok {
					val := v
					orig[k] = &val
				} else {
					orig[k] = nil
				}
				// ensure clean state before setting up
				os.Unsetenv(k)
			}

			// Setup preset env vars for this case
			for k, v := range tt.preset {
				if err := os.Setenv(k, v); err != nil {
					t.Fatalf("failed to set env %s: %v", k, err)
				}
			}

			// Call Require
			err := lxenv.Require(tt.keys...)

			// Restore originals
			for k, vptr := range orig {
				if vptr == nil {
					os.Unsetenv(k)
				} else {
					os.Setenv(k, *vptr)
				}
			}

			if tt.expectsErr {
				if err == nil {
					t.Fatalf("Require(%v) expected error, got nil", tt.keys)
				}
				// check sentinel
				if !errors.Is(err, lxenv.ErrKeyNotFound) {
					t.Fatalf("Require(%v) expected ErrKeyNotFound, got: %v", tt.keys, err)
				}
				// check missing keys mentioned in the error text
				for _, missing := range tt.missing {
					if !strings.Contains(err.Error(), missing) {
						t.Fatalf("error %q did not mention missing key %q", err, missing)
					}
				}
				return
			}

			// expects no error
			if err != nil {
				t.Fatalf("Require(%v) returned unexpected error: %v", tt.keys, err)
			}
		})
	}
}
