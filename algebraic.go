package tree

type Eq struct{ Key, Value string }

func (eq Eq) Eval(p Params) bool {
	if val, found := p[eq.Key]; found {
		return val == eq.Value
	}
	return false
}

type Gt struct{ Key, Value string }

func (gt Gt) Eval(p Params) bool {
	if val, found := p[gt.Key]; found {
		return val > gt.Value
	}
	return false
}

type Lt struct{ Key, Value string }

func (lt Lt) Eval(p Params) bool {
	if val, found := p[lt.Key]; found {
		return val < lt.Value
	}
	return false
}
