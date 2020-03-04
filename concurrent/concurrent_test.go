package concurrent_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/go-board/x-go/concurrent"
)

func TestConcurrent(t *testing.T) {
	t.Run("wait success", func(t *testing.T) {
		now := time.Now()
		c, _ := concurrent.NewConcurrent(context.Background())
		c.Spawn(func() error {
			time.Sleep(time.Second)
			return nil
		})
		c.Spawn(func() error {
			time.Sleep(time.Second * 2)
			return nil
		})
		err := c.Wait(context.Background())
		require.Nil(t, err, "errCh must be nil")
		require.Equal(t, 2, int(time.Since(now).Seconds()), "longest elapsed time is two seconds")
	})
	t.Run("wait error", func(t *testing.T) {
		now := time.Now()
		c, _ := concurrent.NewConcurrent(context.Background())
		c.Spawn(func() error {
			time.Sleep(time.Second)
			return nil
		})
		c.Spawn(func() error {
			time.Sleep(time.Second * 2)
			return errors.New("not found")
		})
		c.Spawn(func() error {
			time.Sleep(time.Second * 3)
			return nil
		})
		err := c.Wait(context.Background())
		require.EqualError(t, err, "not found", "errCh must be not found")
		require.Equal(t, 3, int(time.Since(now).Seconds()), "longest elapsed time is two seconds")
	})
	t.Run("wait error last", func(t *testing.T) {
		now := time.Now()
		c, _ := concurrent.NewConcurrent(context.Background())
		var a int
		c.Spawn(func() error {
			a = 123
			time.Sleep(time.Second)
			return nil
		})
		c.Spawn(func() error {
			time.Sleep(time.Second * 2)
			return nil
		})
		c.Spawn(func() error {
			time.Sleep(time.Second * 3)
			return errors.New("timeout")
		})
		err := c.Wait(context.Background())
		t.Logf("%+v\n", a)
		require.EqualError(t, err, "timeout", "errCh must be timeout")
		require.Equal(t, 3, int(time.Since(now).Seconds()), "longest elapsed time is two seconds")
	})
}
