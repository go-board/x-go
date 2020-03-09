package xctx

import (
	"context"
	"time"
)

// ReadInt try read int stored in context.Context identified by key.
func ReadInt(ctx context.Context, key interface{}) (int, bool) {
	v := ctx.Value(key)
	if v == nil {
		return 0, false
	}
	switch i := v.(type) {
	case int:
		return i, true
	case int64:
		return int(i), true
	case int32:
		return int(i), true
	default:
		return 0, false
	}
}

// ReadInt64 try read int64 stored in context.Context identified by key.
func ReadInt64(ctx context.Context, key interface{}) (int64, bool) {
	v := ctx.Value(key)
	if v == nil {
		return 0, false
	}
	switch i := v.(type) {
	case int:
		return int64(i), true
	case int64:
		return i, true
	case int32:
		return int64(i), true
	default:
		return 0, false
	}
}

// ReadString try read string stored in context.Context identified by key.
func ReadString(ctx context.Context, key interface{}) (string, bool) {
	v := ctx.Value(key)
	if v == nil {
		return "", false
	}
	switch i := v.(type) {
	case string:
		return i, true
	case []byte:
		return string(i), true
	default:
		return "", false
	}
}

// ReadTime try read time.Time stored in context.Context identified by key.
func ReadTime(ctx context.Context, key interface{}) (time.Time, bool) {
	v := ctx.Value(key)
	if v == nil {
		return time.Time{}, false
	}
	switch i := v.(type) {
	case time.Time:
		return i, true
	case *time.Time:
		return *i, true
	default:
		return time.Time{}, false
	}
}
