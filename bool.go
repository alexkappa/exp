package exp

// Boolean

// Bool is a wrapper for the native bool type which provides an Eval function so
// that it satisfies the Exp interface.
type Bool bool

// Eval will always return the boolean value of b and disregard p.
func (b Bool) Eval(p Params) bool {
	return bool(b)
}

var (
	// True is an expression that always evaluates to true.
	True = Bool(true)
	// False is an expression that always evaluates to false.
	False = Bool(false)
)
