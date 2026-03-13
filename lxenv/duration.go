package lxenv

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var durationRE = regexp.MustCompile(`^([-+]?([0-9]*\.[0-9]+|[0-9]+))([a-zA-Zµμ]+)$`)

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
func parseDuration(s string) (time.Duration, error) {
	if s == "" {
		return 0, fmt.Errorf("lxenv: empty duration")
	}

	// Try standard parsing first for compatibility
	if d, err := time.ParseDuration(s); err == nil {
		return d, nil
	}

	// For extended units or combined units not handled by stdlib
	// We'll support simple ones first as requested (3d, etc.)
	// If they want combined like "1d2h", we'll need a more robust parser.
	// Let's implement a robust one that handles multiple parts.

	re := regexp.MustCompile(`([-+]?([0-9]*\.[0-9]+|[0-9]+))([a-zA-Zµμ]+)`)
	matches := re.FindAllStringSubmatch(s, -1)
	if len(matches) == 0 {
		return 0, fmt.Errorf("lxenv: invalid duration %q", s)
	}

	var total time.Duration
	for _, match := range matches {
		valStr := match[1]
		unit := match[3]

		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return 0, fmt.Errorf("lxenv: invalid value %q in duration %q", valStr, s)
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
			return 0, fmt.Errorf("lxenv: unknown unit %q in duration %q", unit, s)
		}

		total += time.Duration(val * float64(factor))
	}

	return total, nil
}
