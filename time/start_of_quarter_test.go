package lxtime_test

import (
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/time"
)

func TestStartOfQuarter_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Q1 - January",
			input:    time.Date(2026, 1, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q1 - February",
			input:    time.Date(2026, 2, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q1 - March",
			input:    time.Date(2026, 3, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q2 - April",
			input:    time.Date(2026, 4, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q2 - May",
			input:    time.Date(2026, 5, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q2 - June",
			input:    time.Date(2026, 6, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q3 - July",
			input:    time.Date(2026, 7, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 7, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q3 - August",
			input:    time.Date(2026, 8, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 7, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q3 - September",
			input:    time.Date(2026, 9, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 7, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q4 - October",
			input:    time.Date(2026, 10, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q4 - November",
			input:    time.Date(2026, 11, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q4 - December",
			input:    time.Date(2026, 12, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.StartOfQuarter(tt.input)
			if !result.Equal(tt.expected) {
				t.Errorf("StartOfQuarter(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStartOfQuarter_EdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "quarter start day",
			check: func() bool {
				input := time.Date(2026, 4, 1, 10, 30, 0, 0, time.UTC)
				result := lxtime.StartOfQuarter(input)
				expected := time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC)
				return result.Equal(expected)
			},
		},
		{
			name: "quarter end day",
			check: func() bool {
				input := time.Date(2026, 6, 30, 10, 30, 0, 0, time.UTC)
				result := lxtime.StartOfQuarter(input)
				expected := time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC)
				return result.Equal(expected)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("StartOfQuarter() check failed")
			}
		})
	}
}

func TestStartOfQuarter_PreservesTimezone(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "preserves UTC",
			check: func() bool {
				input := time.Date(2026, 6, 15, 15, 30, 0, 0, time.UTC)
				result := lxtime.StartOfQuarter(input)
				return result.Location() == time.UTC
			},
		},
		{
			name: "preserves EST",
			check: func() bool {
				est, _ := time.LoadLocation("America/New_York")
				input := time.Date(2026, 6, 15, 15, 30, 0, 0, est)
				result := lxtime.StartOfQuarter(input)
				return result.Location() == est
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("StartOfQuarter() timezone check failed")
			}
		})
	}
}
