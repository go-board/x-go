package xtoml

import (
	"bytes"

	"github.com/BurntSushi/toml"

	"github.com/go-board/x-go/xcodec"
)

const Name = "toml"

type tomlc struct{}

func (tomlc) Name() string { return Name }

func (tomlc) Marshal(v interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	err := toml.NewEncoder(buf).Encode(v)
	return buf.Bytes(), err
}

func (tomlc) Unmarshal(data []byte, v interface{}) error {
	return toml.Unmarshal(data, v)
}

func init() {
	xcodec.Register(tomlc{})
}
