package binding

import (
	"fmt"
	"reflect"
	"strings"
)

// joinFunc struct
type joinFunc struct {
	join bool
	sp   string
}

// JoinFunc method
func JoinFunc(join bool, sp string) Func {
	if sp == "" {
		sp = ","
	}
	return &joinFunc{
		join: join,
		sp:   sp,
	}
}

// Bind method
func (j *joinFunc) Bind(inValue reflect.Value) (outValue reflect.Value) {
	outValue = inValue
	if j.join && inValue.IsValid() && (inValue.Type().Kind() == reflect.Slice || inValue.Type().Kind() == reflect.Array) {
		joins := make([]string, 0)
		for i := 0; i < inValue.Len(); i++ {
			joins = append(joins, fmt.Sprintf("%v", inValue.Index(i)))
		}
		outValue = reflect.ValueOf(strings.Join(joins, j.sp))
	}
	return outValue
}
