package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxtime"
)

func TestMin_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		a        time.Time
		b        time.Time
		expected time.Time
	}{
		{
			name:     "a is earlier",
			a:        time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			b:        time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
		},
		{
			name:     "b is earlier",
			a:        time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			b:        time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
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
			result := lxtime.Min(tt.a, tt.b)
			if !result.Equal(tt.expected) {
				t.Errorf("Min(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMin_DifferentDays(t *testing.T) {
	tests := []struct {
		name     string
		a        time.Time
		b        time.Time
		expected time.Time
	}{
		{
			name:     "one day apart",
			a:        time.Date(2026, 4, 16, 10, 0, 0, 0, time.UTC),
			b:        time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
		},
		{
			name:     "one week apart",
			a:        time.Date(2026, 4, 22, 10, 0, 0, 0, time.UTC),
			b:        time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
		},
		{
			name:     "one month apart",
			a:        time.Date(2026, 5, 15, 10, 0, 0, 0, time.UTC),
			b:        time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
		},
		{
			name:     "one year apart",
			a:        time.Date(2027, 4, 15, 10, 0, 0, 0, time.UTC),
			b:        time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Min(tt.a, tt.b)
			if !result.Equal(tt.expected) {
				t.Errorf("Min(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMin_Timezones(t *testing.T) {
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
			a:        time.Date(2026, 4, 15, 14, 0, 0, 0, utc),
			b:        time.Date(2026, 4, 15, 10, 0, 0, 0, est),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 0, est),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Min(tt.a, tt.b)
			if !result.Equal(tt.expected) {
				t.Errorf("Min(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMin_NanosecondPrecision(t *testing.T) {
	tests := []struct {
		name     string
		a        time.Time
		b        time.Time
		expected time.Time
	}{
		{
			name:     "one nanosecond apart",
			a:        time.Date(2026, 4, 15, 10, 0, 0, 100, time.UTC),
			b:        time.Date(2026, 4, 15, 10, 0, 0, 99, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 99, time.UTC),
		},
		{
			name:     "microsecond apart",
			a:        time.Date(2026, 4, 15, 10, 0, 0, 1000, time.UTC),
			b:        time.Date(2026, 4, 15, 10, 0, 0, 999, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 999, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Min(tt.a, tt.b)
			if !result.Equal(tt.expected) {
				t.Errorf("Min(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMin_ZeroValue(t *testing.T) {
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
			expected: time.Time{},
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
			result := lxtime.Min(tt.a, tt.b)
			if !result.Equal(tt.expected) {
				t.Errorf("Min(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMin_Commutative(t *testing.T) {
	t.Run("min is commutative", func(t *testing.T) {
		a := time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC)
		b := time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC)

		minAB := lxtime.Min(a, b)
		minBA := lxtime.Min(b, a)

		if !minAB.Equal(minBA) {
			t.Errorf("Min(a, b) = %v, Min(b, a) = %v, should be equal", minAB, minBA)
		}
	})
}

func TestMin_Idempotent(t *testing.T) {
	t.Run("min is idempotent", func(t *testing.T) {
		a := time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC)

		minAA := lxtime.Min(a, a)

		if !minAA.Equal(a) {
			t.Errorf("Min(a, a) = %v, want %v", minAA, a)
		}
	})
}

func ExampleMin() {
	t1 := time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC)
	earlier := lxtime.Min(t1, t2)
	// earlier: 2026-04-15 10:00:00 +0000 UTC
	_ = earlier
}
