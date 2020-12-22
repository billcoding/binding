package funcs

import "reflect"

//Define BFunc struct
type BFunc interface {
	Bind(inValue reflect.Value) (outValue reflect.Value)
}
