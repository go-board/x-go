package xstrings

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-board/x-go/xcodec"
)

func JoinInt(i []int, sep string) string {
	buf := strings.Builder{}
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	}

	return buf.String()
}

func JoinInt8(i []int8, sep string) string {
	buf := strings.Builder{}
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	}

	return buf.String()
}

func JoinInt16(i []int16, sep string) string {
	buf := strings.Builder{}
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	}

	return buf.String()
}

func JoinInt32(i []int32, sep string) string {
	buf := strings.Builder{}
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	}

	return buf.String()
}

func JoinInt64(i []int64, sep string) string {
	buf := strings.Builder{}
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatInt(v, 10))
	}

	return buf.String()
}

func JoinUint(i []uint, sep string) string {
	buf := strings.Builder{}
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatUint(uint64(v), 10))
	}

	return buf.String()
}

func JoinUint8(i []uint8, sep string) string {
	buf := strings.Builder{}
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatUint(uint64(v), 10))
	}

	return buf.String()
}

func JoinUint16(i []uint16, sep string) string {
	buf := strings.Builder{}
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatUint(uint64(v), 10))
	}

	return buf.String()
}

func JoinUint32(i []uint32, sep string) string {
	buf := strings.Builder{}
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatUint(uint64(v), 10))
	}

	return buf.String()
}

func JoinUint64(i []uint64, sep string) string {
	buf := strings.Builder{}
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatUint(v, 10))
	}

	return buf.String()
}

func JoinStringer(i []fmt.Stringer, sep string) string {
	buf := strings.Builder{}
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(v.String())
	}

	return buf.String()
}

func JoinAny(i []interface{}, sep string, codec xcodec.Codec) (string, error) {
	buf := strings.Builder{}
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		bytes, err := codec.Marshal(v)
		if err != nil {
			return "", err
		}
		buf.Write(bytes)
	}

	return buf.String(), nil
}
