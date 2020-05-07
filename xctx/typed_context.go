package xctx

import (
	"context"
	"time"

	"github.com/go-board/x-go/types"
)

// TypedContext store typed data in context and can retrieve data with type.
type TypedContext struct{ ctx context.Context }

// NewTyped create derived context from parent context.Context
func NewTyped(ctx context.Context) *TypedContext {
	return &TypedContext{ctx: ctx}
}

func (c *TypedContext) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *TypedContext) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *TypedContext) Err() error {
	return c.ctx.Err()
}

func (c *TypedContext) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}

func (c *TypedContext) With(v interface{}) {
	c.ctx = context.WithValue(c.ctx, types.Typeof(v), v)
}

func (c *TypedContext) ReadTyped(v interface{}) (interface{}, bool) {
	return c.ReadNamedData(types.Typeof(v))
}

func (c *TypedContext) ReadNamedData(t types.TypeId) (interface{}, bool) {
	val := c.ctx.Value(t)
	if val == nil {
		return nil, false
	}
	return val, true
}

func (c *TypedContext) WithInt(i int) { c.With(i) }

func (c *TypedContext) ReadInt() (int, bool) {
	val := c.ctx.Value(types.Typeof(int(0)))
	if val == nil {
		return 0, false
	}
	return val.(int), true
}

func (c *TypedContext) WithInt8(i int8) { c.With(i) }

func (c *TypedContext) ReadInt8() (int8, bool) {
	val := c.ctx.Value(types.Typeof(int8(0)))
	if val == nil {
		return 0, false
	}
	return val.(int8), true
}

func (c *TypedContext) WithInt16(i int16) { c.With(i) }

func (c *TypedContext) ReadInt16() (int16, bool) {
	val := c.ctx.Value(types.Typeof(int16(0)))
	if val == nil {
		return 0, false
	}
	return val.(int16), true
}

func (c *TypedContext) WithInt32(i int32) { c.With(i) }

func (c *TypedContext) ReadInt32() (int32, bool) {
	val := c.ctx.Value(types.Typeof(int32(0)))
	if val == nil {
		return 0, false
	}
	return val.(int32), true
}

func (c *TypedContext) WithInt64(i int64) { c.With(i) }

func (c *TypedContext) ReadInt64() (int64, bool) {
	val := c.ctx.Value(types.Typeof(int64(0)))
	if val == nil {
		return 0, false
	}
	return val.(int64), true
}

func (c *TypedContext) WithUint(i uint) { c.With(i) }

func (c *TypedContext) ReadUint() (uint, bool) {
	val := c.ctx.Value(types.Typeof(uint(0)))
	if val == nil {
		return 0, false
	}
	return val.(uint), true
}

func (c *TypedContext) WithUint8(i uint8) { c.With(i) }

func (c *TypedContext) ReadUint8() (uint8, bool) {
	val := c.ctx.Value(types.Typeof(uint8(0)))
	if val == nil {
		return 0, false
	}
	return val.(uint8), true
}

func (c *TypedContext) WithUint16(i uint16) { c.With(i) }

func (c *TypedContext) ReadUint16() (uint16, bool) {
	val := c.ctx.Value(types.Typeof(uint16(0)))
	if val == nil {
		return 0, false
	}
	return val.(uint16), true
}

func (c *TypedContext) WithUint32(i uint32) { c.With(i) }

func (c *TypedContext) ReadUint32() (uint32, bool) {
	val := c.ctx.Value(types.Typeof(uint32(0)))
	if val == nil {
		return 0, false
	}
	return val.(uint32), true
}

func (c *TypedContext) WithUint64(i uint64) { c.With(i) }

func (c *TypedContext) ReadUint64() (uint64, bool) {
	val := c.ctx.Value(types.Typeof(uint64(0)))
	if val == nil {
		return 0, false
	}
	return val.(uint64), true
}

func (c *TypedContext) WithFloat32(f float32) { c.With(f) }

func (c *TypedContext) ReadFloat32() (float32, bool) {
	val := c.ctx.Value(types.Typeof(float32(0)))
	if val == nil {
		return 0, false
	}
	return val.(float32), true
}

func (c *TypedContext) WithFloat64(f float64) { c.With(f) }

func (c *TypedContext) ReadFloat64() (float64, bool) {
	val := c.ctx.Value(types.Typeof(float64(0)))
	if val == nil {
		return 0, false
	}
	return val.(float64), true
}

func (c *TypedContext) WithBool(b bool) { c.With(b) }

func (c *TypedContext) ReadBool() (bool, bool) {
	val := c.ctx.Value(types.Typeof(false))
	if val == nil {
		return false, false
	}
	return val.(bool), true
}

func (c *TypedContext) WithString(s string) { c.With(s) }

func (c *TypedContext) ReadString() (string, bool) {
	val := c.ctx.Value(types.Typeof(""))
	if val == nil {
		return "", false
	}
	return val.(string), true
}

func (c *TypedContext) WithTime(t time.Time) {
	c.ctx = context.WithValue(c.ctx, types.Typeof(t), t)
}

func (c *TypedContext) ReadTime() (time.Time, bool) {
	val := c.ctx.Value(types.Typeof(time.Time{}))
	if val == nil {
		return time.Time{}, false
	}
	return val.(time.Time), true
}
