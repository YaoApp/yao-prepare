//Package maps the map type variable utils
package maps

import (
	"fmt"
	"reflect"

	"github.com/yaoapp/yao/lib/exception"
	"github.com/yaoapp/yao/lib/t"
)

// Keys get the given map[string]... type variable keys
func Keys(m interface{}) []string {
	keys := KeysAny(m)
	keysString := []string{}
	for _, key := range keys {
		keysString = append(keysString, fmt.Sprintf("%s", key))
	}
	return keysString
}

// KeysAny get the given map type variable keys
func KeysAny(m interface{}) []interface{} {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		exception.New("given variable must be a map[string]...", 400).
			Ctx(t.M{"var": m, "kind": val.Kind()}).
			Throw()
	}
	iter := val.MapRange()
	keys := []interface{}{}
	for iter.Next() {
		key := iter.Key()
		keys = append(keys, key)
	}
	return keys
}
