package xctx

import (
	"context"
	"reflect"
	"time"
)

// TypeName return unique type's name for different type
func TypeName(value interface{}) string {
	return typename(value)
}

func typename(value interface{}) string {
	rt := reflect.TypeOf(value)
	name := rt.String()

	// But for named types (or pointers to them), qualify with import path (but see inner comment).
	// Dereference one pointer looking for a named type.
	star := ""
	if rt.Name() == "" {
		if pt := rt; pt.Kind() == reflect.Ptr {
			star = "*"
			// NOTE: The following line should be rt = pt.Elem() to implement
			// what the comment above claims, but fixing it would break compatibility
			// with existing gobs.
			//
			// Given package p imported as "full/p" with these definitions:
			//     package p
			//     type T1 struct { ... }
			// this table shows the intended and actual strings used by gob to
			// name the types:
			//
			// Type      Correct string     Actual string
			//
			// T1        full/p.T1          full/p.T1
			// *T1       *full/p.T1         *full/p.T1
			//
			//rt = pt
			rt = pt.Elem()
		}
	}
	if rt.Name() != "" {
		if rt.PkgPath() == "" {
			name = star + rt.Name()
		} else {
			name = star + rt.PkgPath() + "." + rt.Name()
		}
	}
	return name
}

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
	c.ctx = context.WithValue(c.ctx, typename(v), v)
}

func (c *TypedContext) ReadTyped(v interface{}) (interface{}, bool) {
	return c.ReadNamedData(typename(v))
}

func (c *TypedContext) ReadNamedData(t string) (interface{}, bool) {
	val := c.ctx.Value(t)
	if val == nil {
		return nil, false
	}
	return val, true
}

func (c *TypedContext) WithInt(i int) { c.With(i) }

func (c *TypedContext) ReadInt() (int, bool) {
	val := c.ctx.Value(typename(int(0)))
	if val == nil {
		return 0, false
	}
	return val.(int), true
}

func (c *TypedContext) WithInt8(i int8) { c.With(i) }

func (c *TypedContext) ReadInt8() (int8, bool) {
	val := c.ctx.Value(typename(int8(0)))
	if val == nil {
		return 0, false
	}
	return val.(int8), true
}

func (c *TypedContext) WithInt16(i int16) { c.With(i) }

func (c *TypedContext) ReadInt16() (int16, bool) {
	val := c.ctx.Value(typename(int16(0)))
	if val == nil {
		return 0, false
	}
	return val.(int16), true
}

func (c *TypedContext) WithInt32(i int32) { c.With(i) }

func (c *TypedContext) ReadInt32() (int32, bool) {
	val := c.ctx.Value(typename(int32(0)))
	if val == nil {
		return 0, false
	}
	return val.(int32), true
}

func (c *TypedContext) WithInt64(i int64) { c.With(i) }

func (c *TypedContext) ReadInt64() (int64, bool) {
	val := c.ctx.Value(typename(int64(0)))
	if val == nil {
		return 0, false
	}
	return val.(int64), true
}

func (c *TypedContext) WithUint(i uint) { c.With(i) }

func (c *TypedContext) ReadUint() (uint, bool) {
	val := c.ctx.Value(typename(uint(0)))
	if val == nil {
		return 0, false
	}
	return val.(uint), true
}

func (c *TypedContext) WithUint8(i uint8) { c.With(i) }

func (c *TypedContext) ReadUint8() (uint8, bool) {
	val := c.ctx.Value(typename(uint8(0)))
	if val == nil {
		return 0, false
	}
	return val.(uint8), true
}

func (c *TypedContext) WithUint16(i uint16) { c.With(i) }

func (c *TypedContext) ReadUint16() (uint16, bool) {
	val := c.ctx.Value(typename(uint16(0)))
	if val == nil {
		return 0, false
	}
	return val.(uint16), true
}

func (c *TypedContext) WithUint32(i uint32) { c.With(i) }

func (c *TypedContext) ReadUint32() (uint32, bool) {
	val := c.ctx.Value(typename(uint32(0)))
	if val == nil {
		return 0, false
	}
	return val.(uint32), true
}

func (c *TypedContext) WithUint64(i uint64) { c.With(i) }

func (c *TypedContext) ReadUint64() (uint64, bool) {
	val := c.ctx.Value(typename(uint64(0)))
	if val == nil {
		return 0, false
	}
	return val.(uint64), true
}

func (c *TypedContext) WithFloat32(f float32) { c.With(f) }

func (c *TypedContext) ReadFloat32() (float32, bool) {
	val := c.ctx.Value(typename(float32(0)))
	if val == nil {
		return 0, false
	}
	return val.(float32), true
}

func (c *TypedContext) WithFloat64(f float64) { c.With(f) }

func (c *TypedContext) ReadFloat64() (float64, bool) {
	val := c.ctx.Value(typename(float64(0)))
	if val == nil {
		return 0, false
	}
	return val.(float64), true
}

func (c *TypedContext) WithBool(b bool) { c.With(b) }

func (c *TypedContext) ReadBool() (bool, bool) {
	val := c.ctx.Value(typename(false))
	if val == nil {
		return false, false
	}
	return val.(bool), true
}

func (c *TypedContext) WithString(s string) { c.With(s) }

func (c *TypedContext) ReadString() (string, bool) {
	val := c.ctx.Value(typename(""))
	if val == nil {
		return "", false
	}
	return val.(string), true
}

func (c *TypedContext) WithTime(t time.Time) {
	c.ctx = context.WithValue(c.ctx, typename(t), t)
}

func (c *TypedContext) ReadTime() (time.Time, bool) {
	val := c.ctx.Value(typename(time.Time{}))
	if val == nil {
		return time.Time{}, false
	}
	return val.(time.Time), true
}
