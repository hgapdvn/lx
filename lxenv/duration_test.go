package lxenv

import (
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

		// Extended units
		{"3d", 72 * time.Hour, false},
		{"1w", 168 * time.Hour, false},
		{"1y", 365 * 24 * time.Hour, false},
		{"1.5d", 36 * time.Hour, false},

		// Combinations
		{"1h30m", 90 * time.Minute, false},
		{"1d12h", 36 * time.Hour, false},
		{"1w2d", 9 * 24 * time.Hour, false},

		// Case insensitivity and full names
		{"1 DAY", 24 * time.Hour, false},
		{"2 days", 48 * time.Hour, false},
		{"1 week", 168 * time.Hour, false},
		{"1 Year", 365 * 24 * time.Hour, false},

		// Spaces and signs
		{" 1d ", 24 * time.Hour, false},
		{"1 d", 24 * time.Hour, false},
		{"+1d", 24 * time.Hour, false},
		{"-1d", -24 * time.Hour, false},

		// Errors
		{"", 0, true},
		{"abc", 0, true},
		{"1x", 0, true},
		{"1.2.3d", 0, true},
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
