package time

import "time"

var _now func() time.Time = time.Now

// DaysBetween calculates the amount of days between two dates.
// The order of the parameters does not matter. The second date can be before the first date.
func DaysBetween(a, b time.Time) int {
	if a.After(b) {
		a, b = b, a
	}

	return int(b.Sub(a).Hours() / 24)
}

// DaysFromNow calculates the amount of days between now and the given date.
func DaysFromNow(t time.Time) int {
	return DaysBetween(_now(), t)
}

// SecondsBetween calculates the amount of seconds between two dates.
// The order of the parameters does not matter. The second date can be before the first date.
func SecondsBetween(a, b time.Time) int64 {
	if a.After(b) {
		a, b = b, a
	}
	return b.Unix() - a.Unix()
}

// SecondsFromNow calculates the amount of seconds between now and the given date.
func SecondsFromNow(t time.Time) int64 {
	return SecondsBetween(_now(), t)
}
