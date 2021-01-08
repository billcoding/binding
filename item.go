package binding

import (
	"encoding/json"
	"github.com/billcoding/binding/funcs"
	"github.com/billcoding/reflectx"
	"reflect"
)

// Item struct
type Item struct {
	Name    string `alias:"name"`
	Default string `alias:"default"`
	Split   bool   `alias:"split"`
	Splitsp string `alias:"splitsp"`
	Join    bool   `alias:"join"`
	Joinsp  string `alias:"joinsp"`
	Prefix  string `alias:"prefix"`
	Suffix  string `alias:"suffix"`
}

// bfuncs
func (i *Item) bfuncs() []funcs.BFunc {
	return []funcs.BFunc{
		funcs.DefaultFunc(i.Default),
		funcs.SplitFunc(i.Split, i.Splitsp),
		funcs.JoinFunc(i.Join, i.Joinsp),
		funcs.PrefixFunc(i.Prefix),
		funcs.SuffixFunc(i.Suffix),
	}
}

// Bind call
func (i *Item) Bind(field *reflect.StructField, value reflect.Value, dataMap map[string]interface{}) {
	bfuncs := i.bfuncs()
	name := i.Name
	if name == "" {
		name = field.Name
	}
	bindVal := reflect.ValueOf(dataMap[name])
	for _, bFunc := range bfuncs {
		bindVal = bFunc.Bind(bindVal)
	}
	reflectx.SetValue(bindVal, value)
}

// String
func (i *Item) String() string {
	bytes, _ := json.Marshal(i)
	return string(bytes)
}
