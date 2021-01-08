package binding

var (
	// Header Type
	Header = &Type{"header"}
	// Param Type
	Param = &Type{"param"}
	// Body Type
	Body = &Type{"body"}
)

// Type struct
type Type struct {
	t string
}
