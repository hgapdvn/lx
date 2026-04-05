package lxtime_test

import (
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/lxtime"
)

func TestAddBusinessDays_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		days     int
		expected time.Time
	}{
		{
			name:     "add zero days from Monday",
			input:    time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC), // Monday
			days:     0,
			expected: time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC), // Monday
		},
		{
			name:     "add 1 business day from Monday",
			input:    time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC), // Monday
			days:     1,
			expected: time.Date(2026, 4, 7, 15, 30, 0, 0, time.UTC), // Tuesday
		},
		{
			name:     "add 5 business days from Monday",
			input:    time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC), // Monday
			days:     5,
			expected: time.Date(2026, 4, 13, 15, 30, 0, 0, time.UTC), // Monday of next week
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.AddBusinessDays(tt.input, tt.days)
			if !result.Equal(tt.expected) {
				t.Errorf("AddBusinessDays(%v, %d) = %v, want %v", tt.input, tt.days, result, tt.expected)
			}
		})
	}
}

func TestAddBusinessDays_SkipWeekend(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		days     int
		expected time.Time
	}{
		{
			name:     "Friday + 1 day = Monday",
			input:    time.Date(2026, 4, 3, 15, 30, 0, 0, time.UTC), // Friday
			days:     1,
			expected: time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC), // Monday
		},
		{
			name:     "Friday + 3 days = Wednesday",
			input:    time.Date(2026, 4, 3, 15, 30, 0, 0, time.UTC), // Friday
			days:     3,
			expected: time.Date(2026, 4, 8, 15, 30, 0, 0, time.UTC), // Wednesday
		},
		{
			name:     "Thursday + 1 day = Friday",
			input:    time.Date(2026, 4, 2, 15, 30, 0, 0, time.UTC), // Thursday
			days:     1,
			expected: time.Date(2026, 4, 3, 15, 30, 0, 0, time.UTC), // Friday
		},
		{
			name:     "Thursday + 2 days = Monday",
			input:    time.Date(2026, 4, 2, 15, 30, 0, 0, time.UTC), // Thursday
			days:     2,
			expected: time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC), // Monday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.AddBusinessDays(tt.input, tt.days)
			if !result.Equal(tt.expected) {
				t.Errorf("AddBusinessDays(%v, %d) = %v, want %v", tt.input, tt.days, result, tt.expected)
			}
		})
	}
}

func TestAddBusinessDays_Negative(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		days     int
		expected time.Time
	}{
		{
			name:     "subtract 1 business day from Monday",
			input:    time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC), // Monday
			days:     -1,
			expected: time.Date(2026, 4, 3, 15, 30, 0, 0, time.UTC), // Friday
		},
		{
			name:     "subtract 5 business days from Monday",
			input:    time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC), // Monday
			days:     -5,
			expected: time.Date(2026, 3, 30, 15, 30, 0, 0, time.UTC), // Monday of previous week
		},
		{
			name:     "Monday - 2 days = Thursday",
			input:    time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC), // Monday
			days:     -2,
			expected: time.Date(2026, 4, 2, 15, 30, 0, 0, time.UTC), // Thursday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.AddBusinessDays(tt.input, tt.days)
			if !result.Equal(tt.expected) {
				t.Errorf("AddBusinessDays(%v, %d) = %v, want %v", tt.input, tt.days, result, tt.expected)
			}
		})
	}
}

func TestAddBusinessDays_NegativeSkipWeekend(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		days     int
		expected time.Time
	}{
		{
			name:     "Monday - 1 day = Friday",
			input:    time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC), // Monday
			days:     -1,
			expected: time.Date(2026, 4, 3, 15, 30, 0, 0, time.UTC), // Friday
		},
		{
			name:     "Tuesday - 3 days = Thursday",
			input:    time.Date(2026, 4, 7, 15, 30, 0, 0, time.UTC), // Tuesday
			days:     -3,
			expected: time.Date(2026, 4, 2, 15, 30, 0, 0, time.UTC), // Thursday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.AddBusinessDays(tt.input, tt.days)
			if !result.Equal(tt.expected) {
				t.Errorf("AddBusinessDays(%v, %d) = %v, want %v", tt.input, tt.days, result, tt.expected)
			}
		})
	}
}

func TestAddBusinessDays_StartingOnWeekend(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		days     int
		expected time.Time
	}{
		{
			name:     "Saturday + 0 days = Monday",
			input:    time.Date(2026, 4, 4, 15, 30, 0, 0, time.UTC), // Saturday
			days:     0,
			expected: time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC), // Monday
		},
		{
			name:     "Sunday + 0 days = Monday",
			input:    time.Date(2026, 4, 5, 15, 30, 0, 0, time.UTC), // Sunday
			days:     0,
			expected: time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC), // Monday
		},
		{
			name:     "Saturday + 1 day = Monday (after moving to Monday)",
			input:    time.Date(2026, 4, 4, 15, 30, 0, 0, time.UTC), // Saturday
			days:     1,
			expected: time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC), // Monday
		},
		{
			name:     "Sunday + 1 day = Monday (after moving to Monday)",
			input:    time.Date(2026, 4, 5, 15, 30, 0, 0, time.UTC), // Sunday
			days:     1,
			expected: time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC), // Monday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.AddBusinessDays(tt.input, tt.days)
			if !result.Equal(tt.expected) {
				t.Errorf("AddBusinessDays(%v, %d) = %v, want %v", tt.input, tt.days, result, tt.expected)
			}
		})
	}
}

func TestAddBusinessDays_PreservesTime(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "preserves hour",
			check: func() bool {
				input := time.Date(2026, 4, 6, 15, 30, 45, 123456789, time.UTC)
				result := lxtime.AddBusinessDays(input, 1)
				return result.Hour() == 15
			},
		},
		{
			name: "preserves minute",
			check: func() bool {
				input := time.Date(2026, 4, 6, 15, 30, 45, 123456789, time.UTC)
				result := lxtime.AddBusinessDays(input, 1)
				return result.Minute() == 30
			},
		},
		{
			name: "preserves second",
			check: func() bool {
				input := time.Date(2026, 4, 6, 15, 30, 45, 123456789, time.UTC)
				result := lxtime.AddBusinessDays(input, 1)
				return result.Second() == 45
			},
		},
		{
			name: "preserves nanosecond",
			check: func() bool {
				input := time.Date(2026, 4, 6, 15, 30, 45, 123456789, time.UTC)
				result := lxtime.AddBusinessDays(input, 1)
				return result.Nanosecond() == 123456789
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("AddBusinessDays() preserve time check failed")
			}
		})
	}
}

func TestAddBusinessDays_PreservesTimezone(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "preserves UTC timezone",
			check: func() bool {
				input := time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC)
				result := lxtime.AddBusinessDays(input, 1)
				return result.Location() == time.UTC
			},
		},
		{
			name: "preserves EST timezone",
			check: func() bool {
				est, _ := time.LoadLocation("America/New_York")
				input := time.Date(2026, 4, 6, 15, 30, 0, 0, est)
				result := lxtime.AddBusinessDays(input, 1)
				return result.Location() == est
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("AddBusinessDays() preserve timezone check failed")
			}
		})
	}
}

func TestAddBusinessDays_MonthBoundary(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		days     int
		expected time.Time
	}{
		{
			name:     "end of March to April",
			input:    time.Date(2026, 3, 31, 15, 30, 0, 0, time.UTC), // Tuesday, March 31
			days:     1,
			expected: time.Date(2026, 4, 1, 15, 30, 0, 0, time.UTC), // Wednesday, April 1
		},
		{
			name:     "Thursday March 26 + 5 days = Wednesday April 1",
			input:    time.Date(2026, 3, 26, 15, 30, 0, 0, time.UTC), // Thursday, March 26
			days:     5,
			expected: time.Date(2026, 4, 2, 15, 30, 0, 0, time.UTC), // Thursday, April 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.AddBusinessDays(tt.input, tt.days)
			if !result.Equal(tt.expected) {
				t.Errorf("AddBusinessDays(%v, %d) = %v, want %v", tt.input, tt.days, result, tt.expected)
			}
		})
	}
}

func TestAddBusinessDays_LargeValues(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "add 100 business days",
			check: func() bool {
				input := time.Date(2026, 4, 6, 0, 0, 0, 0, time.UTC) // Monday
				result := lxtime.AddBusinessDays(input, 100)
				// Should be a weekday
				return result.Weekday() != time.Saturday && result.Weekday() != time.Sunday
			},
		},
		{
			name: "subtract 100 business days",
			check: func() bool {
				input := time.Date(2026, 4, 6, 0, 0, 0, 0, time.UTC) // Monday
				result := lxtime.AddBusinessDays(input, -100)
				// Should be a weekday
				return result.Weekday() != time.Saturday && result.Weekday() != time.Sunday
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("AddBusinessDays() large values check failed")
			}
		})
	}
}

func TestAddBusinessDays_Consistency(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "add then subtract returns same date",
			check: func() bool {
				input := time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC)
				added := lxtime.AddBusinessDays(input, 5)
				subtracted := lxtime.AddBusinessDays(added, -5)
				return input.Equal(subtracted)
			},
		},
		{
			name: "result is always a weekday",
			check: func() bool {
				input := time.Date(2026, 4, 3, 15, 30, 0, 0, time.UTC) // Friday
				for i := -10; i <= 10; i++ {
					result := lxtime.AddBusinessDays(input, i)
					weekday := result.Weekday()
					if weekday == time.Saturday || weekday == time.Sunday {
						return false
					}
				}
				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("AddBusinessDays() consistency check failed")
			}
		})
	}
}
