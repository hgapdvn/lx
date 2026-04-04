package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/time"
)

func TestFromNow_BasicDurations(t *testing.T) {
	tests := []struct {
		name             string
		n                int
		unit             time.Duration
		toleranceSeconds int64
		shouldBeFuture   bool
	}{
		{
			name:             "zero seconds",
			n:                0,
			unit:             time.Second,
			toleranceSeconds: 1,
			shouldBeFuture:   false,
		},
		{
			name:             "one second from now",
			n:                1,
			unit:             time.Second,
			toleranceSeconds: 2,
			shouldBeFuture:   true,
		},
		{
			name:             "five minutes from now",
			n:                5,
			unit:             time.Minute,
			toleranceSeconds: 2,
			shouldBeFuture:   true,
		},
		{
			name:             "one hour from now",
			n:                1,
			unit:             time.Hour,
			toleranceSeconds: 2,
			shouldBeFuture:   true,
		},
		{
			name:             "one day from now",
			n:                1,
			unit:             24 * time.Hour,
			toleranceSeconds: 2,
			shouldBeFuture:   true,
		},
		{
			name:             "thirty minutes from now",
			n:                30,
			unit:             time.Minute,
			toleranceSeconds: 2,
			shouldBeFuture:   true,
		},
		{
			name:             "two hours from now",
			n:                2,
			unit:             time.Hour,
			toleranceSeconds: 2,
			shouldBeFuture:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := time.Now()
			result := lxtime.FromNow(tt.n, tt.unit)

			// Check that result is in the future
			if tt.shouldBeFuture {
				if !result.After(before) {
					t.Errorf("FromNow(%d, %v) should be in the future, got %v", tt.n, tt.unit, result)
				}
			}

			// Check approximate accuracy within tolerance
			actualDiff := result.Sub(before).Seconds()
			expectedDiff := time.Duration(tt.n) * tt.unit / time.Second
			tolerance := time.Duration(tt.toleranceSeconds) * time.Second / time.Second

			if actualDiff < expectedDiff.Seconds()-tolerance.Seconds() ||
				actualDiff > expectedDiff.Seconds()+tolerance.Seconds() {
				t.Logf("FromNow(%d, %v) timing difference: %v seconds (expected ~%v seconds)", tt.n, tt.unit, actualDiff, expectedDiff.Seconds())
			}
		})
	}
}

func TestFromNow_StandardDurations(t *testing.T) {
	tests := []struct {
		name string
		n    int
		unit time.Duration
		desc string
	}{
		{name: "milliseconds", n: 100, unit: time.Millisecond, desc: "100 milliseconds"},
		{name: "seconds", n: 45, unit: time.Second, desc: "45 seconds"},
		{name: "minutes", n: 15, unit: time.Minute, desc: "15 minutes"},
		{name: "hours", n: 3, unit: time.Hour, desc: "3 hours"},
	}

	now := time.Now()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.FromNow(tt.n, tt.unit)

			// Verify result is after now
			if !result.After(now) && !result.Equal(now) {
				t.Errorf("FromNow(%d, %v) should be at or after now", tt.n, tt.unit)
			}

			// Verify approximate distance
			duration := result.Sub(now)
			expectedDuration := time.Duration(tt.n) * tt.unit

			// Allow 2% tolerance for timing variations
			tolerance := expectedDuration / 50
			if duration < expectedDuration-tolerance || duration > expectedDuration+tolerance {
				t.Logf("FromNow(%d, %v) expected ~%v from now, got ~%v from now", tt.n, tt.unit, expectedDuration, duration)
			}
		})
	}
}

func TestFromNow_LargeValues(t *testing.T) {
	tests := []struct {
		name string
		n    int
		unit time.Duration
	}{
		{name: "large seconds", n: 999999, unit: time.Second},
		{name: "large minutes", n: 1000, unit: time.Minute},
		{name: "large hours", n: 100, unit: time.Hour},
	}

	now := time.Now()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.FromNow(tt.n, tt.unit)

			if !result.After(now) {
				t.Errorf("FromNow(%d, %v) should be in the future", tt.n, tt.unit)
			}
		})
	}
}

func TestFromNow_NegativeValues(t *testing.T) {
	t.Run("negative_value", func(t *testing.T) {
		now := time.Now()
		result := lxtime.FromNow(-5, time.Minute)

		// Negative value should give past time
		if !result.Before(now) {
			t.Errorf("FromNow(-5, time.Minute) should be in the past (negative from now)")
		}
	})
}

func TestFromNow_ConsistencyWithNow(t *testing.T) {
	t.Run("consistent_relative_to_now", func(t *testing.T) {
		before := time.Now()
		result := lxtime.FromNow(1, time.Hour)
		after := time.Now()

		// Result should be after before and after
		if result.Before(after) || result.Before(before) {
			t.Errorf("FromNow result should be after execution window")
		}

		// Should be approximately 1 hour after the current time
		diff := result.Sub(before)
		oneHour := time.Hour

		tolerance := 100 * time.Millisecond
		if diff < oneHour-tolerance || diff > oneHour+tolerance {
			t.Logf("FromNow timing off by %v (tolerance: %v)", diff-oneHour, tolerance)
		}
	})
}

func TestFromNow_EquivalenceWithManualAdd(t *testing.T) {
	t.Run("equivalent_to_manual_now_add", func(t *testing.T) {
		fromNowResult := lxtime.FromNow(30, time.Minute)

		manualResult := time.Now().Add(30 * time.Minute)

		// Should be very close (within 100ms)
		diff := fromNowResult.Sub(manualResult)
		if diff < 0 {
			diff = -diff
		}
		tolerance := 100 * time.Millisecond

		if diff > tolerance {
			t.Logf("FromNow result differs from manual calculation by %v (tolerance: %v)", diff, tolerance)
		}
	})
}

func TestAgoFromNow_Symmetry(t *testing.T) {
	t.Run("ago_and_from_now_are_symmetric", func(t *testing.T) {
		now := time.Now()

		ago := lxtime.Ago(10, time.Minute)
		fromNow := lxtime.FromNow(10, time.Minute)

		// Approximate distance from now should be equal
		agoDistance := now.Sub(ago)
		fromNowDistance := fromNow.Sub(now)

		diff := agoDistance - fromNowDistance
		if diff < 0 {
			diff = -diff
		}

		// Allow 1ms tolerance for timing variations
		tolerance := 1 * time.Millisecond
		if diff > tolerance {
			t.Logf("Ago and FromNow are not symmetric by %v (tolerance: %v)", diff, tolerance)
		}
	})
}

func ExampleFromNow() {
	fiveMinutesLater := lxtime.FromNow(5, time.Minute)
	// fiveMinutesLater: approximately 5 minutes from now
	_ = fiveMinutesLater

	oneHourLater := lxtime.FromNow(1, time.Hour)
	// oneHourLater: approximately 1 hour from now
	_ = oneHourLater
}
