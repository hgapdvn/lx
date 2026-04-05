package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxtime"
)

func TestWeeks(t *testing.T) {
	tests := []struct {
		name     string
		weeks    int
		expected time.Duration
	}{
		{
			name:     "zero weeks",
			weeks:    0,
			expected: 0,
		},
		{
			name:     "single week",
			weeks:    1,
			expected: 7 * 24 * time.Hour,
		},
		{
			name:     "multiple weeks",
			weeks:    2,
			expected: 14 * 24 * time.Hour,
		},
		{
			name:     "large number of weeks",
			weeks:    52,
			expected: 52 * 7 * 24 * time.Hour,
		},
		{
			name:     "negative weeks",
			weeks:    -3,
			expected: -3 * 7 * 24 * time.Hour,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Weeks(tt.weeks)
			if result != tt.expected {
				t.Errorf("Weeks(%d) = %v, want %v", tt.weeks, result, tt.expected)
			}
		})
	}
}

func ExampleWeeks() {
	duration := lxtime.Weeks(2)
	// duration: 336h0m0s
	_ = duration
}
