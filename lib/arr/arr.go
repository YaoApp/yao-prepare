package arr

import (
	"reflect"

	"github.com/yaoapp/yao/lib/exception"
	"github.com/yaoapp/yao/lib/json"
)

// Pluck pluck a list of the given key / value pairs from the array.
func Pluck(array interface{}, find interface{}, v interface{}) {
	values := reflect.ValueOf(array)
	if values.Kind() != reflect.Array && values.Kind() != reflect.Slice {
		exception.New("Argument 1 must be an Array or Slice", 500).
			Ctx(json.M{"array": array, "kind": values.Kind()}).
			Throw()
	}

	method := reflect.ValueOf(find)
	if method.Kind() != reflect.Func {
		exception.New("Argument 2 must be a Function", 500).
			Ctx(json.M{"kind": method.Kind()}).
			Throw()
	}
	methodType := reflect.TypeOf(find)
	if methodType.NumIn() != 1 {
		exception.New("the Function must have one return value", 500).
			Ctx(json.M{"numIn": methodType.NumIn()}).
			Throw()
	}
	if methodType.NumOut() != 1 {
		exception.New("the Function must have one return value", 500).
			Ctx(json.M{"numIn": methodType.NumOut()}).
			Throw()
	}

	result := reflect.ValueOf(v)
	if result.Kind() != reflect.Ptr {
		exception.New("Argument 3 must be an Pointer", 500).
			Ctx(json.M{"v": v, "kind": values.Kind()}).
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
