package lock

import (
	"context"
	"time"
)

type Locker interface {
	Do(ctx context.Context, name string, duration time.Duration, fn func(ctx context.Context) error) error
}
