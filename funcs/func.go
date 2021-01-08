package funcs

import "reflect"

// BFunc interface
type BFunc interface {
	// Bind call
	Bind(inValue reflect.Value) (outValue reflect.Value)
}
