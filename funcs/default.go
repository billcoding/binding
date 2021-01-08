package funcs

import (
	"github.com/billcoding/calls"
	"reflect"
)

// defaultFunc struct
type defaultFunc struct {
	defaultVal string
}

// DefaultFunc method
func DefaultFunc(defaultVal string) BFunc {
	return &defaultFunc{defaultVal}
}

// Bind method
func (d *defaultFunc) Bind(inValue reflect.Value) (outValue reflect.Value) {
	outValue = inValue
	calls.True(!inValue.IsValid(), func() {
		outValue = reflect.ValueOf(d.defaultVal)
	})
	return outValue
}
