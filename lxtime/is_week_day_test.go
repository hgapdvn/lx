package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxtime"
)

func TestIsWeekDay_BasicWeekdays(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "Monday is weekday",
			time:     time.Date(2026, 4, 6, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Tuesday is weekday",
			time:     time.Date(2026, 4, 7, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Wednesday is weekday",
			time:     time.Date(2026, 4, 8, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Thursday is weekday",
			time:     time.Date(2026, 4, 9, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Friday is weekday",
			time:     time.Date(2026, 4, 10, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsWeekDay(tt.time)
			if result != tt.expected {
				t.Errorf("IsWeekDay() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWeekDay_Weekends(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "Saturday is not weekday",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Sunday is not weekday",
			time:     time.Date(2026, 4, 5, 12, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsWeekDay(tt.time)
			if result != tt.expected {
				t.Errorf("IsWeekDay() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWeekDay_DifferentTimes(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "Monday midnight",
			time:     time.Date(2026, 4, 6, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Monday noon",
			time:     time.Date(2026, 4, 6, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Monday end of day",
			time:     time.Date(2026, 4, 6, 23, 59, 59, 999999999, time.UTC),
			expected: true,
		},
		{
			name:     "Saturday midnight",
			time:     time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Saturday noon",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Sunday end of day",
			time:     time.Date(2026, 4, 5, 23, 59, 59, 999999999, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsWeekDay(tt.time)
			if result != tt.expected {
				t.Errorf("IsWeekDay() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWeekDay_MultipleWeeks(t *testing.T) {
	tests := []struct {
		name     string
		date     int
		month    time.Month
		expected bool
	}{
		// April 2026 starts on Wednesday
		{
			name:     "April 1, 2026 (Wednesday)",
			date:     1,
			month:    time.April,
			expected: true,
		},
		{
			name:     "April 4, 2026 (Saturday)",
			date:     4,
			month:    time.April,
			expected: false,
		},
		{
			name:     "April 6, 2026 (Monday)",
			date:     6,
			month:    time.April,
			expected: true,
		},
		{
			name:     "April 11, 2026 (Saturday)",
			date:     11,
			month:    time.April,
			expected: false,
		},
		{
			name:     "April 13, 2026 (Monday)",
			date:     13,
			month:    time.April,
			expected: true,
		},
		{
			name:     "April 17, 2026 (Friday)",
			date:     17,
			month:    time.April,
			expected: true,
		},
		{
			name:     "April 18, 2026 (Saturday)",
			date:     18,
			month:    time.April,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTime := time.Date(2026, tt.month, tt.date, 12, 0, 0, 0, time.UTC)
			result := lxtime.IsWeekDay(testTime)
			if result != tt.expected {
				t.Errorf("IsWeekDay() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWeekDay_DifferentYears(t *testing.T) {
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
			expected: true,
		},
		{
			name:     "January 6, 2024 (Saturday)",
			year:     2024,
			month:    time.January,
			day:      6,
			expected: false,
		},
		{
			name:     "December 25, 2025 (Thursday)",
			year:     2025,
			month:    time.December,
			day:      25,
			expected: true,
		},
		{
			name:     "December 27, 2025 (Saturday)",
			year:     2025,
			month:    time.December,
			day:      27,
			expected: false,
		},
		{
			name:     "January 1, 2030 (Friday)",
			year:     2030,
			month:    time.January,
			day:      1,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTime := time.Date(tt.year, tt.month, tt.day, 12, 0, 0, 0, time.UTC)
			result := lxtime.IsWeekDay(testTime)
			if result != tt.expected {
				t.Errorf("IsWeekDay() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWeekDay_DifferentTimezones(t *testing.T) {
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")
	pst, _ := time.LoadLocation("America/Los_Angeles")

	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "Monday in UTC",
			time:     time.Date(2026, 4, 6, 12, 0, 0, 0, utc),
			expected: true,
		},
		{
			name:     "Monday in EST",
			time:     time.Date(2026, 4, 6, 12, 0, 0, 0, est),
			expected: true,
		},
		{
			name:     "Monday in PST",
			time:     time.Date(2026, 4, 6, 12, 0, 0, 0, pst),
			expected: true,
		},
		{
			name:     "Saturday in UTC",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 0, utc),
			expected: false,
		},
		{
			name:     "Saturday in EST",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 0, est),
			expected: false,
		},
		{
			name:     "Saturday in PST",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 0, pst),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsWeekDay(tt.time)
			if result != tt.expected {
				t.Errorf("IsWeekDay() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWeekDay_NanosecondPrecision(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "Monday with nanoseconds",
			time:     time.Date(2026, 4, 6, 12, 0, 0, 123456789, time.UTC),
			expected: true,
		},
		{
			name:     "Friday with max nanoseconds",
			time:     time.Date(2026, 4, 10, 23, 59, 59, 999999999, time.UTC),
			expected: true,
		},
		{
			name:     "Saturday with nanoseconds",
			time:     time.Date(2026, 4, 4, 12, 0, 0, 123456789, time.UTC),
			expected: false,
		},
		{
			name:     "Sunday with max nanoseconds",
			time:     time.Date(2026, 4, 5, 23, 59, 59, 999999999, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsWeekDay(tt.time)
			if result != tt.expected {
				t.Errorf("IsWeekDay() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWeekDay_AllDaysOfWeek(t *testing.T) {
	// April 2026: week starting Monday April 6
	days := []struct {
		name      string
		dayOfWeek int
		date      int
		month     time.Month
		expected  bool
	}{
		{
			name:      "Sunday",
			dayOfWeek: 0,
			date:      5,
			month:     time.April,
			expected:  false,
		},
		{
			name:      "Monday",
			dayOfWeek: 1,
			date:      6,
			month:     time.April,
			expected:  true,
		},
		{
			name:      "Tuesday",
			dayOfWeek: 2,
			date:      7,
			month:     time.April,
			expected:  true,
		},
		{
			name:      "Wednesday",
			dayOfWeek: 3,
			date:      8,
			month:     time.April,
			expected:  true,
		},
		{
			name:      "Thursday",
			dayOfWeek: 4,
			date:      9,
			month:     time.April,
			expected:  true,
		},
		{
			name:      "Friday",
			dayOfWeek: 5,
			date:      10,
			month:     time.April,
			expected:  true,
		},
		{
			name:      "Saturday",
			dayOfWeek: 6,
			date:      4,
			month:     time.April,
			expected:  false,
		},
	}

	for _, day := range days {
		t.Run(day.name, func(t *testing.T) {
			testTime := time.Date(2026, day.month, day.date, 12, 0, 0, 0, time.UTC)
			result := lxtime.IsWeekDay(testTime)
			if result != day.expected {
				t.Errorf("IsWeekDay() for %s = %v, want %v", day.name, result, day.expected)
			}
		})
	}
}

func TestIsWeekDay_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "Year boundary - Friday to Saturday",
			time:     time.Date(2025, 12, 26, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Year boundary - Saturday",
			time:     time.Date(2025, 12, 27, 12, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Leap year February Monday",
			time:     time.Date(2024, 2, 5, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Leap year February Saturday",
			time:     time.Date(2024, 2, 3, 12, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "First day of year - Sunday",
			time:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "First day of year - Monday",
			time:     time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsWeekDay(tt.time)
			if result != tt.expected {
				t.Errorf("IsWeekDay() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWeekDay_Now(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "current day has consistent result",
			check: func() bool {
				now := time.Now()
				result1 := lxtime.IsWeekDay(now)
				result2 := lxtime.IsWeekDay(now)
				return result1 == result2
			},
		},
		{
			name: "weekday is consistent",
			check: func() bool {
				now := time.Now()
				result := lxtime.IsWeekDay(now)
				weekday := now.Weekday()
				expected := weekday != time.Saturday && weekday != time.Sunday
				return result == expected
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsWeekDay() check failed")
			}
		})
	}
}

func TestIsWeekDay_Consistency(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "same day in different times returns same result",
			check: func() bool {
				base := time.Date(2026, 4, 6, 0, 0, 0, 0, time.UTC)
				result1 := lxtime.IsWeekDay(base)

				different := time.Date(2026, 4, 6, 23, 59, 59, 999999999, time.UTC)
				result2 := lxtime.IsWeekDay(different)

				return result1 == result2
			},
		},
		{
			name: "monday always returns true",
			check: func() bool {
				for i := 0; i < 52; i++ {
					monday := time.Date(2026, 4, 6, 12, 0, 0, 0, time.UTC).AddDate(0, 0, i*7)
					if monday.Weekday() == time.Monday {
						if !lxtime.IsWeekDay(monday) {
							return false
						}
					}
				}
				return true
			},
		},
		{
			name: "saturday always returns false",
			check: func() bool {
				for i := 0; i < 52; i++ {
					saturday := time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC).AddDate(0, 0, i*7)
					if saturday.Weekday() == time.Saturday {
						if lxtime.IsWeekDay(saturday) {
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
				t.Errorf("IsWeekDay() consistency check failed")
			}
		})
	}
}
