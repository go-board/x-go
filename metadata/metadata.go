package metadata

import (
	"context"
	"strings"

	"go.uber.org/zap/zapcore"
)

// Metadata is an attribute list, and will be propagation to the whole call chain.
// In process, will inject to context.Context.
// Cross process, will serialize to header for HTTPHeader/GRPC.
type Metadata map[string]string

func (md Metadata) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	for k, v := range md {
		enc.AddString(k, v)
	}
	return nil
}

// Clone return the copy of original Metadata
func (md Metadata) Clone() Metadata {
	newMd := make(Metadata, len(md))
	for k, v := range md {
		newMd[k] = v
	}
	return newMd
}

func (md Metadata) ValueList(key string) []string { return strings.Split(md[key], ",") }

var metadataKey = struct{}{}

// FromContext retrieve Metadata from context.Context with default metadataKey.
//
// 使用默认的key从context.Context中取出Metadata，如果不存在默认的key，会发生panic
func FromContext(ctx context.Context) Metadata {
	return FromContextKey(ctx, metadataKey)
}

// FromContextKey retrieve Metadata from context.Context with given key, if nil, panic.
//
// 使用给定的key从context.Context中取出Metadata, 如果不存在给定的key，会发生panic
func FromContextKey(ctx context.Context, key interface{}) Metadata {
	return ctx.Value(key).(Metadata)
}

// TryFromContext retrieve Metadata from context.Context with default metadataKey.
//
// 尝试使用默认的key从context.Context中取出Metadata
func TryFromContext(ctx context.Context) (Metadata, bool) {
	return TryFromContextKey(ctx, metadataKey)
}

// TryFromContextKey retrieve Metadata from context.Context with given key.
//
// 尝试使用给定的key从context.Context中取出Metadata
func TryFromContextKey(ctx context.Context, key interface{}) (Metadata, bool) {
	val := ctx.Value(key)
	if val == nil {
		return nil, false
	}
	md, ok := val.(Metadata)
	return md, ok
}

// IntoContext will append Metadata to context.Context use default metadataKey, ignore previous one.
//
// 使用默认的metadataKey讲Metadata注入到context.Context中，会忽略掉之前设置的。
func IntoContext(ctx context.Context, md Metadata) context.Context {
	return IntoContextKey(ctx, metadataKey, md)
}

// IntoContextKey will append Metadata to context.Context use given key, ignore previous one.
//
// 使用给定的key讲Metadata注入到context.Context中, 会忽略掉之前设置的。
func IntoContextKey(ctx context.Context, key interface{}, md Metadata) context.Context {
	return context.WithValue(ctx, key, md)
}

// MergeIntoContext will append Metadata to context.Context use default metadataKey, if has previous one, merge then update.
//
// 使用默认的metadataKey讲Metadata注入到context.Context中，如果存在上一个，先合并再更新
func MergeIntoContext(ctx context.Context, md Metadata) context.Context {
	oldOne, ok := TryFromContext(ctx)
	if !ok {
		return IntoContext(ctx, md)
	}
	for k, v := range oldOne {
		md[k] = v
	}
	return IntoContext(ctx, md)
}

// MergeIntoContextKey will append Metadata to context.Context use default metadataKey, if has previous one, merge then update.
//
// 使用给定的keyey讲Metadata注入到context.Context中，如果存在上一个，先合并再更新
func MergeIntoContextKey(ctx context.Context, key interface{}, md Metadata) context.Context {
	oldOne, ok := TryFromContextKey(ctx, key)
	if !ok {
		return IntoContextKey(ctx, key, md)
	}
	for k, v := range oldOne {
		md[k] = v
	}
	return IntoContextKey(ctx, key, md)
}
