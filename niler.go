package niler

import (
	"reflect"
)

var nillables = []reflect.Kind{
	reflect.Chan,
	reflect.Func,
	reflect.Interface,
	reflect.Map,
	reflect.Ptr,
	reflect.Slice,
}

func IsNil(v any) bool {

	// Untyped nil?
	if v == nil {
		return true
	}

	// Typed pointer?
	value := reflect.ValueOf(v)
	kind := value.Kind()
	for _, v := range nillables {
		if v == kind {
			return value.IsNil()
		}
	}

	// Has zero value
	return false
}
