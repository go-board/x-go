package types

import (
	"reflect"
)

type TypeId string

func (t TypeId) String() string { return string(t) }

func typeid(value interface{}) string {
	rt := reflect.TypeOf(value)
	name := rt.String()

	// But for named types (or pointers to them), qualify with import path (but see inner comment).
	// Dereference one pointer looking for a named type.
	star := ""
	if rt.Name() == "" {
		if pt := rt; pt.Kind() == reflect.Ptr {
			star = "*"
			// NOTE: The following line should be rt = pt.Elem() to implement
			// what the comment above claims, but fixing it would break compatibility
			// with existing gobs.
			//
			// Given package p imported as "full/p" with these definitions:
			//     package p
			//     type T1 struct { ... }
			// this table shows the intended and actual strings used by gob to
			// name the types:
			//
			// Type      Correct string     Actual string
			//
			// T1        full/p.T1          full/p.T1
			// *T1       *full/p.T1         *full/p.T1
			//
			//rt = pt
			rt = pt.Elem()
		}
	}
	if rt.Name() != "" {
		if rt.PkgPath() == "" {
			name = star + rt.Name()
		} else {
			name = star + rt.PkgPath() + "." + rt.Name()
		}
	}
	return name
}

type typeof struct{}

func (typeof) Typeof(value interface{}) TypeId { return TypeId(typeid(value)) }

var Type = typeof{}

func Typeof(value interface{}) TypeId { return Type.Typeof(value) }
