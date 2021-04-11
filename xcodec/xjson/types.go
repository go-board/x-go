package xjson

import (
	bytes2 "bytes"
	"fmt"
	"strconv"
)

// Int64String is a go int64 type.
//
// which can marshal into json string, and unmarshal from json string or json number.
//	type User struct {
//	  UID Int64String `json:"uid"`
//	}
// so we can marshal into
//	{"uid": "123456"}
// unmarshal from both
//	{"uid": 123456}
// and
//	{"uid": "123456"}
type Int64String int64

func (i Int64String) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", i)), nil
}

func (i *Int64String) UnmarshalJSON(bytes []byte) error {
	if bytes2.HasPrefix(bytes, []byte{'"'}) {
		bytes = bytes[1 : len(bytes)-1]
	}
	m, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		return err
	}
	*i = Int64String(m)
	return nil
}

func (i Int64String) Int64() int64 { return int64(i) }
func (i Int64String) Int32() int32 { return int32(i) }
func (i Int64String) Int16() int16 { return int16(i) }
func (i Int64String) Int8() int8   { return int8(i) }
func (i Int64String) Int() int     { return int(i) }
