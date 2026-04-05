package lxtime_test

import (
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/lxtime"
)

func TestStartOfYear_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "January 1 stays January 1",
			input:    time.Date(2026, 1, 1, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "mid-year goes to January 1",
			input:    time.Date(2026, 6, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "December 31 goes to January 1",
			input:    time.Date(2026, 12, 31, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.StartOfYear(tt.input)
			if !result.Equal(tt.expected) {
				t.Errorf("StartOfYear(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStartOfYear_PreservesTimezone(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "preserves UTC",
			check: func() bool {
				input := time.Date(2026, 6, 15, 15, 30, 0, 0, time.UTC)
				result := lxtime.StartOfYear(input)
				return result.Location() == time.UTC
			},
		},
		{
			name: "preserves EST",
			check: func() bool {
				est, _ := time.LoadLocation("America/New_York")
				input := time.Date(2026, 6, 15, 15, 30, 0, 0, est)
				result := lxtime.StartOfYear(input)
				return result.Location() == est
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("StartOfYear() timezone check failed")
			}
		})
	}
}

func TestStartOfYear_DifferentYears(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "2024 goes to Jan 1, 2024",
			check: func() bool {
				input := time.Date(2024, 6, 15, 15, 30, 0, 0, time.UTC)
				result := lxtime.StartOfYear(input)
				return result.Year() == 2024 && result.Month() == time.January && result.Day() == 1
			},
		},
		{
			name: "2025 goes to Jan 1, 2025",
			check: func() bool {
				input := time.Date(2025, 6, 15, 15, 30, 0, 0, time.UTC)
				result := lxtime.StartOfYear(input)
				return result.Year() == 2025 && result.Month() == time.January && result.Day() == 1
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("StartOfYear() different years check failed")
			}
		})
	}
}
