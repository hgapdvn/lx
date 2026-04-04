package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/time"
)

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected bool
	}{
		{
			name:     "regular leap year (divisible by 4)",
			date:     time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "regular non-leap year",
			date:     time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "leap year divisible by 4",
			date:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "non-leap year: century year not divisible by 400",
			date:     time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "leap year: century year divisible by 400",
			date:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "non-leap year: 2100",
			date:     time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "non-leap year: 2200",
			date:     time.Date(2200, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "leap year: 2400",
			date:     time.Date(2400, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "historical leap year: 1904",
			date:     time.Date(1904, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "historical non-leap year: 1801",
			date:     time.Date(1801, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "leap year with different month and day",
			date:     time.Date(2024, 2, 29, 15, 30, 45, 0, time.UTC),
			expected: true,
		},
		{
			name:     "non-leap year with different month and day",
			date:     time.Date(2026, 5, 15, 10, 20, 30, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsLeapYear(tt.date)
			if result != tt.expected {
				t.Errorf("IsLeapYear(%v) = %v, want %v", tt.date, result, tt.expected)
			}
		})
	}
}

func ExampleIsLeapYear() {
	t := time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC)
	isLeap := lxtime.IsLeapYear(t)
	// isLeap: true
	_ = isLeap

	t2 := time.Date(2026, 2, 15, 0, 0, 0, 0, time.UTC)
	isLeap2 := lxtime.IsLeapYear(t2)
	// isLeap2: false
	_ = isLeap2
}
