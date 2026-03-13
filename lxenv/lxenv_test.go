package lxenv_test

import (
	"os"
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
				os.Setenv(tt.key, tt.setValue)
				defer os.Unsetenv(tt.key)
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
				os.Setenv(tt.key, tt.setValue)
				defer os.Unsetenv(tt.key)
			}

			result := lxenv.GetOr(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("GetOr(%q, %q) = %q, want %q", tt.key, tt.defaultValue, result, tt.expected)
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
				os.Setenv(tt.key, tt.preset)
				defer os.Unsetenv(tt.key)
			}

			result := lxenv.Has(tt.key)
			if result != tt.expected {
				t.Errorf("Has(%q) = %v, want %v", tt.key, result, tt.expected)
			}
		})
	}
}
