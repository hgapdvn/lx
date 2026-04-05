package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxtime"
)

func TestQuarterOf_Q1(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "January 1st",
			date:     time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "January 31st",
			date:     time.Date(2026, 1, 31, 23, 59, 59, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "February 14th",
			date:     time.Date(2026, 2, 14, 12, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "February 28th",
			date:     time.Date(2026, 2, 28, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "March 1st",
			date:     time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "March 31st",
			date:     time.Date(2026, 3, 31, 23, 59, 59, 999999999, time.UTC),
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.QuarterOf(tt.date)
			if result != tt.expected {
				t.Errorf("QuarterOf(%v) = %d, want %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestQuarterOf_Q2(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "April 1st",
			date:     time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
			expected: 2,
		},
		{
			name:     "April 30th",
			date:     time.Date(2026, 4, 30, 23, 59, 59, 0, time.UTC),
			expected: 2,
		},
		{
			name:     "May 15th",
			date:     time.Date(2026, 5, 15, 12, 30, 45, 0, time.UTC),
			expected: 2,
		},
		{
			name:     "June 1st",
			date:     time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC),
			expected: 2,
		},
		{
			name:     "June 30th",
			date:     time.Date(2026, 6, 30, 23, 59, 59, 999999999, time.UTC),
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.QuarterOf(tt.date)
			if result != tt.expected {
				t.Errorf("QuarterOf(%v) = %d, want %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestQuarterOf_Q3(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "July 1st",
			date:     time.Date(2026, 7, 1, 0, 0, 0, 0, time.UTC),
			expected: 3,
		},
		{
			name:     "July 31st",
			date:     time.Date(2026, 7, 31, 23, 59, 59, 0, time.UTC),
			expected: 3,
		},
		{
			name:     "August 15th",
			date:     time.Date(2026, 8, 15, 12, 0, 0, 0, time.UTC),
			expected: 3,
		},
		{
			name:     "September 1st",
			date:     time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC),
			expected: 3,
		},
		{
			name:     "September 30th",
			date:     time.Date(2026, 9, 30, 23, 59, 59, 999999999, time.UTC),
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.QuarterOf(tt.date)
			if result != tt.expected {
				t.Errorf("QuarterOf(%v) = %d, want %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestQuarterOf_Q4(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "October 1st",
			date:     time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC),
			expected: 4,
		},
		{
			name:     "October 31st",
			date:     time.Date(2026, 10, 31, 23, 59, 59, 0, time.UTC),
			expected: 4,
		},
		{
			name:     "November 15th",
			date:     time.Date(2026, 11, 15, 12, 0, 0, 0, time.UTC),
			expected: 4,
		},
		{
			name:     "December 1st",
			date:     time.Date(2026, 12, 1, 0, 0, 0, 0, time.UTC),
			expected: 4,
		},
		{
			name:     "December 31st",
			date:     time.Date(2026, 12, 31, 23, 59, 59, 999999999, time.UTC),
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.QuarterOf(tt.date)
			if result != tt.expected {
				t.Errorf("QuarterOf(%v) = %d, want %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestQuarterOf_DifferentYears(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "2024 January",
			date:     time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "2024 April",
			date:     time.Date(2024, 4, 15, 0, 0, 0, 0, time.UTC),
			expected: 2,
		},
		{
			name:     "2024 July",
			date:     time.Date(2024, 7, 15, 0, 0, 0, 0, time.UTC),
			expected: 3,
		},
		{
			name:     "2024 October",
			date:     time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC),
			expected: 4,
		},
		{
			name:     "2025 February",
			date:     time.Date(2025, 2, 15, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "2025 May",
			date:     time.Date(2025, 5, 15, 0, 0, 0, 0, time.UTC),
			expected: 2,
		},
		{
			name:     "2025 August",
			date:     time.Date(2025, 8, 15, 0, 0, 0, 0, time.UTC),
			expected: 3,
		},
		{
			name:     "2025 November",
			date:     time.Date(2025, 11, 15, 0, 0, 0, 0, time.UTC),
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.QuarterOf(tt.date)
			if result != tt.expected {
				t.Errorf("QuarterOf(%v) = %d, want %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestQuarterOf_DifferentTimezones(t *testing.T) {
	est, _ := time.LoadLocation("America/New_York")
	pst, _ := time.LoadLocation("America/Los_Angeles")
	jst, _ := time.LoadLocation("Asia/Tokyo")

	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "Q1 in EST",
			date:     time.Date(2026, 2, 15, 0, 0, 0, 0, est),
			expected: 1,
		},
		{
			name:     "Q2 in PST",
			date:     time.Date(2026, 5, 15, 0, 0, 0, 0, pst),
			expected: 2,
		},
		{
			name:     "Q3 in JST",
			date:     time.Date(2026, 8, 15, 0, 0, 0, 0, jst),
			expected: 3,
		},
		{
			name:     "Q4 in EST",
			date:     time.Date(2026, 11, 15, 0, 0, 0, 0, est),
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.QuarterOf(tt.date)
			if result != tt.expected {
				t.Errorf("QuarterOf(%v) = %d, want %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestQuarterOf_TimeComponentIgnored(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "January midnight",
			date:     time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "January noon",
			date:     time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "January end of day",
			date:     time.Date(2026, 1, 15, 23, 59, 59, 999999999, time.UTC),
			expected: 1,
		},
		{
			name:     "April midnight",
			date:     time.Date(2026, 4, 15, 0, 0, 0, 0, time.UTC),
			expected: 2,
		},
		{
			name:     "April noon",
			date:     time.Date(2026, 4, 15, 12, 30, 45, 0, time.UTC),
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.QuarterOf(tt.date)
			if result != tt.expected {
				t.Errorf("QuarterOf(%v) = %d, want %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestQuarterOf_Consistency(t *testing.T) {
	// Test that same month always returns same quarter
	t.Run("same_month_same_quarter", func(t *testing.T) {
		q1 := lxtime.QuarterOf(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC))
		q2 := lxtime.QuarterOf(time.Date(2026, 1, 31, 23, 59, 59, 0, time.UTC))
		q3 := lxtime.QuarterOf(time.Date(2026, 1, 15, 12, 30, 45, 500, time.UTC))

		if q1 != q2 || q2 != q3 || q1 != 1 {
			t.Errorf("January dates should all return Q1 (1), got %d, %d, %d", q1, q2, q3)
		}
	})

	// Test that adjacent months in different quarters return different values
	t.Run("adjacent_months_different_quarters", func(t *testing.T) {
		marchQuarter := lxtime.QuarterOf(time.Date(2026, 3, 31, 23, 59, 59, 0, time.UTC))
		aprilQuarter := lxtime.QuarterOf(time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC))

		if marchQuarter != 1 {
			t.Errorf("March should be Q1 (1), got %d", marchQuarter)
		}
		if aprilQuarter != 2 {
			t.Errorf("April should be Q2 (2), got %d", aprilQuarter)
		}
		if marchQuarter == aprilQuarter {
			t.Errorf("March and April should return different quarters, both got %d", marchQuarter)
		}
	})

	// Test year boundary
	t.Run("year_boundary", func(t *testing.T) {
		decemberQuarter := lxtime.QuarterOf(time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC))
		januaryQuarter := lxtime.QuarterOf(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC))

		if decemberQuarter != 4 {
			t.Errorf("December should be Q4 (4), got %d", decemberQuarter)
		}
		if januaryQuarter != 1 {
			t.Errorf("January should be Q1 (1), got %d", januaryQuarter)
		}
	})
}

func TestQuarterOf_BoundaryMonths(t *testing.T) {
	tests := []struct {
		name     string
		month    time.Month
		expected int
	}{
		{name: "January", month: 1, expected: 1},
		{name: "February", month: 2, expected: 1},
		{name: "March", month: 3, expected: 1},
		{name: "April", month: 4, expected: 2},
		{name: "May", month: 5, expected: 2},
		{name: "June", month: 6, expected: 2},
		{name: "July", month: 7, expected: 3},
		{name: "August", month: 8, expected: 3},
		{name: "September", month: 9, expected: 3},
		{name: "October", month: 10, expected: 4},
		{name: "November", month: 11, expected: 4},
		{name: "December", month: 12, expected: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date := time.Date(2026, tt.month, 15, 0, 0, 0, 0, time.UTC)
			result := lxtime.QuarterOf(date)
			if result != tt.expected {
				t.Errorf("QuarterOf(%v) = %d, want %d", date, result, tt.expected)
			}
		})
	}
}

func ExampleQuarterOf() {
	t := time.Date(2026, 4, 15, 0, 0, 0, 0, time.UTC)
	quarter := lxtime.QuarterOf(t)
	// quarter: 2
	_ = quarter

	t2 := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	quarter2 := lxtime.QuarterOf(t2)
	// quarter2: 1
	_ = quarter2

	t3 := time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC)
	quarter3 := lxtime.QuarterOf(t3)
	// quarter3: 3
	_ = quarter3

	t4 := time.Date(2026, 10, 15, 0, 0, 0, 0, time.UTC)
	quarter4 := lxtime.QuarterOf(t4)
	// quarter4: 4
	_ = quarter4
}
