package binding

import (
	"encoding/json"
	"github.com/billcoding/reflectx"
	"reflect"
)

// Item struct
type Item struct {
	Name    string `alias:"name"`
	Default string `alias:"default"`
	Split   bool   `alias:"split"`
	SplitSp string `alias:"split_sp"`
	Trim    bool   `alias:"trim"`
	TrimSp  string `alias:"trim_sp"`
	Join    bool   `alias:"join"`
	JoinSp  string `alias:"join_sp"`
	Prefix  string `alias:"prefix"`
	Suffix  string `alias:"suffix"`
}

func (i *Item) fs() []Func {
	return []Func{
		DefaultFunc(i.Default),
		TrimFunc(i.Trim, i.TrimSp),
		SplitFunc(i.Split, i.SplitSp),
		JoinFunc(i.Join, i.JoinSp),
		PrefixFunc(i.Prefix),
		SuffixFunc(i.Suffix),
	}
}

// Bind call
func (i *Item) Bind(field *reflect.StructField, value reflect.Value, dataMap map[string]interface{}) {
	fs := i.fs()
	name := i.Name
	if name == "" {
		name = field.Name
	}
	bindVal := reflect.ValueOf(dataMap[name])
	for _, bFunc := range fs {
		bindVal = bFunc.Bind(bindVal)
	}
	reflectx.SetValue(bindVal, value)
}

// String
func (i *Item) String() string {
	bytes, _ := json.Marshal(i)
	return string(bytes)
}
