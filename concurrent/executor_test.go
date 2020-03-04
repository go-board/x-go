package concurrent_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/go-board/x-go/concurrent"
)

func TestPool(t *testing.T) {
	t.Run("spawn", func(t *testing.T) {
		p := concurrent.NewPool(1, 10)
		now := time.Now()
		task := p.Spawn(context.Background(), func(ctx context.Context) error {
			time.Sleep(time.Second)
			return nil
		})
		err := task.Wait()
		require.Nil(t, err, "err must be nil")
		require.Equal(t, 1, int(time.Since(now).Seconds()), "task spent 1 second")
	})
	t.Run("spawn timeout", func(t *testing.T) {
		p := concurrent.NewPool(1, 10)
		now := time.Now()
		task := p.Spawn(context.Background(), func(ctx context.Context) error {
			time.Sleep(time.Second * 2)
			return nil
		})
		err := task.WaitTimeout(time.Second)
		require.NotNil(t, err, "err must be not nil")
		require.Equal(t, 1, int(time.Since(now).Seconds()), "task spent 1 second")
	})
}
