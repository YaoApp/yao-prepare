package exception

import (
	"encoding/json"
	"fmt"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
)

// Exception the Exception type
type Exception struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Context interface{} `json:"context"`
}

// New Create a new exception instance
func New(message string, code int) *Exception {
	return &Exception{
		Message: message,
		Code:    code,
	}
}

// Err Create an exception instance from the error
func Err(err error, code int) *Exception {
	return &Exception{
		Message: err.Error(),
		Code:    code,
	}
}

// CatchPrint Catch the exception and print it
func CatchPrint() {
	if r := recover(); r != nil {
		switch r.(type) {
		case *Exception:
			color.Red(r.(*Exception).Message)
			r.(*Exception).Print()
			break
		case string:
			color.Red(r.(string))
			break
		case error:
			color.Red(r.(error).Error())
			break
		default:
			color.Red("%#v\n", r)
		}
	}
}

// Ctx Add the context for the exception.
func (exception *Exception) Ctx(context interface{}) *Exception {
	exception.Context = context
	return exception
}

// Print print the exception
func (exception *Exception) Print() {
	f := colorjson.NewFormatter()
	f.Indent = 2
	var res interface{}
	txt, _ := json.Marshal(exception)
	json.Unmarshal(txt, &res)
	s, _ := colorjson.Marshal(res)
	fmt.Println(string(s))
}

// Throw Throw the exception and terminal progress.
func (exception *Exception) Throw() {
	panic(exception)
}
