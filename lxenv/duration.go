package lxenv

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// durationRe matches a single numeric component followed by a unit in a duration string.
var durationRe = regexp.MustCompile(`^\s*([-+]?([0-9]*\.[0-9]+|[0-9]+))\s*([a-zA-Zµμ]+)`)

// GetDuration retrieves an environment variable as a duration.
// Supports standard Go duration strings and extended units: d (days), w (weeks), y (years).
// Returns (value, true) if the variable is set and can be parsed as a duration.
// Returns (0, false) if the variable is not set or cannot be parsed.
//
// Example:
//
//	if timeout, ok := lxenv.GetDuration("TIMEOUT"); ok {
//	    // Use timeout as time.Duration
//	}
func GetDuration(key string) (time.Duration, bool) {
	value := os.Getenv(key)
	if value == "" {
		return 0, false
	}
	parsed, err := parseDuration(value)
	if err != nil {
		return 0, false
	}
	return parsed, true
}

// GetDurationOr retrieves an environment variable as a duration or returns a default value.
// Returns the parsed duration if the variable is set and valid, otherwise returns defaultValue.
//
// Example:
//
//	timeout := lxenv.GetDurationOr("TIMEOUT", 30*time.Second)
//	// timeout: 30s if TIMEOUT is not set or invalid
func GetDurationOr(key string, defaultValue time.Duration) time.Duration {
	if value, ok := GetDuration(key); ok {
		return value
	}
	return defaultValue
}

// MustGetDuration retrieves an environment variable as a duration.
// Panics if the variable is not set or cannot be parsed as a duration.
//
// Example:
//
//	timeout := lxenv.MustGetDuration("TIMEOUT")
//	// timeout: 30s
func MustGetDuration(key string) time.Duration {
	value, ok := GetDuration(key)
	if !ok {
		panic("lxenv: environment variable " + key + " is not set or not a valid duration")
	}
	return value
}

// parseDuration parses a duration string that supports extended units:
// y (years), w (weeks), d (days).
//
// Units supported:
//   - ns, us, µs, μs, ms, s, m, h (Go standard)
//   - d (day = 24h)
//   - w (week = 7d)
//   - y (year = 365d)
//
// Examples:
//   - "3d" -> 72 * time.Hour
//   - "1w" -> 168 * time.Hour
//   - "1.5d" -> 36 * time.Hour
//   - "1d 12h" -> 36 * time.Hour
func parseDuration(s string) (time.Duration, error) {
	orig := strings.TrimSpace(s)
	if orig == "" {
		return 0, fmt.Errorf("lxenv: empty duration")
	}

	// Try standard parsing first for compatibility
	if d, err := time.ParseDuration(orig); err == nil {
		return d, nil
	}

	// Robust envLoader that consumes the string part by part
	// Supports optional spaces between value and unit, and between parts.
	remaining := orig
	var total time.Duration
	for remaining != "" {
		matches := durationRe.FindStringSubmatch(remaining)
		if matches == nil {
			return 0, fmt.Errorf("lxenv: invalid duration component in %q", orig)
		}

		fullMatch := matches[0]
		valStr := matches[1]
		unit := strings.ToLower(matches[3])

		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return 0, fmt.Errorf("lxenv: invalid value %q in duration %q", valStr, orig)
		}

		var factor time.Duration
		switch unit {
		case "y", "yr", "year", "years":
			factor = 365 * 24 * time.Hour
		case "w", "wk", "week", "weeks":
			factor = 7 * 24 * time.Hour
		case "d", "day", "days":
			factor = 24 * time.Hour
		case "h", "hr", "hour", "hours":
			factor = time.Hour
		case "m", "min", "minute", "minutes":
			factor = time.Minute
		case "s", "sec", "second", "seconds":
			factor = time.Second
		case "ms", "msec", "millisecond", "milliseconds":
			factor = time.Millisecond
		case "us", "µs", "μs", "usec", "microsecond", "microseconds":
			factor = time.Microsecond
		case "ns", "nsec", "nanosecond", "nanoseconds":
			factor = time.Nanosecond
		default:
			return 0, fmt.Errorf("lxenv: unknown unit %q in duration %q", unit, orig)
		}

		total += time.Duration(val * float64(factor))
		remaining = remaining[len(fullMatch):]
		remaining = strings.TrimLeft(remaining, " \t\n\r")
	}

	return total, nil
}
