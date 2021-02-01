package exception

import (
	"encoding/json"
	"fmt"

	"github.com/TylerBrock/colorjson"
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

// Ctx Add the context for the exception.
func (exception *Exception) Ctx(context interface{}) *Exception {
	exception.Context = context
	return exception
}

// Print print the exception
func (exception *Exception) Print() {
	f := colorjson.NewFormatter()
	f.Indent = 4
	var res interface{}
	txt, _ := json.Marshal(exception)
	json.Unmarshal(txt, &res)
	s, _ := json.Marshal(res)
	fmt.Printf("%s\n", s)
}

// Throw Throw the exception and terminal progress.
func (exception *Exception) Throw() {
	exception.Print()
	panic(exception)
}
