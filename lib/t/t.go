package t

import (
	"fmt"
	"strconv"

	"github.com/yaoapp/yao/lib/exception"
)

// M Alias map[string]interface{}
type M map[string]interface{}

// Int64 covering any to *int64
func Int64(v interface{}) *int64 {
	switch v.(type) {
	case string, int32, int16, int8, int, uint32, uint16, uint8, uint, float32, float64:
		val := fmt.Sprintf("%v", v)
		res, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			exception.Err(err, 500).Ctx(M{"v": v}).Throw()
		}
		return &res
	case int64, uint64:
		res := v.(int64)
		return &res
	case bool:
		val := v.(bool)
		var res int64 = 0
		if val {
			res = 1
		}
		return &res
	}
	return nil
}

// Int covering any to *int
func Int(v interface{}) *int {
	switch v.(type) {
	case string, int32, int16, int8, int64, uint32, uint16, uint8, uint64, float32, float64:
		val := fmt.Sprintf("%v", v)
		res, err := strconv.Atoi(val)
		if err != nil {
			exception.Err(err, 500).Ctx(M{"v": v}).Throw()
		}
		return &res
	case int, uint:
		res := v.(int)
		return &res
	case bool:
		val := v.(bool)
		var res int = 0
		if val {
			res = 1
		}
		return &res
	}
	return nil
}

// Bool covering any to *bool
func Bool(v interface{}) *bool {
	switch v.(type) {
	case bool:
		val := v.(bool)
		return &val
	case int, uint, int32, int16, int8, int64, uint32, uint16, uint8, uint64, float32, float64:
		val, err := strconv.Atoi(fmt.Sprintf("%v", v))
		if err != nil {
			exception.Err(err, 500).Ctx(M{"v": v}).Throw()
		}
		res := false
		if val != 0 {
			res = true
		}
		return &res
	default:
		val := fmt.Sprintf("%v", v)
		res := false
		if val != "" {
			res = true
		}
		return &res
	}
}
