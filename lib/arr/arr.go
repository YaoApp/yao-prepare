package arr

import (
	"reflect"

	"github.com/yaoapp/yao/lib/exception"
	"github.com/yaoapp/yao/lib/t"
)

// Pluck pluck a list of the given key / value pairs from the array/slice.
func Pluck(array interface{}, find interface{}, v interface{}) {
	values := reflect.ValueOf(array)
	if values.Kind() != reflect.Array && values.Kind() != reflect.Slice {
		exception.New("Argument 1 must be an Array or Slice", 500).
			Ctx(t.M{"array": array, "kind": values.Kind()}).
			Throw()
	}

	method := reflect.ValueOf(find)
	if method.Kind() != reflect.Func {
		exception.New("Argument 2 must be a Function", 500).
			Ctx(t.M{"kind": method.Kind()}).
			Throw()
	}
	methodType := reflect.TypeOf(find)
	if methodType.NumIn() != 1 {
		exception.New("the Function must have one return value", 500).
			Ctx(t.M{"numIn": methodType.NumIn()}).
			Throw()
	}
	if methodType.NumOut() != 1 {
		exception.New("the Function must have one return value", 500).
			Ctx(t.M{"numIn": methodType.NumOut()}).
			Throw()
	}

	result := reflect.ValueOf(v)
	if result.Kind() != reflect.Ptr {
		exception.New("Argument 3 must be an Pointer", 500).
			Ctx(t.M{"v": v, "kind": values.Kind()}).
			Throw()
	}

	length := values.Len()
	typ := methodType.Out(0)
	slice := reflect.MakeSlice(reflect.SliceOf(typ), 0, length)
	for i := 0; i < length; i++ {
		elm := values.Index(i)
		vals := method.Call([]reflect.Value{elm})
		slice = reflect.Append(slice, vals[0])
	}

	result.Elem().Set(slice)
}

// Have Check element exists in a array/slice
func Have(array interface{}, item interface{}) bool {
	MustBeArray(array)
	values := reflect.ValueOf(array)
	length := values.Len()
	for i := 0; i < length; i++ {
		elm := values.Index(i)
		if elm.Interface() == item {
			return true
		}

	}
	return false
}

// HaveMany Check each element exists in a array/slice
func HaveMany(array interface{}, items interface{}) bool {
	MustBeArray(items)
	values := reflect.ValueOf(items)
	length := values.Len()
	for i := 0; i < length; i++ {
		item := values.Index(i)
		if !Have(array, item.Interface()) {
			return false
		}
	}
	return true
}

// MustBeArray the given variable must be array or slice
func MustBeArray(array interface{}) {
	values := reflect.ValueOf(array)
	if values.Kind() != reflect.Array && values.Kind() != reflect.Slice {
		exception.New("the variable type is not Array or Slice", 500).
			Ctx(t.M{"array": array, "kind": values.Kind()}).
			Throw()
	}
}
