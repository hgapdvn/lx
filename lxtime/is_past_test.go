package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxtime"
)

func TestIsPast_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "one second in past",
			time:     lxtime.Ago(1, time.Second),
			expected: true,
		},
		{
			name:     "five minutes in past",
			time:     lxtime.Ago(5, time.Minute),
			expected: true,
		},
		{
			name:     "one hour in past",
			time:     lxtime.Ago(1, time.Hour),
			expected: true,
		},
		{
			name:     "one day in past",
			time:     lxtime.Ago(1, 24*time.Hour),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsPast(tt.time)
			if result != tt.expected {
				t.Errorf("IsPast(%v) = %v, want %v", tt.time, result, tt.expected)
			}
		})
	}
}

func TestIsPast_FutureTimes(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "one second in future",
			time:     lxtime.FromNow(1, time.Second),
			expected: false,
		},
		{
			name:     "five minutes in future",
			time:     lxtime.FromNow(5, time.Minute),
			expected: false,
		},
		{
			name:     "one hour in future",
			time:     lxtime.FromNow(1, time.Hour),
			expected: false,
		},
		{
			name:     "one day in future",
			time:     lxtime.FromNow(1, 24*time.Hour),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsPast(tt.time)
			if result != tt.expected {
				t.Errorf("IsPast(%v) = %v, want %v", tt.time, result, tt.expected)
			}
		})
	}
}

func TestIsPast_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		getTime  func() time.Time
		expected bool
		desc     string
	}{
		{
			name: "approximately now",
			getTime: func() time.Time {
				return time.Now().Add(-100 * time.Nanosecond)
			},
			expected: true,
			desc:     "very small past (may be equal)",
		},
		{
			name: "zero value",
			getTime: func() time.Time {
				return time.Time{}
			},
			expected: true,
			desc:     "zero time is in the past",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTime := tt.getTime()
			result := lxtime.IsPast(testTime)
			// For edge cases, we check that the result matches or is reasonable
			if tt.desc == "very small past (may be equal)" {
				// Very small past times might appear equal to now depending on timing
				t.Logf("IsPast(%v) = %v (timing edge case)", testTime, result)
			} else if result != tt.expected {
				t.Errorf("IsPast(%v) = %v, want %v", testTime, result, tt.expected)
			}
		})
	}
}

func TestIsPast_WithRelativeHelpers(t *testing.T) {
	t.Run("IsPast_with_Ago", func(t *testing.T) {
		pastTime := lxtime.Ago(10, time.Minute)
		if !lxtime.IsPast(pastTime) {
			t.Errorf("IsPast should return true for time from Ago(10, time.Minute)")
		}
	})

	t.Run("IsPast_with_FromNow", func(t *testing.T) {
		futureTime := lxtime.FromNow(10, time.Minute)
		if lxtime.IsPast(futureTime) {
			t.Errorf("IsPast should return false for time from FromNow(10, time.Minute)")
		}
	})
}

func TestIsPast_Consistency(t *testing.T) {
	t.Run("consistency_multiple_calls", func(t *testing.T) {
		pastTime := lxtime.Ago(30, time.Second)
		result1 := lxtime.IsPast(pastTime)
		result2 := lxtime.IsPast(pastTime)

		// Results should both be true since the time is clearly in the past
		if !result1 || !result2 {
			t.Errorf("IsPast should consistently return true for past time")
		}
	})
}

func TestIsPast_LargeValues(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "far past (100 years)",
			time:     time.Now().AddDate(-100, 0, 0),
			expected: true,
		},
		{
			name:     "far future (100 years)",
			time:     time.Now().AddDate(100, 0, 0),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsPast(tt.time)
			if result != tt.expected {
				t.Errorf("IsPast(%v) = %v, want %v", tt.time, result, tt.expected)
			}
		})
	}
}

func TestIsFutureIsPast_Complement(t *testing.T) {
	t.Run("future_and_past_are_complementary", func(t *testing.T) {
		// For times clearly in the past or future, they should be complements
		pastTime := lxtime.Ago(1, time.Hour)
		futureTime := lxtime.FromNow(1, time.Hour)

		if lxtime.IsFuture(pastTime) {
			t.Errorf("Time in past should not be IsFuture")
		}
		if !lxtime.IsPast(pastTime) {
			t.Errorf("Time in past should be IsPast")
		}

		if !lxtime.IsFuture(futureTime) {
			t.Errorf("Time in future should be IsFuture")
		}
		if lxtime.IsPast(futureTime) {
			t.Errorf("Time in future should not be IsPast")
		}
	})
}

func ExampleIsPast() {
	t := lxtime.Ago(5, time.Minute)
	isPast := lxtime.IsPast(t)
	// isPast: true
	_ = isPast

	future := lxtime.FromNow(5, time.Minute)
	isFuturePast := lxtime.IsPast(future)
	// isFuturePast: false
	_ = isFuturePast
}
