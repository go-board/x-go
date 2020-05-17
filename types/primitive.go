package types

import (
	"reflect"
)

// IsPrimitiveType return whether type of v is primitive type or not.
func IsPrimitiveType(v interface{}) bool {
	rt := reflect.TypeOf(v)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	k := rt.Kind()
	return k == reflect.Int || k == reflect.Int8 || k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64 ||
		k == reflect.Uint || k == reflect.Uint8 || k == reflect.Uint16 || k == reflect.Uint32 || k == reflect.Uint64 || k == reflect.Uintptr ||
		k == reflect.Complex64 || k == reflect.Complex128 ||
		k == reflect.Float32 || k == reflect.Float64 || k == reflect.Bool || k == reflect.String
}

// ZeroValue generate zero value of given data.
// For example:
//    give int 8 will generate int 0
//    give string "hello" will generate string ""
//    give struct User{name: "lily", age: 123} will generate struct User{}
func ZeroValue(v interface{}) interface{} {
	return reflect.Zero(reflect.TypeOf(v)).Interface()
}

// IsSameType check a's runtime type equal b's runtime type
func IsSameType(a interface{}, b interface{}) bool {
	ra := reflect.TypeOf(a)
	pa := false
	if ra.Kind() == reflect.Ptr {
		ra = ra.Elem()
		pa = true
	}
	rb := reflect.TypeOf(b)
	pb := false
	if rb.Kind() == reflect.Ptr {
		rb = rb.Elem()
		pb = true
	}
	return pa == pb && ra.Kind() == rb.Kind()
}
