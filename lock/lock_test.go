package lock

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRedisLocker(t *testing.T) {
	l := NewRedisLocker("127.0.0.1:6379", "", 0)
	testLocker(t, l)
}

func TestProcessLocker(t *testing.T) {
	l := NewProcessLocker()
	testLocker(t, l)
}

func testLocker(t *testing.T, l Locker) {
	t.Run("success", func(t *testing.T) {
		g := sync.WaitGroup{}
		g.Add(2)
		go func() {
			err := l.Do(context.Background(), "hello", time.Second, func(ctx context.Context) error {
				time.Sleep(time.Second * 3)
				return nil
			})
			require.Nil(t, err, "err must be nil")
			g.Done()
		}()
		go func() {
			time.Sleep(time.Second)
			err := l.Do(context.Background(), "world", time.Second, func(ctx context.Context) error {
				time.Sleep(time.Second * 3)
				return nil
			})
			require.Nil(t, err, "err must be nil")
			g.Done()
		}()
		g.Wait()
	})
	t.Run("success on previous timeout", func(t *testing.T) {
		g := sync.WaitGroup{}
		g.Add(2)
		go func() {
			err := l.Do(context.Background(), "hello", time.Second, func(ctx context.Context) error {
				time.Sleep(time.Second * 2)
				return nil
			})
			require.Nil(t, err, "err must be nil")
			g.Done()
		}()
		go func() {
			time.Sleep(time.Second * 2)
			err := l.Do(context.Background(), "hello", time.Second, func(ctx context.Context) error {
				time.Sleep(time.Second * 2)
				return nil
			})
			require.Nil(t, err, "err must be nil")
			g.Done()
		}()
		g.Wait()
	})
	t.Run("failed on previous valid", func(t *testing.T) {
		g := sync.WaitGroup{}
		g.Add(2)
		go func() {
			err := l.Do(context.Background(), "hello", time.Second*2, func(ctx context.Context) error {
				time.Sleep(time.Second * 2)
				return nil
			})
			require.Nil(t, err, "err must be nil")
			g.Done()
		}()
		go func() {
			time.Sleep(time.Second)
			err := l.Do(context.Background(), "hello", time.Second*2, func(ctx context.Context) error {
				time.Sleep(time.Second * 2)
				return nil
			})
			require.NotNil(t, err, "err must be not nil")
			t.Logf("%+v\n", err)
			g.Done()
		}()
		g.Wait()
	})
}

func TestLockerFunctional(t *testing.T) {
	t.Run("redis", func(t *testing.T) {
		l := NewRedisLocker("", "", 0)
		testFunctional(t, l)
	})
	t.Run("process", func(t *testing.T) {
		l := NewProcessLocker()
		testFunctional(t, l)
	})
}

func testFunctional(t *testing.T, l Locker) {
	t.Run("timeout", func(t *testing.T) {
		err := l.Do(context.Background(), "hello", time.Second, func(ctx context.Context) error {
			t := time.NewTimer(time.Second * 2)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-t.C:
				return nil
			}
		})
		require.NotNil(t, err, "err should not be nil")
	})
	t.Run("success", func(t *testing.T) {
		err := l.Do(context.Background(), "hello", time.Second*2, func(ctx context.Context) error {
			t := time.NewTimer(time.Second)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-t.C:
				return nil
			}
		})
		require.Nil(t, err, "err should be nil")
	})
}
