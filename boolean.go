package tree

type Node interface {
	Eval(p map[string]string) bool
}

type And struct {
	Left, Right Node
}

func (a And) Eval(p map[string]string) bool {
	return a.Left.Eval(p) && a.Right.Eval(p)
}

type Or struct {
	Left, Right Node
}

func (o Or) Eval(p map[string]string) bool {
	return o.Left.Eval(p) || o.Right.Eval(p)
}

type Not struct {
	Elem Node
}

func (n Not) Eval(p map[string]string) bool {
	return !n.Elem.Eval(p)
}
