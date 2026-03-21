package lxenv

import (
	"os"
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
		wantErr  bool
	}{
		// Standard units
		{"1h", time.Hour, false},
		{"30m", 30 * time.Minute, false},
		{"10s", 10 * time.Second, false},
		{"100ms", 100 * time.Millisecond, false},
		{"100us", 100 * time.Microsecond, false},
		{"100µs", 100 * time.Microsecond, false},
		{"100μs", 100 * time.Microsecond, false},
		{"100ns", 100 * time.Nanosecond, false},

		// Extended units
		{"3d", 72 * time.Hour, false},
		{"1w", 168 * time.Hour, false},
		{"1y", 365 * 24 * time.Hour, false},
		{"1.5d", 36 * time.Hour, false},
		{"0.5w", 3.5 * 24 * time.Hour, false},

		// Combinations
		{"1h30m", 90 * time.Minute, false},
		{"1d12h", 36 * time.Hour, false},
		{"1w2d", 9 * 24 * time.Hour, false},
		{"1y1w1d1h1m1s", (365*24+7*24+24+1)*time.Hour + 1*time.Minute + 1*time.Second, false},

		// Case insensitivity and full names
		{"1 DAY", 24 * time.Hour, false},
		{"2 days", 48 * time.Hour, false},
		{"1 week", 168 * time.Hour, false},
		{"1 Year", 365 * 24 * time.Hour, false},
		{"1 hr", time.Hour, false},
		{"1 min", time.Minute, false},
		{"1 sec", time.Second, false},
		{"1 msec", time.Millisecond, false},
		{"1 usec", time.Microsecond, false},
		{"1 nsec", time.Nanosecond, false},

		// Spaces and signs
		{" 1d ", 24 * time.Hour, false},
		{"1 d", 24 * time.Hour, false},
		{"+1d", 24 * time.Hour, false},
		{"-1d", -24 * time.Hour, false},
		{"1d 2h 3m", 24*time.Hour + 2*time.Hour + 3*time.Minute, false},
		{"- 1d", 0, true}, // Invalid space between sign and value

		// Floating point
		{"0.5h", 30 * time.Minute, false},
		{".5h", 30 * time.Minute, false},
		{"1.h", time.Hour, false}, // Standard ParseDuration allows this

		// Errors
		{"", 0, true},
		{"abc", 0, true},
		{"1x", 0, true},
		{"1.2.3d", 0, true},
		{"1d 2", 0, true}, // Missing unit for the second part
		{"d", 0, true},    // Missing value
		{"1.2.3", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := parseDuration(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDuration(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("parseDuration(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

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
			name:          "extended duration weeks",
			key:           "TEST_DURATION_WEEKS",
			preset:        "2w",
			setVar:        true,
			expectedValue: 336 * time.Hour,
			expectedOk:    true,
		},
		{
			name:          "extended duration years",
			key:           "TEST_DURATION_YEARS",
			preset:        "1y",
			setVar:        true,
			expectedValue: 8760 * time.Hour,
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
			name:          "combination with spaces",
			key:           "TEST_DURATION_COMBO_SPACES",
			preset:        "1d 2h 30m",
			setVar:        true,
			expectedValue: 26*time.Hour + 30*time.Minute,
			expectedOk:    true,
		},
		{
			name:          "negative duration",
			key:           "TEST_DURATION_NEGATIVE",
			preset:        "-1.5h",
			setVar:        true,
			expectedValue: -90 * time.Minute,
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
			name:          "empty value returns false",
			key:           "TEST_DURATION_EMPTY",
			preset:        "",
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
				t.Setenv(tt.key, tt.preset)
			}

			value, ok := GetDuration(tt.key)
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
			name:         "combination returns value",
			key:          "TEST_DURATIONOR_COMBO",
			preset:       "1w 2d",
			setVar:       true,
			defaultValue: 1 * time.Hour,
			expected:     216 * time.Hour,
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
			name:         "empty duration returns default",
			key:          "TEST_DURATIONOR_EMPTY",
			preset:       "",
			setVar:       true,
			defaultValue: 10 * time.Minute,
			expected:     10 * time.Minute,
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
				t.Setenv(tt.key, tt.preset)
			}

			result := GetDurationOr(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("GetDurationOr(%q, %v) = %v, want %v", tt.key, tt.defaultValue, result, tt.expected)
			}
		})
	}
}

func TestMustGetDuration(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		preset    string
		setVar    bool
		want      time.Duration
		wantPanic bool
	}{
		{
			name:   "standard duration",
			key:    "TEST_MUST_DURATION_STD",
			preset: "1h30m",
			setVar: true,
			want:   90 * time.Minute,
		},
		{
			name:   "extended days",
			key:    "TEST_MUST_DURATION_DAYS",
			preset: "3d",
			setVar: true,
			want:   72 * time.Hour,
		},
		{
			name:   "combination with spaces",
			key:    "TEST_MUST_DURATION_COMBO",
			preset: "1d 2h 30m",
			setVar: true,
			want:   26*time.Hour + 30*time.Minute,
		},
		{
			name:   "negative duration",
			key:    "TEST_MUST_DURATION_NEG",
			preset: "-1.5h",
			setVar: true,
			want:   -90 * time.Minute,
		},
		{
			name:      "missing key panics",
			key:       "TEST_MUST_DURATION_MISSING",
			setVar:    false,
			wantPanic: true,
		},
		{
			name:      "invalid duration panics",
			key:       "TEST_MUST_DURATION_INVALID",
			preset:    "not_a_duration",
			setVar:    true,
			wantPanic: true,
		},
		{
			name:      "empty value panics",
			key:       "TEST_MUST_DURATION_EMPTY",
			preset:    "",
			setVar:    true,
			wantPanic: true,
		},
		{
			name:   "whitespace allowed",
			key:    "TEST_MUST_DURATION_SPACE",
			preset: " 1d ",
			setVar: true,
			want:   24 * time.Hour,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVar {
				t.Setenv(tt.key, tt.preset)
			} else {
				os.Unsetenv(tt.key)
			}

			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Fatalf("MustGetDuration(%q) did not panic", tt.key)
					}
				}()
			}

			got := MustGetDuration(tt.key)
			if !tt.wantPanic && got != tt.want {
				t.Fatalf("MustGetDuration(%q) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}
