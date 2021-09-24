package binding

import (
	"reflect"
)

// suffixFunc struct
type suffixFunc struct {
	suffix string
}

// SuffixFunc method
func SuffixFunc(suffix string) Func {
	return &suffixFunc{suffix}
}

// Bind method
func (s *suffixFunc) Bind(inValue reflect.Value) (outValue reflect.Value) {
	outValue = inValue
	if s.suffix != "" && inValue.IsValid() && inValue.Type().Kind() == reflect.String {
		outValue = reflect.ValueOf(inValue.String() + s.suffix)
	}
	return outValue
}
