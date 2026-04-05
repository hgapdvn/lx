package lxtime_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxtime"
)

func TestDaysInMonth(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "January (31 days)",
			date:     time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			expected: 31,
		},
		{
			name:     "February non-leap year (28 days)",
			date:     time.Date(2026, 2, 15, 0, 0, 0, 0, time.UTC),
			expected: 28,
		},
		{
			name:     "February leap year (29 days)",
			date:     time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC),
			expected: 29,
		},
		{
			name:     "April (30 days)",
			date:     time.Date(2026, 4, 15, 0, 0, 0, 0, time.UTC),
			expected: 30,
		},
		{
			name:     "May (31 days)",
			date:     time.Date(2026, 5, 15, 0, 0, 0, 0, time.UTC),
			expected: 31,
		},
		{
			name:     "June (30 days)",
			date:     time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC),
			expected: 30,
		},
		{
			name:     "July (31 days)",
			date:     time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
			expected: 31,
		},
		{
			name:     "August (31 days)",
			date:     time.Date(2026, 8, 15, 0, 0, 0, 0, time.UTC),
			expected: 31,
		},
		{
			name:     "September (30 days)",
			date:     time.Date(2026, 9, 15, 0, 0, 0, 0, time.UTC),
			expected: 30,
		},
		{
			name:     "October (31 days)",
			date:     time.Date(2026, 10, 15, 0, 0, 0, 0, time.UTC),
			expected: 31,
		},
		{
			name:     "November (30 days)",
			date:     time.Date(2026, 11, 15, 0, 0, 0, 0, time.UTC),
			expected: 30,
		},
		{
			name:     "December (31 days)",
			date:     time.Date(2026, 12, 15, 0, 0, 0, 0, time.UTC),
			expected: 31,
		},
		{
			name:     "First day of month",
			date:     time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
			expected: 30,
		},
		{
			name:     "Last day of month",
			date:     time.Date(2026, 4, 30, 0, 0, 0, 0, time.UTC),
			expected: 30,
		},
		{
			name:     "Leap year boundary: century year divisible by 400",
			date:     time.Date(2000, 2, 15, 0, 0, 0, 0, time.UTC),
			expected: 29,
		},
		{
			name:     "Non-leap year boundary: century year not divisible by 400",
			date:     time.Date(1900, 2, 15, 0, 0, 0, 0, time.UTC),
			expected: 28,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.TotalDaysInMonth(tt.date)
			if result != tt.expected {
				t.Errorf("TotalDaysInMonth(%v) = %d, want %d", tt.date, result, tt.expected)
			}
		})
	}
}

func ExampleTotalDaysInMonth() {
	t := time.Date(2026, 4, 15, 0, 0, 0, 0, time.UTC)
	fmt.Println(lxtime.TotalDaysInMonth(t))
	// Output: 30
}
