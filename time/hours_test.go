package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/time"
)

func TestHours_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected time.Duration
	}{
		{
			name:     "zero hours",
			input:    0,
			expected: 0,
		},
		{
			name:     "one hour",
			input:    1,
			expected: time.Hour,
		},
		{
			name:     "three hours",
			input:    3,
			expected: 3 * time.Hour,
		},
		{
			name:     "twenty four hours",
			input:    24,
			expected: 24 * time.Hour,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Hours(tt.input)
			if result != tt.expected {
				t.Errorf("Hours(%d) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestHours_Usage(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "can be used with Add",
			check: func() bool {
				base := time.Date(2026, 4, 4, 10, 0, 0, 0, time.UTC)
				result := base.Add(lxtime.Hours(5))
				expected := time.Date(2026, 4, 4, 15, 0, 0, 0, time.UTC)
				return result.Equal(expected)
			},
		},
		{
			name: "hours value equals time.Hour",
			check: func() bool {
				return lxtime.Hours(1) == time.Hour
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("Hours() usage check failed")
			}
		})
	}
}

func TestHours_Negative(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "negative hours work",
			check: func() bool {
				base := time.Date(2026, 4, 4, 15, 0, 0, 0, time.UTC)
				result := base.Add(lxtime.Hours(-5))
				expected := time.Date(2026, 4, 4, 10, 0, 0, 0, time.UTC)
				return result.Equal(expected)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("Hours() negative check failed")
			}
		})
	}
}
