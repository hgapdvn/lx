package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxtime"
)

func TestDays_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected time.Duration
	}{
		{
			name:     "zero days",
			input:    0,
			expected: 0,
		},
		{
			name:     "one day",
			input:    1,
			expected: 24 * time.Hour,
		},
		{
			name:     "five days",
			input:    5,
			expected: 120 * time.Hour,
		},
		{
			name:     "seven days",
			input:    7,
			expected: 168 * time.Hour,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Days(tt.input)
			if result != tt.expected {
				t.Errorf("Days(%d) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDays_Usage(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "can be used with Add",
			check: func() bool {
				base := time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC)
				result := base.Add(lxtime.Days(5))
				expected := time.Date(2026, 4, 9, 0, 0, 0, 0, time.UTC)
				return result.Equal(expected)
			},
		},
		{
			name: "days value equals 24 hours",
			check: func() bool {
				return lxtime.Days(1) == 24*time.Hour
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("Days() usage check failed")
			}
		})
	}
}

func TestDays_Negative(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "negative days work",
			check: func() bool {
				base := time.Date(2026, 4, 9, 0, 0, 0, 0, time.UTC)
				result := base.Add(lxtime.Days(-5))
				expected := time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC)
				return result.Equal(expected)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("Days() negative check failed")
			}
		})
	}
}
