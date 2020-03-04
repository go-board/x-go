package concurrent

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Concurrent struct {
	group       *errgroup.Group
	panicHandle func(x interface{})
}

func (c *Concurrent) SpawnContext(ctx context.Context, fn func(ctx context.Context) error) {
	c.group.Go(func() error {
		defer func() {
			if x := recover(); x != nil {
				if c.panicHandle != nil {
					c.panicHandle(x)
				}
			}
		}()
		return fn(ctx)
	})
}

func (c *Concurrent) Spawn(fn func() error) {
	c.group.Go(func() error {
		defer func() {
			if x := recover(); x != nil {
				if c.panicHandle != nil {
					c.panicHandle(x)
				}
			}
		}()
		return fn()
	})
}

func (c *Concurrent) Wait(ctx context.Context) error {
	return c.group.Wait()
}

func (c *Concurrent) SpawnAndWait(ctx context.Context, fns ...func(context.Context) error) error {
	for _, fn := range fns {
		c.SpawnContext(ctx, fn)
	}
	return c.Wait(ctx)
}

func NewConcurrent(ctx context.Context) (*Concurrent, context.Context) {
	group, ctx := errgroup.WithContext(ctx)
	return &Concurrent{group: group}, ctx
}
