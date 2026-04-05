package lxtime

import "time"

// AddBusinessDays adds n business days (skipping Saturday and Sunday) to the given time.
// Negative values subtract business days.
// Returns a time on a weekday. If the result falls on a weekend, moves to the next/previous weekday.
//
// Example:
//
//	// Friday April 3, 2026 + 1 business day = Monday April 6
//	result := lxtime.AddBusinessDays(friday, 1)
//	// result: Monday April 6
//
//	// Friday April 3, 2026 + 3 business days = Wednesday April 8
//	result := lxtime.AddBusinessDays(friday, 3)
//	// result: Wednesday April 8
func AddBusinessDays(t time.Time, days int) time.Time {
	if days == 0 {
		return moveToWeekday(t)
	}

	current := startOfDay(t)

	direction := 1
	if days < 0 {
		direction = -1
		days = -days
	}

	added := 0
	for added < days {
		current = current.AddDate(0, 0, direction)

		switch current.Weekday() {
		case time.Saturday, time.Sunday:
			continue
		}

		added++
	}

	return restoreTime(current, t)
}

func moveToWeekday(t time.Time) time.Time {
	switch t.Weekday() {
	case time.Saturday:
		return t.AddDate(0, 0, 2)
	case time.Sunday:
		return t.AddDate(0, 0, 1)
	default:
		return t
	}
}

func startOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func restoreTime(date time.Time, original time.Time) time.Time {
	y, m, d := date.Date()
	h, min, s := original.Clock()

	return time.Date(
		y, m, d,
		h, min, s,
		original.Nanosecond(),
		original.Location(),
	)
}
