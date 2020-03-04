package xyaml

import (
	"gopkg.in/yaml.v2"

	"github.com/go-board/x-go/xcodec"
)

const Name = "yaml"

type yamlc struct{}

func (yamlc) Name() string {
	return Name
}

func (yamlc) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (yamlc) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

func init() {
	xcodec.Register(yamlc{})
}
