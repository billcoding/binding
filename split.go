package binding

import (
	"reflect"
	"strings"
)

// splitFunc struct
type splitFunc struct {
	split bool
	sp    string
}

// SplitFunc method
func SplitFunc(split bool, sp string) Func {
	if sp == "" {
		sp = ","
	}
	return &splitFunc{
		split: split,
		sp:    sp,
	}
}

// Bind method
func (s *splitFunc) Bind(inValue reflect.Value) (outValue reflect.Value) {
	outValue = inValue
	if s.split && inValue.IsValid() && inValue.Type().Kind() == reflect.String {
		outValue = reflect.ValueOf(strings.Split(inValue.String(), s.sp))
	}
	return outValue
}
