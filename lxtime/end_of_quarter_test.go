package lxtime_test

import (
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/lxtime"
)

func TestEndOfQuarter_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Q1 - January",
			input:    time.Date(2026, 1, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 3, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q1 - February",
			input:    time.Date(2026, 2, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 3, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q1 - March",
			input:    time.Date(2026, 3, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 3, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q2 - April",
			input:    time.Date(2026, 4, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 6, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q2 - May",
			input:    time.Date(2026, 5, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 6, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q2 - June",
			input:    time.Date(2026, 6, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 6, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q3 - July",
			input:    time.Date(2026, 7, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 9, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q3 - August",
			input:    time.Date(2026, 8, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 9, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q3 - September",
			input:    time.Date(2026, 9, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 9, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q4 - October",
			input:    time.Date(2026, 10, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q4 - November",
			input:    time.Date(2026, 11, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q4 - December",
			input:    time.Date(2026, 12, 15, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2026, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.EndOfQuarter(tt.input)
			if !result.Equal(tt.expected) {
				t.Errorf("EndOfQuarter(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestEndOfQuarter_EdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "quarter end day",
			check: func() bool {
				input := time.Date(2026, 6, 30, 10, 30, 0, 0, time.UTC)
				result := lxtime.EndOfQuarter(input)
				expected := time.Date(2026, 6, 30, 23, 59, 59, 999999999, time.UTC)
				return result.Equal(expected)
			},
		},
		{
			name: "quarter start day",
			check: func() bool {
				input := time.Date(2026, 4, 1, 10, 30, 0, 0, time.UTC)
				result := lxtime.EndOfQuarter(input)
				expected := time.Date(2026, 6, 30, 23, 59, 59, 999999999, time.UTC)
				return result.Equal(expected)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("EndOfQuarter() check failed")
			}
		})
	}
}

func TestEndOfQuarter_PreservesTimezone(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "preserves UTC",
			check: func() bool {
				input := time.Date(2026, 6, 15, 15, 30, 0, 0, time.UTC)
				result := lxtime.EndOfQuarter(input)
				return result.Location() == time.UTC
			},
		},
		{
			name: "preserves EST",
			check: func() bool {
				est, _ := time.LoadLocation("America/New_York")
				input := time.Date(2026, 6, 15, 15, 30, 0, 0, est)
				result := lxtime.EndOfQuarter(input)
				return result.Location() == est
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("EndOfQuarter() timezone check failed")
			}
		})
	}
}

func TestStartEndOfQuarter_Consistency(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "start and end are same year",
			check: func() bool {
				input := time.Date(2026, 6, 15, 15, 30, 0, 0, time.UTC)
				start := lxtime.StartOfQuarter(input)
				end := lxtime.EndOfQuarter(input)
				return start.Year() == end.Year()
			},
		},
		{
			name: "start is first day, end is last day of same quarter",
			check: func() bool {
				input := time.Date(2026, 6, 15, 15, 30, 0, 0, time.UTC)
				start := lxtime.StartOfQuarter(input)
				end := lxtime.EndOfQuarter(input)
				return start.Day() == 1 && end.Day() >= 28
			},
		},
		{
			name: "end is after start",
			check: func() bool {
				input := time.Date(2026, 6, 15, 15, 30, 0, 0, time.UTC)
				start := lxtime.StartOfQuarter(input)
				end := lxtime.EndOfQuarter(input)
				return end.After(start)
			},
		},
		{
			name: "Q1 and Q2 don't overlap",
			check: func() bool {
				q1End := lxtime.EndOfQuarter(time.Date(2026, 3, 15, 0, 0, 0, 0, time.UTC))
				q2Start := lxtime.StartOfQuarter(time.Date(2026, 4, 15, 0, 0, 0, 0, time.UTC))
				return q1End.Before(q2Start)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("StartOfQuarter/EndOfQuarter consistency check failed")
			}
		})
	}
}

func TestEndOfQuarter_YearBoundary(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "Q4 ends on December 31",
			check: func() bool {
				input := time.Date(2025, 12, 15, 15, 30, 0, 0, time.UTC)
				result := lxtime.EndOfQuarter(input)
				return result.Month() == time.December && result.Day() == 31 && result.Year() == 2025
			},
		},
		{
			name: "next year Q1 starts on January 1",
			check: func() bool {
				input := time.Date(2026, 1, 15, 15, 30, 0, 0, time.UTC)
				result := lxtime.StartOfQuarter(input)
				return result.Month() == time.January && result.Day() == 1 && result.Year() == 2026
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("EndOfQuarter() year boundary check failed")
			}
		})
	}
}
