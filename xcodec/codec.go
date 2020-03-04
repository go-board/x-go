package xcodec

import (
	"fmt"
)

// Codec is definition of encoding can marshal/unmarshal data
type Codec interface {
	Name() string
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

var codecMap = make(map[string]Codec)

func Register(codec Codec) {
	if _, ok := codecMap[codec.Name()]; ok {
		panic(fmt.Sprintf("codec [%s] already registered", codec.Name()))
	}
	codecMap[codec.Name()] = codec
}

func Get(name string) Codec {
	return codecMap[name]
}
