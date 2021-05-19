package binding

var (
	// Header Type
	Header = &Type{"header"}
	// Param Type
	Param = &Type{"param"}
)

// Type struct
type Type struct {
	t string
}
