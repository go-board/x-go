package xctx

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestReadInt(t *testing.T) {
	ctx := context.Background()
	newCtx := context.WithValue(ctx, "int", 1)
	i, ok := ReadInt(newCtx, "int")
	require.Equal(t, true, ok)
	require.Equal(t, 1, i, "store value is 1")
}

func TestReadTime(t *testing.T) {
	ctx := context.Background()
	n := time.Now()
	newCtx := context.WithValue(ctx, "now", n)
	now, ok := ReadTime(newCtx, "now")
	require.Equal(t, true, ok)
	require.Equal(t, n.UnixNano(), now.UnixNano())
}
