package lxtime_test

import (
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/time"
)

func TestSeconds_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected time.Duration
	}{
		{
			name:     "zero seconds",
			input:    0,
			expected: 0,
		},
		{
			name:     "one second",
			input:    1,
			expected: time.Second,
		},
		{
			name:     "forty five seconds",
			input:    45,
			expected: 45 * time.Second,
		},
		{
			name:     "sixty seconds",
			input:    60,
			expected: 60 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Seconds(tt.input)
			if result != tt.expected {
				t.Errorf("Seconds(%d) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSeconds_Usage(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "can be used with Add",
			check: func() bool {
				base := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
				result := base.Add(lxtime.Seconds(30))
				expected := time.Date(2026, 4, 4, 10, 30, 30, 0, time.UTC)
				return result.Equal(expected)
			},
		},
		{
			name: "seconds value equals time.Second",
			check: func() bool {
				return lxtime.Seconds(1) == time.Second
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("Seconds() usage check failed")
			}
		})
	}
}

func TestSeconds_Negative(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "negative seconds work",
			check: func() bool {
				base := time.Date(2026, 4, 4, 10, 30, 30, 0, time.UTC)
				result := base.Add(lxtime.Seconds(-30))
				expected := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
				return result.Equal(expected)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("Seconds() negative check failed")
			}
		})
	}
}
