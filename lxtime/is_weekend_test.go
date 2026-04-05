package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxtime"
)

func TestIsWeekend_BasicWeekends(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "Saturday is weekend",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Sunday is weekend",
			time:     time.Date(2026, 4, 5, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsWeekend(tt.time)
			if result != tt.expected {
				t.Errorf("IsWeekend() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWeekend_Weekdays(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "Monday is not weekend",
			time:     time.Date(2026, 4, 6, 12, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Tuesday is not weekend",
			time:     time.Date(2026, 4, 7, 12, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Wednesday is not weekend",
			time:     time.Date(2026, 4, 8, 12, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Thursday is not weekend",
			time:     time.Date(2026, 4, 9, 12, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Friday is not weekend",
			time:     time.Date(2026, 4, 10, 12, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsWeekend(tt.time)
			if result != tt.expected {
				t.Errorf("IsWeekend() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWeekend_DifferentTimes(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "Saturday midnight",
			time:     time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Saturday noon",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Saturday end of day",
			time:     time.Date(2026, 4, 4, 23, 59, 59, 999999999, time.UTC),
			expected: true,
		},
		{
			name:     "Sunday midnight",
			time:     time.Date(2026, 4, 5, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Sunday noon",
			time:     time.Date(2026, 4, 5, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Sunday end of day",
			time:     time.Date(2026, 4, 5, 23, 59, 59, 999999999, time.UTC),
			expected: true,
		},
		{
			name:     "Monday midnight",
			time:     time.Date(2026, 4, 6, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Friday end of day",
			time:     time.Date(2026, 4, 10, 23, 59, 59, 999999999, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsWeekend(tt.time)
			if result != tt.expected {
				t.Errorf("IsWeekend() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWeekend_MultipleWeeks(t *testing.T) {
	tests := []struct {
		name     string
		date     int
		month    time.Month
		expected bool
	}{
		// April 2026 starts on Wednesday
		{
			name:     "April 4, 2026 (Saturday)",
			date:     4,
			month:    time.April,
			expected: true,
		},
		{
			name:     "April 5, 2026 (Sunday)",
			date:     5,
			month:    time.April,
			expected: true,
		},
		{
			name:     "April 6, 2026 (Monday)",
			date:     6,
			month:    time.April,
			expected: false,
		},
		{
			name:     "April 11, 2026 (Saturday)",
			date:     11,
			month:    time.April,
			expected: true,
		},
		{
			name:     "April 12, 2026 (Sunday)",
			date:     12,
			month:    time.April,
			expected: true,
		},
		{
			name:     "April 13, 2026 (Monday)",
			date:     13,
			month:    time.April,
			expected: false,
		},
		{
			name:     "April 18, 2026 (Saturday)",
			date:     18,
			month:    time.April,
			expected: true,
		},
		{
			name:     "April 19, 2026 (Sunday)",
			date:     19,
			month:    time.April,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTime := time.Date(2026, tt.month, tt.date, 12, 0, 0, 0, time.UTC)
			result := lxtime.IsWeekend(testTime)
			if result != tt.expected {
				t.Errorf("IsWeekend() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWeekend_DifferentYears(t *testing.T) {
	tests := []struct {
		name     string
		year     int
		month    time.Month
		day      int
		expected bool
	}{
		{
			name:     "January 4, 2024 (Friday)",
			year:     2024,
			month:    time.January,
			day:      5,
			expected: false,
		},
		{
			name:     "January 6, 2024 (Saturday)",
			year:     2024,
			month:    time.January,
			day:      6,
			expected: true,
		},
		{
			name:     "January 7, 2024 (Sunday)",
			year:     2024,
			month:    time.January,
			day:      7,
			expected: true,
		},
		{
			name:     "December 25, 2025 (Thursday)",
			year:     2025,
			month:    time.December,
			day:      25,
			expected: false,
		},
		{
			name:     "December 27, 2025 (Saturday)",
			year:     2025,
			month:    time.December,
			day:      27,
			expected: true,
		},
		{
			name:     "December 28, 2025 (Sunday)",
			year:     2025,
			month:    time.December,
			day:      28,
			expected: true,
		},
		{
			name:     "January 1, 2030 (Friday)",
			year:     2030,
			month:    time.January,
			day:      1,
			expected: false,
		},
		{
			name:     "January 5, 2030 (Saturday)",
			year:     2030,
			month:    time.January,
			day:      5,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTime := time.Date(tt.year, tt.month, tt.day, 12, 0, 0, 0, time.UTC)
			result := lxtime.IsWeekend(testTime)
			if result != tt.expected {
				t.Errorf("IsWeekend() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWeekend_DifferentTimezones(t *testing.T) {
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")
	pst, _ := time.LoadLocation("America/Los_Angeles")

	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "Saturday in UTC",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 0, utc),
			expected: true,
		},
		{
			name:     "Saturday in EST",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 0, est),
			expected: true,
		},
		{
			name:     "Saturday in PST",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 0, pst),
			expected: true,
		},
		{
			name:     "Sunday in UTC",
			time:     time.Date(2026, 4, 5, 12, 0, 0, 0, utc),
			expected: true,
		},
		{
			name:     "Sunday in EST",
			time:     time.Date(2026, 4, 5, 12, 0, 0, 0, est),
			expected: true,
		},
		{
			name:     "Monday in UTC",
			time:     time.Date(2026, 4, 6, 12, 0, 0, 0, utc),
			expected: false,
		},
		{
			name:     "Monday in EST",
			time:     time.Date(2026, 4, 6, 12, 0, 0, 0, est),
			expected: false,
		},
		{
			name:     "Friday in PST",
			time:     time.Date(2026, 4, 10, 12, 0, 0, 0, pst),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsWeekend(tt.time)
			if result != tt.expected {
				t.Errorf("IsWeekend() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWeekend_NanosecondPrecision(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "Saturday with nanoseconds",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 123456789, time.UTC),
			expected: true,
		},
		{
			name:     "Saturday with max nanoseconds",
			time:     time.Date(2026, 4, 4, 23, 59, 59, 999999999, time.UTC),
			expected: true,
		},
		{
			name:     "Sunday with nanoseconds",
			time:     time.Date(2026, 4, 5, 12, 0, 0, 123456789, time.UTC),
			expected: true,
		},
		{
			name:     "Monday with nanoseconds",
			time:     time.Date(2026, 4, 6, 12, 0, 0, 123456789, time.UTC),
			expected: false,
		},
		{
			name:     "Friday with max nanoseconds",
			time:     time.Date(2026, 4, 10, 23, 59, 59, 999999999, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsWeekend(tt.time)
			if result != tt.expected {
				t.Errorf("IsWeekend() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWeekend_AllDaysOfWeek(t *testing.T) {
	// April 2026: week starting Monday April 6
	days := []struct {
		name     string
		date     int
		month    time.Month
		expected bool
	}{
		{
			name:     "Sunday",
			date:     5,
			month:    time.April,
			expected: true,
		},
		{
			name:     "Monday",
			date:     6,
			month:    time.April,
			expected: false,
		},
		{
			name:     "Tuesday",
			date:     7,
			month:    time.April,
			expected: false,
		},
		{
			name:     "Wednesday",
			date:     8,
			month:    time.April,
			expected: false,
		},
		{
			name:     "Thursday",
			date:     9,
			month:    time.April,
			expected: false,
		},
		{
			name:     "Friday",
			date:     10,
			month:    time.April,
			expected: false,
		},
		{
			name:     "Saturday",
			date:     4,
			month:    time.April,
			expected: true,
		},
	}

	for _, day := range days {
		t.Run(day.name, func(t *testing.T) {
			testTime := time.Date(2026, day.month, day.date, 12, 0, 0, 0, time.UTC)
			result := lxtime.IsWeekend(testTime)
			if result != day.expected {
				t.Errorf("IsWeekend() for %s = %v, want %v", day.name, result, day.expected)
			}
		})
	}
}

func TestIsWeekend_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "Year boundary - Friday to Saturday",
			time:     time.Date(2025, 12, 27, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Year boundary - Sunday to Monday",
			time:     time.Date(2025, 12, 28, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Leap year February Saturday",
			time:     time.Date(2024, 2, 3, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Leap year February Sunday",
			time:     time.Date(2024, 2, 4, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "First day of year - Sunday",
			time:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "First day of year - Monday",
			time:     time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Last day of year - Saturday",
			time:     time.Date(2025, 12, 27, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Last day of year - Sunday",
			time:     time.Date(2025, 12, 28, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsWeekend(tt.time)
			if result != tt.expected {
				t.Errorf("IsWeekend() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWeekend_Now(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "current day has consistent result",
			check: func() bool {
				now := time.Now()
				result1 := lxtime.IsWeekend(now)
				result2 := lxtime.IsWeekend(now)
				return result1 == result2
			},
		},
		{
			name: "weekday is consistent",
			check: func() bool {
				now := time.Now()
				result := lxtime.IsWeekend(now)
				weekday := now.Weekday()
				expected := weekday == time.Saturday || weekday == time.Sunday
				return result == expected
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsWeekend() check failed")
			}
		})
	}
}

func TestIsWeekend_Consistency(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "same day in different times returns same result",
			check: func() bool {
				base := time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC)
				result1 := lxtime.IsWeekend(base)

				different := time.Date(2026, 4, 4, 23, 59, 59, 999999999, time.UTC)
				result2 := lxtime.IsWeekend(different)

				return result1 == result2
			},
		},
		{
			name: "saturday always returns true",
			check: func() bool {
				for i := 0; i < 52; i++ {
					saturday := time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC).AddDate(0, 0, i*7)
					if saturday.Weekday() == time.Saturday {
						if !lxtime.IsWeekend(saturday) {
							return false
						}
					}
				}
				return true
			},
		},
		{
			name: "sunday always returns true",
			check: func() bool {
				for i := 0; i < 52; i++ {
					sunday := time.Date(2026, 4, 5, 12, 0, 0, 0, time.UTC).AddDate(0, 0, i*7)
					if sunday.Weekday() == time.Sunday {
						if !lxtime.IsWeekend(sunday) {
							return false
						}
					}
				}
				return true
			},
		},
		{
			name: "monday always returns false",
			check: func() bool {
				for i := 0; i < 52; i++ {
					monday := time.Date(2026, 4, 6, 12, 0, 0, 0, time.UTC).AddDate(0, 0, i*7)
					if monday.Weekday() == time.Monday {
						if lxtime.IsWeekend(monday) {
							return false
						}
					}
				}
				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsWeekend() consistency check failed")
			}
		})
	}
}

func TestIsWeekend_ComplementToIsWeekday(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "IsWeekend is complement of IsWeekday",
			check: func() bool {
				for day := 1; day <= 30; day++ {
					testTime := time.Date(2026, 4, day, 12, 0, 0, 0, time.UTC)
					isWeekend := lxtime.IsWeekend(testTime)
					isWeekday := lxtime.IsWeekDay(testTime)
					if isWeekend == isWeekday {
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
				t.Errorf("IsWeekend() complement check failed")
			}
		})
	}
}
