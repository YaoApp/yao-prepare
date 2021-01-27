package json

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TylerBrock/colorjson"
	jsoniter "github.com/json-iterator/go"
	"github.com/yaoapp/yao/lib/exception"
)

// M Alias map[string]interface{}
type M map[string]interface{}

// Decode Decode the JSON format text
func Decode(content string, v interface{}) {
	DecodeBytes([]byte(content), v)
}

// DecodeBytes Decode the JSON format by
func DecodeBytes(content []byte, v interface{}) {
	jsoniter.Unmarshal(content, v)
}

// Encode Encode to the JSON format text
func Encode(v interface{}) string {
	jsonbytes := EncodeBytes(v)
	return string(jsonbytes)
}

// EncodeBytes Encode to the JSON format text
func EncodeBytes(v interface{}) []byte {
	jsonbytes, err := jsoniter.Marshal(v)
	if err != nil {
		exception.Error(err, 500).
			Ctx(M{"v": v}).
			Throw()
	}
	return jsonbytes
}

// DecodeFile Decode the JSON format text file
func DecodeFile(file string, v interface{}) {

	if _, err := os.Stat(file); os.IsNotExist(err) {
		exception.New("file does not exist.", 400).
			Ctx(M{"filepath": file}).
			Throw()
	}

	jsonbytes, err := ioutil.ReadFile(file)
	if err != nil {
		exception.Error(err, 500).
			Ctx(M{"filepath": file}).
			Throw()
	}
	DecodeBytes(jsonbytes, v)
}

// SaveFile Encode the JSON format text and save it as a file
func SaveFile(v interface{}, file string) {
	jsonbytes := EncodeBytes(v)
	err := ioutil.WriteFile(file, jsonbytes, 0644)
	if err != nil {
		exception.Error(err, 500).
			Ctx(M{"filepath": file}).
			Throw()
	}
}

// PrettyPrint pretty print var
func PrettyPrint(v interface{}) {
	f := colorjson.NewFormatter()
	f.Indent = 4
	var res map[string]interface{}
	txt, _ := jsoniter.Marshal(v)
	jsoniter.Unmarshal(txt, &res)
	s, _ := f.Marshal(res)
	fmt.Printf("%s\n", s)
}
