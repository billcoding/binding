package funcs

import (
	"github.com/billcoding/calls"
	"reflect"
)

//Define defaultFunc
type defaultFunc struct {
	defaultVal string
}

//DefaultFunc
func DefaultFunc(defaultVal string) BFunc {
	return &defaultFunc{defaultVal}
}

//Bind
func (d *defaultFunc) Bind(inValue reflect.Value) (outValue reflect.Value) {
	outValue = inValue
	calls.True(!inValue.IsValid(), func() {
		outValue = reflect.ValueOf(d.defaultVal)
	})
	return outValue
}
