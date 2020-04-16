package xsql

import (
	"database/sql/driver"
	"errors"
	"net"
	"strings"
)

// StringArray present a slice of array which store in database in string mode,
// but in application is string slice mode, this will translate automatically.
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

// Enum present a selectable of value which store in database in int mode,
// but in application is string mode, this will translate automatically.
type Enum struct {
	value    string
	valueMap map[string]int
	keyMap   map[int]string
}

func NewEnum(valueMap map[string]int, defs ...string) *Enum {
	def := ""
	if len(defs) > 0 {
		def = defs[0]
	}
	return &Enum{
		value:    def,
		valueMap: valueMap,
	}
}

func (c *Enum) String() string {
	return c.value
}

func (c *Enum) fromInt(i int) {
	c.value = c.keyMap[i]
}

func (c *Enum) Scan(src interface{}) error {
	switch x := src.(type) {
	case int:
		c.fromInt(x)
		return nil
	case int64:
		c.fromInt(int(x))
		return nil
	case int32:
		c.fromInt(int(x))
		return nil
	default:
		return nil
	}
}

func (c Enum) ConvertValue(v interface{}) (driver.Value, error) {
	switch x := v.(type) {
	case *Enum:
		return x.Value()
	case Enum:
		return x.Value()
	case string:
		return c.valueMap[x], nil
	default:
		return nil, nil
	}
}

func (c Enum) Value() (driver.Value, error) {
	return c.valueMap[c.value], nil
}

// IPV4 present a selectable of ipv4 address which store in database in int mode,
// but in application is ipv4 mode, this will translate automatically.
type IPV4 struct {
	ipStr string
	ip    net.IP
}

func NewIPV4(ip string) (*IPV4, error) {
	ipaddr := net.ParseIP(ip)
	if ipaddr == nil {
		return nil, errors.New("err: parse ipv4 failed: " + ip)
	}
	return &IPV4{ip: ipaddr, ipStr: ip}, nil
}

func (a IPV4) toUint32() uint32 {
	return uint32(a.ip[0])<<24 + uint32(a.ip[1])<<16 + uint32(a.ip[2])<<8 + uint32(a.ip[3])
}

func (a IPV4) IP() net.IP {
	return a.ip
}

func (a IPV4) String() string {
	return a.ipStr
}

func (a *IPV4) Scan(src interface{}) error {
	panic("implement me")
}

func (a IPV4) Value() (driver.Value, error) {
	return a.toUint32(), nil
}

func (a IPV4) ConvertValue(v interface{}) (driver.Value, error) {
	return a.Value()
}
