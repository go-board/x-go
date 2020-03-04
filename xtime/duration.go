package xtime

import (
	"time"
)

// MeasureDuration will measure duration of an operation.
func MeasureDuration(fn func()) time.Duration {
	now := time.Now()
	fn()
	return time.Since(now)
}
