package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxtime"
)

func TestMax_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		a        time.Time
		b        time.Time
		expected time.Time
	}{
		{
			name:     "a is later",
			a:        time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			b:        time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
		},
		{
			name:     "b is later",
			a:        time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			b:        time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
		},
		{
			name:     "both equal",
			a:        time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC),
			b:        time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Max(tt.a, tt.b)
			if !result.Equal(tt.expected) {
				t.Errorf("Max(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMax_DifferentDays(t *testing.T) {
	tests := []struct {
		name     string
		a        time.Time
		b        time.Time
		expected time.Time
	}{
		{
			name:     "one day apart",
			a:        time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			b:        time.Date(2026, 4, 16, 10, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 16, 10, 0, 0, 0, time.UTC),
		},
		{
			name:     "one week apart",
			a:        time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			b:        time.Date(2026, 4, 22, 10, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 22, 10, 0, 0, 0, time.UTC),
		},
		{
			name:     "one month apart",
			a:        time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			b:        time.Date(2026, 5, 15, 10, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 5, 15, 10, 0, 0, 0, time.UTC),
		},
		{
			name:     "one year apart",
			a:        time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			b:        time.Date(2027, 4, 15, 10, 0, 0, 0, time.UTC),
			expected: time.Date(2027, 4, 15, 10, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Max(tt.a, tt.b)
			if !result.Equal(tt.expected) {
				t.Errorf("Max(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMax_Timezones(t *testing.T) {
	est, _ := time.LoadLocation("America/New_York")
	pst, _ := time.LoadLocation("America/Los_Angeles")
	utc := time.UTC

	tests := []struct {
		name     string
		a        time.Time
		b        time.Time
		expected time.Time
	}{
		{
			name:     "same instant different zones",
			a:        time.Date(2026, 4, 15, 10, 0, 0, 0, est),
			b:        time.Date(2026, 4, 15, 7, 0, 0, 0, pst),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 0, est),
		},
		{
			name:     "UTC vs EST",
			a:        time.Date(2026, 4, 15, 10, 0, 0, 0, est),
			b:        time.Date(2026, 4, 15, 14, 0, 0, 0, utc),
			expected: time.Date(2026, 4, 15, 14, 0, 0, 0, utc),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Max(tt.a, tt.b)
			if !result.Equal(tt.expected) {
				t.Errorf("Max(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMax_NanosecondPrecision(t *testing.T) {
	tests := []struct {
		name     string
		a        time.Time
		b        time.Time
		expected time.Time
	}{
		{
			name:     "one nanosecond apart",
			a:        time.Date(2026, 4, 15, 10, 0, 0, 99, time.UTC),
			b:        time.Date(2026, 4, 15, 10, 0, 0, 100, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 100, time.UTC),
		},
		{
			name:     "microsecond apart",
			a:        time.Date(2026, 4, 15, 10, 0, 0, 999, time.UTC),
			b:        time.Date(2026, 4, 15, 10, 0, 0, 1000, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 1000, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Max(tt.a, tt.b)
			if !result.Equal(tt.expected) {
				t.Errorf("Max(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMax_ZeroValue(t *testing.T) {
	tests := []struct {
		name     string
		a        time.Time
		b        time.Time
		expected time.Time
	}{
		{
			name:     "zero vs now",
			a:        time.Time{},
			b:        time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
		},
		{
			name:     "both zero",
			a:        time.Time{},
			b:        time.Time{},
			expected: time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Max(tt.a, tt.b)
			if !result.Equal(tt.expected) {
				t.Errorf("Max(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMax_Commutative(t *testing.T) {
	t.Run("max is commutative", func(t *testing.T) {
		a := time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC)
		b := time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC)

		maxAB := lxtime.Max(a, b)
		maxBA := lxtime.Max(b, a)

		if !maxAB.Equal(maxBA) {
			t.Errorf("Max(a, b) = %v, Max(b, a) = %v, should be equal", maxAB, maxBA)
		}
	})
}

func TestMax_Idempotent(t *testing.T) {
	t.Run("max is idempotent", func(t *testing.T) {
		a := time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC)

		maxAA := lxtime.Max(a, a)

		if !maxAA.Equal(a) {
			t.Errorf("Max(a, a) = %v, want %v", maxAA, a)
		}
	})
}

func TestMinMax_Consistency(t *testing.T) {
	t.Run("min_and_max_cover_range", func(t *testing.T) {
		a := time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC)
		b := time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC)

		minVal := lxtime.Min(a, b)
		maxVal := lxtime.Max(a, b)

		if minVal.After(maxVal) {
			t.Errorf("Min(%v, %v) = %v is after Max(...) = %v", a, b, minVal, maxVal)
		}

		if (a.Equal(minVal) || b.Equal(minVal)) && (a.Equal(maxVal) || b.Equal(maxVal)) {
			// OK - min and max are both from the inputs
		} else {
			t.Errorf("Min and Max should return input values")
		}
	})
}

func ExampleMax() {
	t1 := time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC)
	later := lxtime.Max(t1, t2)
	// later: 2026-04-15 14:00:00 +0000 UTC
	_ = later
}
