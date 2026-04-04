package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/time"
)

func TestMinutes_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected time.Duration
	}{
		{
			name:     "zero minutes",
			input:    0,
			expected: 0,
		},
		{
			name:     "one minute",
			input:    1,
			expected: time.Minute,
		},
		{
			name:     "thirty minutes",
			input:    30,
			expected: 30 * time.Minute,
		},
		{
			name:     "sixty minutes",
			input:    60,
			expected: 60 * time.Minute,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Minutes(tt.input)
			if result != tt.expected {
				t.Errorf("Minutes(%d) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMinutes_Usage(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "can be used with Add",
			check: func() bool {
				base := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
				result := base.Add(lxtime.Minutes(15))
				expected := time.Date(2026, 4, 4, 10, 45, 0, 0, time.UTC)
				return result.Equal(expected)
			},
		},
		{
			name: "minutes value equals time.Minute",
			check: func() bool {
				return lxtime.Minutes(1) == time.Minute
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("Minutes() usage check failed")
			}
		})
	}
}

func TestMinutes_Negative(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "negative minutes work",
			check: func() bool {
				base := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
				result := base.Add(lxtime.Minutes(-15))
				expected := time.Date(2026, 4, 4, 10, 15, 0, 0, time.UTC)
				return result.Equal(expected)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("Minutes() negative check failed")
			}
		})
	}
}
