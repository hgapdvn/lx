package lxtime_test

import (
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/lxtime"
)

func TestEndOfDay_BasicCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "returns end of day",
			check: func() bool {
				input := time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC)
				result := lxtime.EndOfDay(input)
				return result.Hour() == 23 && result.Minute() == 59 && result.Second() == 59
			},
		},
		{
			name: "preserves date",
			check: func() bool {
				input := time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC)
				result := lxtime.EndOfDay(input)
				return result.Year() == 2026 && result.Month() == time.April && result.Day() == 4
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("EndOfDay() check failed")
			}
		})
	}
}

func TestEndOfDay_TimezoneCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "EST preserves correct date",
			check: func() bool {
				est, _ := time.LoadLocation("America/New_York")
				input := time.Date(2026, 4, 4, 10, 30, 0, 0, est)
				result := lxtime.EndOfDay(input)
				expected := time.Date(2026, 4, 4, 23, 59, 59, 999999999, est)
				return result.Equal(expected)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("EndOfDay() timezone check failed")
			}
		})
	}
}
