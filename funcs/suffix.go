package funcs

import (
	"github.com/billcoding/calls"
	"reflect"
)

// suffixFunc struct
type suffixFunc struct {
	suffix string
}

// SuffixFunc method
func SuffixFunc(suffix string) BFunc {
	return &suffixFunc{suffix}
}

// Bind method
func (s *suffixFunc) Bind(inValue reflect.Value) (outValue reflect.Value) {
	outValue = inValue
	calls.True(s.suffix != "" && inValue.IsValid() && inValue.Type().Kind() == reflect.String, func() {
		outValue = reflect.ValueOf(inValue.String() + s.suffix)
	})
	return outValue
}
