package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxtime"
)

func TestIsBetween_TimeInRange(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		start    time.Time
		end      time.Time
		expected bool
	}{
		{
			name:     "time exactly at start",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 4, 13, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "time exactly at end",
			time:     time.Date(2026, 4, 4, 13, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 4, 13, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "time in middle of range",
			time:     time.Date(2026, 4, 4, 12, 30, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 4, 13, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "time before range",
			time:     time.Date(2026, 4, 4, 11, 59, 59, 0, time.UTC),
			start:    time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 4, 13, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "time after range",
			time:     time.Date(2026, 4, 4, 13, 0, 0, 1, time.UTC),
			start:    time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 4, 13, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "start and end equal",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "start and end equal, time not at point",
			time:     time.Date(2026, 4, 4, 11, 59, 59, 0, time.UTC),
			start:    time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsBetween(tt.time, tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("IsBetween() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsBetween_DayRanges(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		start    time.Time
		end      time.Time
		expected bool
	}{
		{
			name:     "time at start of day range",
			time:     time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 5, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "time at end of day range",
			time:     time.Date(2026, 4, 5, 0, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 5, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "time in middle of day range",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 5, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "time before day range",
			time:     time.Date(2026, 4, 3, 23, 59, 59, 0, time.UTC),
			start:    time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 5, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "time after day range",
			time:     time.Date(2026, 4, 5, 0, 0, 0, 1, time.UTC),
			start:    time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 5, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsBetween(tt.time, tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("IsBetween() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsBetween_MonthRanges(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		start    time.Time
		end      time.Time
		expected bool
	}{
		{
			name:     "time at start of month range",
			time:     time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 5, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "time at end of month range",
			time:     time.Date(2026, 5, 1, 0, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 5, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "time in middle of month range",
			time:     time.Date(2026, 4, 15, 12, 30, 45, 0, time.UTC),
			start:    time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 5, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "time before month range",
			time:     time.Date(2026, 3, 31, 23, 59, 59, 0, time.UTC),
			start:    time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 5, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "time after month range",
			time:     time.Date(2026, 5, 1, 0, 0, 0, 1, time.UTC),
			start:    time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 5, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsBetween(tt.time, tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("IsBetween() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsBetween_YearRanges(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		start    time.Time
		end      time.Time
		expected bool
	}{
		{
			name:     "time at start of year range",
			time:     time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "time at end of year range",
			time:     time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "time in middle of year range",
			time:     time.Date(2026, 6, 15, 12, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "time before year range",
			time:     time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "time after year range",
			time:     time.Date(2027, 1, 1, 0, 0, 0, 1, time.UTC),
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsBetween(tt.time, tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("IsBetween() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsBetween_NanosecondPrecision(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		start    time.Time
		end      time.Time
		expected bool
	}{
		{
			name:     "exact nanosecond at start",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 1000, time.UTC),
			start:    time.Date(2026, 4, 4, 12, 0, 0, 1000, time.UTC),
			end:      time.Date(2026, 4, 4, 13, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "exact nanosecond at end",
			time:     time.Date(2026, 4, 4, 13, 0, 0, 999, time.UTC),
			start:    time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 4, 13, 0, 0, 999, time.UTC),
			expected: true,
		},
		{
			name:     "one nanosecond before start",
			time:     time.Date(2026, 4, 4, 11, 59, 59, 999999999, time.UTC),
			start:    time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 4, 13, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "one nanosecond after end",
			time:     time.Date(2026, 4, 4, 13, 0, 0, 1, time.UTC),
			start:    time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 4, 13, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsBetween(tt.time, tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("IsBetween() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsBetween_DifferentTimezones(t *testing.T) {
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")

	tests := []struct {
		name     string
		time     time.Time
		start    time.Time
		end      time.Time
		expected bool
	}{
		{
			name:     "same time different timezones at start",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 0, est),
			start:    time.Date(2026, 4, 4, 16, 0, 0, 0, utc),
			end:      time.Date(2026, 4, 4, 17, 0, 0, 0, utc),
			expected: true,
		},
		{
			name:     "same time different timezones at end",
			time:     time.Date(2026, 4, 4, 13, 0, 0, 0, est),
			start:    time.Date(2026, 4, 4, 16, 0, 0, 0, utc),
			end:      time.Date(2026, 4, 4, 17, 0, 0, 0, utc),
			expected: true,
		},
		{
			name:     "time in middle with different timezone",
			time:     time.Date(2026, 4, 4, 12, 30, 0, 0, est),
			start:    time.Date(2026, 4, 4, 16, 0, 0, 0, utc),
			end:      time.Date(2026, 4, 4, 17, 0, 0, 0, utc),
			expected: true,
		},
		{
			name:     "time before range different timezone",
			time:     time.Date(2026, 4, 4, 11, 59, 59, 0, est),
			start:    time.Date(2026, 4, 4, 16, 0, 0, 0, utc),
			end:      time.Date(2026, 4, 4, 17, 0, 0, 0, utc),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsBetween(tt.time, tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("IsBetween() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsBetween_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		start    time.Time
		end      time.Time
		expected bool
	}{
		{
			name:     "start after end",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 4, 13, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 4, 11, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "all times identical",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "very small range one nanosecond",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 500, time.UTC),
			start:    time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 4, 12, 0, 0, 1000, time.UTC),
			expected: true,
		},
		{
			name:     "very small range just outside",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 1001, time.UTC),
			start:    time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 4, 12, 0, 0, 1000, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsBetween(tt.time, tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("IsBetween() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsBetween_LargeTimeRanges(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		start    time.Time
		end      time.Time
		expected bool
	}{
		{
			name:     "time in decade range",
			time:     time.Date(2022, 6, 15, 0, 0, 0, 0, time.UTC),
			start:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "time at decade start",
			time:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			start:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "time at decade end",
			time:     time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC),
			start:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "time before decade range",
			time:     time.Date(2019, 12, 31, 23, 59, 59, 0, time.UTC),
			start:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "time after decade range",
			time:     time.Date(2030, 1, 1, 0, 0, 0, 1, time.UTC),
			start:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsBetween(tt.time, tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("IsBetween() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsBetween_Now(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "now is within past and future",
			check: func() bool {
				now := time.Now()
				past := now.Add(-1 * time.Hour)
				future := now.Add(1 * time.Hour)
				return lxtime.IsBetween(now, past, future)
			},
		},
		{
			name: "now is not before past",
			check: func() bool {
				now := time.Now()
				past := now.Add(1 * time.Hour)
				future := now.Add(2 * time.Hour)
				return !lxtime.IsBetween(now, past, future)
			},
		},
		{
			name: "now is equal to start",
			check: func() bool {
				now := time.Now()
				future := now.Add(1 * time.Hour)
				return lxtime.IsBetween(now, now, future)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsBetween() check failed")
			}
		})
	}
}
