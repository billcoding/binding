package funcs

import (
	"github.com/billcoding/calls"
	"reflect"
	"strings"
)

//Define splitFunc struct
type splitFunc struct {
	split   bool
	splitsp string
}

//SplitFunc
func SplitFunc(split bool, splitsp string) BFunc {
	if splitsp == "" {
		splitsp = ","
	}
	return &splitFunc{
		split:   split,
		splitsp: splitsp,
	}
}

//Bind
func (s *splitFunc) Bind(inValue reflect.Value) (outValue reflect.Value) {
	outValue = inValue
	calls.True(s.split && inValue.IsValid() && inValue.Type().Kind() == reflect.String, func() {
		outValue = reflect.ValueOf(strings.Split(inValue.String(), s.splitsp))
	})
	return outValue
}
