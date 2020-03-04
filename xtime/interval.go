package xtime

import (
	"context"
	"time"
)

// Interval return a chan chan emit event every duration and a cancel function to cancel interval.
func Interval(ctx context.Context, duration time.Duration) (chan struct{}, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx)
	ch := make(chan struct{}, 1)
	go func() {
		ticker := time.NewTicker(duration)
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case <-ticker.C:
				ch <- struct{}{}
			}
		}
	}()
	return ch, cancel
}

// RunInterval run user-defined function every duration until cancel.
func RunInterval(ctx context.Context, duration time.Duration, fn func()) context.CancelFunc {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		ticker := time.NewTicker(duration)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				fn()
			}
		}
	}()
	return cancel
}
