package xcodec

import (
	"encoding/binary"
	"io"
)

type StreamCodec interface {
	Encode(w io.Writer, v interface{}) error
	Decode(r io.Reader, v interface{}) error
	Name() string
}

type StreamWrite interface {
	Write(p []byte) error
}

type StreamWriteExt interface {
	StreamWrite
	ByteOrder() binary.ByteOrder
	Codec() Codec
	WriteInt8(i int8) error
	WriteInt16(i int16) error
	WriteInt32(i int32) error
	WriteInt64(i int64) error

	WriteUint8(i uint8) error
	WriteUint16(i uint16) error
	WriteUint32(i uint32) error
	WriteUint64(i uint64) error

	WriteFloat32(f float32) error
	WriteFloat64(f float64) error

	WriteString(str string) error

	WriteBool(b bool) error

	WriteAny(v interface{}) error
}

type StreamRead interface {
	Read(p []byte) error
}

type StreamReadExt interface {
	StreamRead
	ByteOrder() binary.ByteOrder
	Codec() Codec
	ReadInt8(i *int8) error
	ReadInt16(i *int16) error
	ReadInt32(i *int32) error
	ReadInt64(i *int64) error

	ReadUint8(i *uint8) error
	ReadUint16(i *uint16) error
	ReadUint32(i *uint32) error
	ReadUint64(i *uint64) error

	ReadFloat32(f *float32) error
	ReadFloat64(f *float64) error

	ReadString(str *string) error

	ReadBool(b *bool) error

	ReadAny(v interface{}) error
}
