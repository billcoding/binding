package binding

var (
	Header = &Type{"header"}
	Param  = &Type{"param"}
	Body   = &Type{"body"}
)

//Define btype struct
type Type struct {
	t string
}
