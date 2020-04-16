package xcodec_test

//
//import (
//	"encoding/gob"
//	"testing"
//
//	"github.com/stretchr/testify/require"
//
//
//	"github.com/go-board/x-go/xcodec"
//	"github.com/go-board/x-go/xcodec/xgob"
//	"github.com/go-board/x-go/xcodec/xjson"
//)
//
//type testStruct struct {
//	A int
//	B string
//}
//
//func init() {
//	gob.Register(&testStruct{})
//}
//
//func TestCodec(t *testing.T) {
//	t.Run("gob", func(t *testing.T) {
//		codec := xcodec.Get(xgob.Name)
//		testCodec(t, codec)
//	})
//	t.Run("json", func(t *testing.T) {
//		codec := xcodec.Get(xjson.Name)
//		testCodec(t, codec)
//	})
//	t.Run("proto", func(t *testing.T) {
//		codec := xcodec.Get(xproto.Name)
//		bytes, err := codec.Marshal(&internal.Test{
//			A: "hello",
//			B: true,
//			C: 3,
//		})
//		require.Nil(t, err, "err must be nil")
//		x := new(internal.Test)
//		err = codec.Unmarshal(bytes, x)
//		require.Nil(t, err, "err must be nil")
//		require.Equal(t, "hello", x.A, "x.A must be string(hello)")
//		require.Equal(t, true, x.B, "x.B must be bool(false)")
//		require.Equal(t, int64(3), x.C, "x.C must be integer(3)")
//	})
//	t.Run("toml", func(t *testing.T) {
//		codec := xcodec.Get(xtoml.Name)
//		testCodec(t, codec)
//	})
//	t.Run("yaml", func(t *testing.T) {
//		codec := xcodec.Get(xyaml.Name)
//		testCodec(t, codec)
//	})
//}
//
//func testCodec(t *testing.T, codec xcodec.Codec) {
//	bytes, err := codec.Marshal(&testStruct{A: 1, B: "hello"})
//	require.Nil(t, err, "err must be nil")
//	x := new(testStruct)
//	err = codec.Unmarshal(bytes, x)
//	require.Nil(t, err, "err must be nil")
//	require.Equal(t, 1, x.A, "x.A must be integer(1)")
//	require.Equal(t, "hello", x.B, "x.B must be string(hello)")
//}
