package lxtime_test

import (
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/lxtime"
)

func TestEndOfWeek_BasicCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "returns Sunday end of day",
			check: func() bool {
				input := time.Date(2026, 4, 8, 15, 30, 0, 0, time.UTC)
				result := lxtime.EndOfWeek(input)
				return result.Weekday() == time.Sunday && result.Hour() == 23
			},
		},
		{
			name: "Wednesday goes to Sunday",
			check: func() bool {
				input := time.Date(2026, 4, 8, 15, 30, 0, 0, time.UTC)
				result := lxtime.EndOfWeek(input)
				expected := time.Date(2026, 4, 12, 23, 59, 59, 999999999, time.UTC)
				return result.Equal(expected)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("EndOfWeek() check failed")
			}
		})
	}
}

func TestEndOfWeek_TimezoneCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "EST preserves timezone",
			check: func() bool {
				est, _ := time.LoadLocation("America/New_York")
				// 2026-04-08 (Wednesday) in EST
				input := time.Date(2026, 4, 8, 10, 30, 0, 0, est)
				result := lxtime.EndOfWeek(input)
				// Should be Sunday 2026-04-12 at end of day in EST
				expected := time.Date(2026, 4, 12, 23, 59, 59, 999999999, est)
				return result.Equal(expected)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("EndOfWeek() timezone check failed")
			}
		})
	}
}
