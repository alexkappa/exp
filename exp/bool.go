package exp

// Boolean

type Bool bool

func (b Bool) Eval(p Params) bool {
	return bool(b)
}

var (
	// True is an expression that always evaluates to true.
	True = Bool(true)
	// False is an expression that always evaluates to false.
	False = Bool(false)
)
