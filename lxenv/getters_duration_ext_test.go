package lxenv_test

import (
	"os"
	"testing"
	"time"

	"github.com/nthanhhai2909/lx/lxenv"
)

func TestGetDuration(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		preset        string
		setVar        bool
		expectedValue time.Duration
		expectedOk    bool
	}{
		{
			name:          "standard duration",
			key:           "TEST_DURATION_STD",
			preset:        "1h30m",
			setVar:        true,
			expectedValue: 90 * time.Minute,
			expectedOk:    true,
		},
		{
			name:          "extended duration days",
			key:           "TEST_DURATION_DAYS",
			preset:        "3d",
			setVar:        true,
			expectedValue: 72 * time.Hour,
			expectedOk:    true,
		},
		{
			name:          "case insensitive unit",
			key:           "TEST_DURATION_CASE",
			preset:        "1 DAY",
			setVar:        true,
			expectedValue: 24 * time.Hour,
			expectedOk:    true,
		},
		{
			name:          "invalid duration returns false",
			key:           "TEST_DURATION_INVALID",
			preset:        "not_a_duration",
			setVar:        true,
			expectedValue: 0,
			expectedOk:    false,
		},
		{
			name:          "non-existent variable returns false",
			key:           "TEST_DURATION_NONEXISTENT",
			setVar:        false,
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

			value, ok := lxenv.GetDuration(tt.key)
			if value != tt.expectedValue || ok != tt.expectedOk {
				t.Errorf("GetDuration(%q) = (%v, %v), want (%v, %v)", tt.key, value, ok, tt.expectedValue, tt.expectedOk)
			}
		})
	}
}

func TestGetDurationOr(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		preset       string
		setVar       bool
		defaultValue time.Duration
		expected     time.Duration
	}{
		{
			name:         "valid duration returns value",
			key:          "TEST_DURATIONOR_VALID",
			preset:       "2d",
			setVar:       true,
			defaultValue: 1 * time.Hour,
			expected:     48 * time.Hour,
		},
		{
			name:         "invalid duration returns default",
			key:          "TEST_DURATIONOR_INVALID",
			preset:       "invalid",
			setVar:       true,
			defaultValue: 5 * time.Minute,
			expected:     5 * time.Minute,
		},
		{
			name:         "non-existent variable returns default",
			key:          "TEST_DURATIONOR_NONEXISTENT",
			setVar:       false,
			defaultValue: 10 * time.Second,
			expected:     10 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				os.Setenv(tt.key, tt.preset)
				defer os.Unsetenv(tt.key)
			}

			result := lxenv.GetDurationOr(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("GetDurationOr(%q, %v) = %v, want %v", tt.key, tt.defaultValue, result, tt.expected)
			}
		})
	}
}
