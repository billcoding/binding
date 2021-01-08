package funcs

import (
	"github.com/billcoding/calls"
	"reflect"
)

// prefixFunc struct
type prefixFunc struct {
	prefix string
}

// PrefixFunc method
func PrefixFunc(prefix string) BFunc {
	return &prefixFunc{prefix}
}

// Bind method
func (p *prefixFunc) Bind(inValue reflect.Value) (outValue reflect.Value) {
	outValue = inValue
	calls.True(p.prefix != "" && inValue.IsValid() && inValue.Type().Kind() == reflect.String, func() {
		outValue = reflect.ValueOf(p.prefix + inValue.String())
	})
	return outValue
}
