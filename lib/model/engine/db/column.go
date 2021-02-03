package db

import (
	"github.com/yaoapp/yao/lib/arr"
	"github.com/yaoapp/yao/lib/t"
)

func columnLengthDefault(typ string) *int {
	var value int = 200
	switch typ {
	case "char", "string":
		value = 200
		break
	case "bigInteger":
		value = 20
		break
	case "decimal", "double", "float", "unsignedDecimal":
		value = 10
		break
	}
	return &value
}

func columnLength(typ string, args interface{}) *int {
	if arr.Have([]string{"bigInteger", "char", "string"}, typ) {
		return t.Int(args)
	} else if arr.Have([]string{"decimal", "double", "float", "unsignedDecimal"}, typ) {
		switch args.(type) {
		case []int:
			if len(args.([]int)) >= 1 {
				return t.Int(args.([]int)[0])
			}
		case []interface{}:
			if len(args.([]interface{})) >= 1 {
				return t.Int(args.([]interface{})[0])
			}
		case int, interface{}:
			return t.Int(args)
		}
	}
	return nil
}

func (column *Column) comment(comment string) {
}

func (column *Column) defaults(args interface{}) {
}

func (column *Column) first() {
}

func (column *Column) after() {
}

func (column *Column) nullable() {
}

func (column *Column) unsigned() {
}

func (column *Column) bigInteger(args interface{}) {
}

func (column *Column) binary(args interface{}) {
}

func (column *Column) boolean(args interface{}) {
}

func (column *Column) string(args interface{}) {
}
