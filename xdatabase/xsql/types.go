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

func (a StringArray) String() string {
	return strings.Join(a, ",")
}

func (a *StringArray) Scan(src interface{}) error {
	switch x := src.(type) {
	case string:
		*a = strings.Split(x, ",")
		return nil
	case []byte:
		*a = strings.Split(string(x), ",")
		return nil
	default:
		return nil
	}
}

func (a StringArray) Value() (driver.Value, error) {
	x := strings.Join(a, ",")
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
	keyMap := make(map[int]string)
	for v, i := range valueMap {
		keyMap[i] = v
	}
	return &Enum{
		value:    def,
		valueMap: valueMap,
		keyMap:   keyMap,
	}
}

func (a *Enum) String() string {
	return a.value
}

func (a *Enum) fromInt(i int) {
	a.value = a.keyMap[i]
}

func (a *Enum) Scan(src interface{}) error {
	switch x := src.(type) {
	case int:
		a.fromInt(x)
		return nil
	case int64:
		a.fromInt(int(x))
		return nil
	case int32:
		a.fromInt(int(x))
		return nil
	default:
		return nil
	}
}

func (a Enum) Value() (driver.Value, error) {
	return a.valueMap[a.value], nil
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

func (a *IPV4) fromUint32(u uint32) error {
	a.ip[0] = uint8(u >> 24)
	a.ip[1] = uint8(u >> 16)
	a.ip[2] = uint8(u >> 8)
	a.ip[3] = uint8(u)
	a.ipStr = a.ip.String()
	return nil
}

func (a IPV4) IP() net.IP {
	return a.ip
}

func (a IPV4) String() string {
	return a.ipStr
}

func (a *IPV4) Scan(src interface{}) error {
	a.ip = a.ip[:0]
	switch x := src.(type) {
	case uint32:
		return a.fromUint32(x)
	case uint64:
		return a.fromUint32(uint32(x))
	case int64:
		return a.fromUint32(uint32(x))
	default:
		return errors.New("err: unsupported type")
	}
}

func (a IPV4) Value() (driver.Value, error) {
	return a.toUint32(), nil
}
