package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/time"
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
