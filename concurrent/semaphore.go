package concurrent

import (
	"context"
	"errors"

	"golang.org/x/sync/semaphore"
)

type Semaphore struct {
	w *semaphore.Weighted
}

func (s *Semaphore) SpawnContext(ctx context.Context, fn func(ctx context.Context) error) error {
	if err := s.w.Acquire(ctx, 1); err != nil {
		return err
	}
	defer s.w.Release(1)
	return fn(ctx)
}

func (s *Semaphore) Spawn(fn func() error) error {
	if err := s.w.Acquire(context.Background(), 1); err != nil {
		return err
	}
	defer s.w.Release(1)
	return fn()
}

func (s *Semaphore) TrySpawnContext(ctx context.Context, fn func(ctx context.Context) error) error {
	if !s.w.TryAcquire(1) {
		return errors.New("err: try acquire semaphore failed")
	}
	defer s.w.Release(1)
	return fn(ctx)
}

func (s *Semaphore) TrySpawn(fn func() error) error {
	if !s.w.TryAcquire(1) {
		return errors.New("err: try acquire semaphore failed")
	}
	defer s.w.Release(1)
	return fn()
}

func NewSemaphore(n int64) *Semaphore {
	return &Semaphore{w: semaphore.NewWeighted(n)}
}
