package xjson

import (
	"encoding/json"

	"github.com/go-board/x-go/xcodec"
)

const Name = "json"

type jsonc struct{}

func (jsonc) Name() string { return Name }

func (jsonc) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonc) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func init() {
	xcodec.Register(jsonc{})
}
