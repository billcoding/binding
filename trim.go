package binding

import (
	"reflect"
	"strings"
)

// trimFunc struct
type trimFunc struct {
	trim bool
	sp   string
}

// TrimFunc method
func TrimFunc(trim bool, sp string) Func {
	if sp == "" {
		sp = " "
	}
	return &trimFunc{
		trim: trim,
		sp:   sp,
	}
}

// Bind method
func (t *trimFunc) Bind(inValue reflect.Value) (outValue reflect.Value) {
	outValue = inValue
	if t.trim && inValue.IsValid() && inValue.Type().Kind() == reflect.String {
		outValue = reflect.ValueOf(strings.Trim(inValue.String(), t.sp))
	}
	return outValue
}
