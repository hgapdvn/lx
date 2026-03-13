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
				os.Setenv(tt.key, tt.preset)
				defer os.Unsetenv(tt.key)
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
				os.Setenv(tt.key, tt.preset)
				defer os.Unsetenv(tt.key)
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
			name:         "whitespace value returns default",
			key:          "TEST_GETINTOR_WHITESPACE",
			preset:       "   ",
			setVar:       true,
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
				os.Setenv(tt.key, tt.preset)
				defer os.Unsetenv(tt.key)
			}

			result := lxenv.GetIntOr(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("GetIntOr(%q, %d) = %d, want %d", tt.key, tt.defaultValue, result, tt.expected)
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
				os.Setenv(tt.key, tt.preset)
				defer os.Unsetenv(tt.key)
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
				os.Setenv(tt.key, tt.preset)
				defer os.Unsetenv(tt.key)
			}

			result := lxenv.GetBoolOr(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("GetBoolOr(%q, %v) = %v, want %v", tt.key, tt.defaultValue, result, tt.expected)
			}
		})
	}
}
