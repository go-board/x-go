package xsql

import (
	"database/sql/driver"
	"strings"
)

type StringArray []string

func (c *StringArray) Scan(src interface{}) error {
	switch x := src.(type) {
	case string:
		*c = strings.Split(x, ",")
		return nil
	case []byte:
		*c = strings.Split(string(x), ",")
		return nil
	default:
		return nil
	}
}

func (c StringArray) ConvertValue(v interface{}) (driver.Value, error) {
	switch x := v.(type) {
	case *StringArray:
		return x.Value()
	case StringArray:
		return x.Value()
	case []string:
		return StringArray(x).Value()
	case string:
		return x, nil
	case []byte:
		return x, nil
	default:
		return nil, nil
	}
}

func (c StringArray) Value() (driver.Value, error) {
	x := strings.Join(c, ",")
	return x, nil
}
