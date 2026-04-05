package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/time"
)

func TestDaysInYear_LeapYears(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "2024 is a leap year",
			date:     time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 366,
		},
		{
			name:     "2020 is a leap year",
			date:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 366,
		},
		{
			name:     "2000 is a leap year (divisible by 400)",
			date:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 366,
		},
		{
			name:     "2400 is a leap year (divisible by 400)",
			date:     time.Date(2400, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 366,
		},
		{
			name:     "1904 is a leap year",
			date:     time.Date(1904, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 366,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.TotalDaysInYear(tt.date)
			if result != tt.expected {
				t.Errorf("TotalDaysInYear(%v) = %d, want %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestDaysInYear_NonLeapYears(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "2026 is not a leap year",
			date:     time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
		{
			name:     "2025 is not a leap year",
			date:     time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
		{
			name:     "2023 is not a leap year",
			date:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
		{
			name:     "1900 is not a leap year (divisible by 100 but not 400)",
			date:     time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
		{
			name:     "2100 is not a leap year (divisible by 100 but not 400)",
			date:     time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
		{
			name:     "2200 is not a leap year (divisible by 100 but not 400)",
			date:     time.Date(2200, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.TotalDaysInYear(tt.date)
			if result != tt.expected {
				t.Errorf("TotalDaysInYear(%v) = %d, want %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestDaysInYear_DifferentDatesInSameYear(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "Jan 1, 2024 (leap year)",
			date:     time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 366,
		},
		{
			name:     "Feb 29, 2024 (leap day)",
			date:     time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC),
			expected: 366,
		},
		{
			name:     "Dec 31, 2024 (leap year)",
			date:     time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
			expected: 366,
		},
		{
			name:     "Jan 1, 2026 (non-leap year)",
			date:     time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
		{
			name:     "Jun 15, 2026 (non-leap year)",
			date:     time.Date(2026, 6, 15, 12, 30, 45, 0, time.UTC),
			expected: 365,
		},
		{
			name:     "Dec 31, 2026 (non-leap year)",
			date:     time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
			expected: 365,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.TotalDaysInYear(tt.date)
			if result != tt.expected {
				t.Errorf("TotalDaysInYear(%v) = %d, want %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestDaysInYear_TimeComponentIgnored(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "2024 midnight",
			date:     time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 366,
		},
		{
			name:     "2024 noon",
			date:     time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC),
			expected: 366,
		},
		{
			name:     "2024 end of day",
			date:     time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.UTC),
			expected: 366,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.TotalDaysInYear(tt.date)
			if result != tt.expected {
				t.Errorf("TotalDaysInYear(%v) = %d, want %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestDaysInYear_Consistency(t *testing.T) {
	t.Run("same_year_same_result", func(t *testing.T) {
		d1 := lxtime.TotalDaysInYear(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))
		d2 := lxtime.TotalDaysInYear(time.Date(2024, 6, 15, 12, 30, 0, 0, time.UTC))
		d3 := lxtime.TotalDaysInYear(time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC))

		if d1 != d2 || d2 != d3 {
			t.Errorf("TotalDaysInYear should return same value for all dates in same year: %d, %d, %d", d1, d2, d3)
		}
	})

	t.Run("different_years_different_result_leap_vs_non_leap", func(t *testing.T) {
		d2024 := lxtime.TotalDaysInYear(time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC))
		d2026 := lxtime.TotalDaysInYear(time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC))

		if d2024 == d2026 {
			t.Errorf("TotalDaysInYear should differ between leap and non-leap years: %d vs %d", d2024, d2026)
		}

		if d2024 != 366 || d2026 != 365 {
			t.Errorf("Expected 366 for 2024 and 365 for 2026, got %d and %d", d2024, d2026)
		}
	})
}

func TestDaysInYear_HistoricalYears(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "1000 AD (leap year)",
			date:     time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 365, // 1000 is not leap (divisible by 100 but not 400)
		},
		{
			name:     "1600 AD (leap year)",
			date:     time.Date(1600, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 366, // 1600 is leap (divisible by 400)
		},
		{
			name:     "1776 AD (leap year)",
			date:     time.Date(1776, 7, 4, 0, 0, 0, 0, time.UTC),
			expected: 366, // 1776 is leap (divisible by 4, not by 100)
		},
		{
			name:     "1800 AD (non-leap year)",
			date:     time.Date(1800, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 365, // 1800 is not leap (divisible by 100 but not 400)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.TotalDaysInYear(tt.date)
			if result != tt.expected {
				t.Errorf("TotalDaysInYear(%v) = %d, want %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestDaysInYear_DifferentTimezones(t *testing.T) {
	est, _ := time.LoadLocation("America/New_York")
	pst, _ := time.LoadLocation("America/Los_Angeles")
	jst, _ := time.LoadLocation("Asia/Tokyo")

	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "2024 leap year in EST",
			date:     time.Date(2024, 6, 15, 0, 0, 0, 0, est),
			expected: 366,
		},
		{
			name:     "2024 leap year in PST",
			date:     time.Date(2024, 6, 15, 0, 0, 0, 0, pst),
			expected: 366,
		},
		{
			name:     "2024 leap year in JST",
			date:     time.Date(2024, 6, 15, 0, 0, 0, 0, jst),
			expected: 366,
		},
		{
			name:     "2026 non-leap year in EST",
			date:     time.Date(2026, 6, 15, 0, 0, 0, 0, est),
			expected: 365,
		},
		{
			name:     "2026 non-leap year in PST",
			date:     time.Date(2026, 6, 15, 0, 0, 0, 0, pst),
			expected: 365,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.TotalDaysInYear(tt.date)
			if result != tt.expected {
				t.Errorf("TotalDaysInYear(%v) = %d, want %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestDaysInYear_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "year 1 AD",
			date:     time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
		{
			name:     "year 4 AD (leap year)",
			date:     time.Date(4, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 366,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.TotalDaysInYear(tt.date)
			if result != tt.expected {
				t.Errorf("TotalDaysInYear(%v) = %d, want %d", tt.date, result, tt.expected)
			}
		})
	}
}

func ExampleTotalDaysInYear() {
	t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	days := lxtime.TotalDaysInYear(t)
	// days: 366
	_ = days

	t2 := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	days2 := lxtime.TotalDaysInYear(t2)
	// days2: 365
	_ = days2
}
