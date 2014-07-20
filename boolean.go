package tree

type True struct{}

func (t True) Eval(p Params) bool {
	return true
}

type False struct{}

func (f False) Eval(p Params) bool {
	return false
}

type And struct {
	Left, Right Node
}

func (a And) Eval(p Params) bool {
	return a.Left.Eval(p) && a.Right.Eval(p)
}

type Or struct {
	Left, Right Node
}

func (o Or) Eval(p Params) bool {
	return o.Left.Eval(p) || o.Right.Eval(p)
}

type Not struct {
	Elem Node
}

func (n Not) Eval(p Params) bool {
	return !n.Elem.Eval(p)
}
