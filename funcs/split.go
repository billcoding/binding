package funcs

import (
	"reflect"
	"strings"
)

// splitFunc struct
type splitFunc struct {
	split   bool
	splitsp string
}

// SplitFunc method
func SplitFunc(split bool, splitsp string) BFunc {
	if splitsp == "" {
		splitsp = ","
	}
	return &splitFunc{
		split:   split,
		splitsp: splitsp,
	}
}

// Bind method
func (s *splitFunc) Bind(inValue reflect.Value) (outValue reflect.Value) {
	outValue = inValue
	if s.split && inValue.IsValid() && inValue.Type().Kind() == reflect.String {
		outValue = reflect.ValueOf(strings.Split(inValue.String(), s.splitsp))
	}
	return outValue
}
