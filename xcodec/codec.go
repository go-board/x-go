package xcodec

import (
	"fmt"
)

// Codec is definition of encoding can marshal/unmarshal data
type Codec interface {
	// Name is name/id of this codec strategy, like json/yaml etc...
	Name() string
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

var codecMap = make(map[string]Codec)

// Register register codec to global registry.
// If codec with same name exist, the new one will replace the old one
func Register(codec Codec) {
	if _, ok := codecMap[codec.Name()]; ok {
		panic(fmt.Sprintf("codec [%s] already registered", codec.Name()))
	}
	codecMap[codec.Name()] = codec
}

// Get get codec by name
func Get(name string) (Codec, bool) {
	codec, ok := codecMap[name]
	return codec, ok
}

// MustGet must get codec by name, if not exists, panic
func MustGet(name string) Codec {
	codec, ok := Get(name)
	if !ok {
		panic(fmt.Errorf("codec %s not exists", name))
	}
	return codec
}
