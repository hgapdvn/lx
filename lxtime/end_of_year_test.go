package lxtime_test

import (
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/lxtime"
)

func TestEndOfYear_Basic(t *testing.T) {
	input := time.Date(2026, 6, 15, 15, 30, 0, 0, time.UTC)
	result := lxtime.EndOfYear(input)
	expected := time.Date(2026, 12, 31, 23, 59, 59, 999999999, time.UTC)
	if !result.Equal(expected) {
		t.Errorf("EndOfYear() = %v, want %v", result, expected)
	}
}

func TestEndOfYear_Timezone(t *testing.T) {
	est, _ := time.LoadLocation("America/New_York")
	input := time.Date(2026, 6, 15, 15, 30, 0, 0, est)
	result := lxtime.EndOfYear(input)
	if result.Location() != est {
		t.Errorf("EndOfYear() timezone not preserved")
	}
}
