package xproto

import (
	"errors"

	"github.com/golang/protobuf/proto"

	"github.com/go-board/x-go/xcodec"
)

const Name = "protobuf"

var ErrNotProtoMessage = errors.New("err: not proto message")

type protoc struct{}

func (protoc) Name() string { return Name }

func (protoc) Marshal(v interface{}) ([]byte, error) {
	if msg, ok := v.(proto.Message); ok {
		return proto.Marshal(msg)
	}
	return nil, ErrNotProtoMessage
}

func (protoc) Unmarshal(data []byte, v interface{}) error {
	if msg, ok := v.(proto.Message); ok {
		return proto.Unmarshal(data, msg)
	}
	return ErrNotProtoMessage
}

func init() {
	xcodec.Register(protoc{})
}
