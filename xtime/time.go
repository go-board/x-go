package xtime

import (
	"time"
)

// UTC return utc time
func UTC() time.Time {
	return time.Now().In(time.UTC)
}

// Local return local time
func Local() time.Time {
	return time.Now()
}

// NowMillis return now time in milliseconds
func NowMillis() int64 {
	return time.Now().UnixNano() / 1e6
}

func Yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

func Tomorrow() time.Time {
	return time.Now().AddDate(0, 0, 1)
}

func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 999999999, t.Location())
}

func StartOfToday() time.Time {
	return StartOfDay(time.Now())
}

func EndOfToday() time.Time {
	return EndOfDay(time.Now())
}
