package funcs

import (
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
	if !inValue.IsValid() {
		outValue = reflect.ValueOf(d.defaultVal)
	}
	return outValue
}
