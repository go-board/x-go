package lock

import (
	"context"
	"time"

	"golang.org/x/sync/singleflight"
)

type processLocker struct {
	locker *singleflight.Group
}

func (m *processLocker) Do(ctx context.Context, name string, duration time.Duration, fn func(ctx context.Context) error) error {
	ctx, _ = context.WithTimeout(ctx, duration)
	resCh := m.locker.DoChan(name, func() (i interface{}, err error) {
		err = fn(ctx)
		return
	})
	defer m.locker.Forget(name)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case res := <-resCh:
		return res.Err
	}
}

func NewProcessLocker() Locker {
	return &processLocker{locker: &singleflight.Group{}}
}
