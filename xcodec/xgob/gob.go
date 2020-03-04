package xgob

import (
	"bytes"
	"encoding/gob"

	"github.com/go-board/x-go/xcodec"
)

const Name = "gob"

type gobc struct{}

func (gobc) Name() string { return Name }

func (gobc) Marshal(v interface{}) ([]byte, error) {
	b := &bytes.Buffer{}
	err := gob.NewEncoder(b).Encode(v)
	return b.Bytes(), err
}

func (gobc) Unmarshal(data []byte, v interface{}) error {
	r := bytes.NewReader(data)
	return gob.NewDecoder(r).Decode(v)
}

func init() {
	xcodec.Register(gobc{})
}
