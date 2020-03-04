package xmap

import (
	"reflect"
)

func Keys(m interface{}) []interface{} {
	rtype := reflect.ValueOf(m)
	if rtype.Kind() == reflect.Ptr {
		rtype = rtype.Elem()
	}
	keys := make([]interface{}, 0)
	switch rtype.Kind() {
	case reflect.Map:
		values := rtype.MapKeys()
		for _, v := range values {
			keys = append(keys, v.Interface())
		}
	case reflect.Struct:
		for i := 0; i < rtype.NumField(); i++ {
			keys = append(keys, rtype.Type().Field(i).Name)
		}
	}
	return keys
}

func Values(v interface{}) []interface{} {
	rtype := reflect.ValueOf(v)
	if rtype.Kind() == reflect.Ptr {
		rtype = rtype.Elem()
	}
	keys := make([]interface{}, 0)
	switch rtype.Kind() {
	case reflect.Map:
		r := rtype.MapRange()
		for r.Next() {
			keys = append(keys, r.Value().Interface())
		}
	case reflect.Struct:
		for i := 0; i < rtype.NumField(); i++ {
			keys = append(keys, rtype.Field(i).Interface())
		}
	}
	return keys
}

func Contains(v interface{}, x interface{}) bool {
	rtype := reflect.ValueOf(v)
	if rtype.Kind() == reflect.Ptr {
		rtype = rtype.Elem()
	}
	switch rtype.Kind() {
	case reflect.Map:
		for _, t := range rtype.MapKeys() {
			if t.Interface() == x {
				return true
			}
		}
	case reflect.Struct:
		str, ok := x.(string)
		if !ok {
			return false
		}
		for i := 0; i < rtype.NumField(); i++ {
			if rtype.Type().Field(i).Name == str {
				return true
			}
		}
	}
	return false
}
