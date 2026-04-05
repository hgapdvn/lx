package lxtime

import (
	"fmt"
	"strings"
	"time"
)

// DurationString converts a time.Duration to a human-readable string format.
// Only includes non-zero components. Uses proper singular/plural forms.
// Returns "0 seconds" for zero duration.
//
// Example:
//
//	d := 2*time.Hour + 3*time.Minute
//	str := lxtime.DurationString(d)
//	// str: "2 hours, 3 minutes"
//
//	d := 1*time.Hour + 30*time.Second
//	str := lxtime.DurationString(d)
//	// str: "1 hour, 30 seconds"
func DurationString(d time.Duration) string {
	if d == 0 {
		return "0 seconds"
	}

	// Handle negative durations
	isNegative := d < 0
	if isNegative {
		d = -d
	}

	// Extract components in descending order
	var parts []string

	// Hours
	if hours := d / time.Hour; hours > 0 {
		parts = append(parts, pluralizeUnit(int(hours), "hour"))
		d -= hours * time.Hour
	}

	// Minutes
	if minutes := d / time.Minute; minutes > 0 {
		parts = append(parts, pluralizeUnit(int(minutes), "minute"))
		d -= minutes * time.Minute
	}

	// Seconds
	if seconds := d / time.Second; seconds > 0 {
		parts = append(parts, pluralizeUnit(int(seconds), "second"))
		d -= seconds * time.Second
	}

	// Milliseconds
	if millis := d / time.Millisecond; millis > 0 {
		parts = append(parts, pluralizeUnit(int(millis), "millisecond"))
		d -= millis * time.Millisecond
	}

	// Microseconds
	if micros := d / time.Microsecond; micros > 0 {
		parts = append(parts, pluralizeUnit(int(micros), "microsecond"))
		d -= micros * time.Microsecond
	}

	// Nanoseconds
	if nanos := d; nanos > 0 {
		parts = append(parts, pluralizeUnit(int(nanos), "nanosecond"))
	}

	// Join parts with commas
	result := strings.Join(parts, ", ")

	// Add negative prefix if needed
	if isNegative {
		result = "-" + result
	}

	return result
}

// pluralizeUnit returns a formatted string with singular or plural form
func pluralizeUnit(count int, unit string) string {
	if count == 1 {
		return fmt.Sprintf("%d %s", count, unit)
	}
	return fmt.Sprintf("%d %ss", count, unit)
}
