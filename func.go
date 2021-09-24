package binding

import "reflect"

// Func interface
type Func interface {
	// Bind call
	Bind(inValue reflect.Value) (outValue reflect.Value)
}
