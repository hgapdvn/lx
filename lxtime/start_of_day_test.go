package lxtime_test

import (
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/lxtime"
)

func TestStartOfDay_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "midnight stays midnight",
			input:    time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "noon becomes midnight",
			input:    time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "end of day becomes midnight",
			input:    time.Date(2026, 4, 4, 23, 59, 59, 999999999, time.UTC),
			expected: time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.StartOfDay(tt.input)
			if !result.Equal(tt.expected) {
				t.Errorf("StartOfDay() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestStartOfDay_DifferentDates(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "preserves year",
			check: func() bool {
				input := time.Date(2026, 4, 4, 15, 30, 45, 123456789, time.UTC)
				result := lxtime.StartOfDay(input)
				return result.Year() == 2026
			},
		},
		{
			name: "preserves month",
			check: func() bool {
				input := time.Date(2026, 4, 4, 15, 30, 45, 123456789, time.UTC)
				result := lxtime.StartOfDay(input)
				return result.Month() == time.April
			},
		},
		{
			name: "preserves day",
			check: func() bool {
				input := time.Date(2026, 4, 4, 15, 30, 45, 123456789, time.UTC)
				result := lxtime.StartOfDay(input)
				return result.Day() == 4
			},
		},
		{
			name: "sets hour to 0",
			check: func() bool {
				input := time.Date(2026, 4, 4, 15, 30, 45, 123456789, time.UTC)
				result := lxtime.StartOfDay(input)
				return result.Hour() == 0
			},
		},
		{
			name: "sets minute to 0",
			check: func() bool {
				input := time.Date(2026, 4, 4, 15, 30, 45, 123456789, time.UTC)
				result := lxtime.StartOfDay(input)
				return result.Minute() == 0
			},
		},
		{
			name: "sets second to 0",
			check: func() bool {
				input := time.Date(2026, 4, 4, 15, 30, 45, 123456789, time.UTC)
				result := lxtime.StartOfDay(input)
				return result.Second() == 0
			},
		},
		{
			name: "sets nanosecond to 0",
			check: func() bool {
				input := time.Date(2026, 4, 4, 15, 30, 45, 123456789, time.UTC)
				result := lxtime.StartOfDay(input)
				return result.Nanosecond() == 0
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("StartOfDay() check failed")
			}
		})
	}
}

func TestStartOfDay_DifferentTimezones(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "works with UTC",
			check: func() bool {
				input := time.Date(2026, 4, 4, 15, 30, 0, 0, time.UTC)
				result := lxtime.StartOfDay(input)
				return result.Hour() == 0 && result.Minute() == 0
			},
		},
		{
			name: "preserves timezone",
			check: func() bool {
				est, _ := time.LoadLocation("America/New_York")
				input := time.Date(2026, 4, 4, 15, 30, 0, 0, est)
				result := lxtime.StartOfDay(input)
				return result.Location() == est
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("StartOfDay() check failed")
			}
		})
	}
}

func TestStartOfDay_EdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "leap year day",
			check: func() bool {
				input := time.Date(2024, 2, 29, 15, 30, 0, 0, time.UTC)
				result := lxtime.StartOfDay(input)
				expected := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)
				return result.Equal(expected)
			},
		},
		{
			name: "month boundary",
			check: func() bool {
				input := time.Date(2026, 4, 1, 15, 30, 0, 0, time.UTC)
				result := lxtime.StartOfDay(input)
				expected := time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC)
				return result.Equal(expected)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("StartOfDay() check failed")
			}
		})
	}
}

func TestStartOfDay_TimezoneCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "EST preserves correct date",
			check: func() bool {
				est, _ := time.LoadLocation("America/New_York")
				input := time.Date(2026, 4, 4, 10, 30, 0, 0, est)
				result := lxtime.StartOfDay(input)
				expected := time.Date(2026, 4, 4, 0, 0, 0, 0, est)
				return result.Equal(expected)
			},
		},
		{
			name: "PST preserves correct date",
			check: func() bool {
				pst, _ := time.LoadLocation("America/Los_Angeles")
				input := time.Date(2026, 4, 4, 10, 30, 0, 0, pst)
				result := lxtime.StartOfDay(input)
				expected := time.Date(2026, 4, 4, 0, 0, 0, 0, pst)
				return result.Equal(expected)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("StartOfDay() timezone check failed")
			}
		})
	}
}
