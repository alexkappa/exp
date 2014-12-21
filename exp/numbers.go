package exp

import "strconv"

// Eq

type expEq struct {
	key, value string
}

func (eq expEq) Eval(p Params) bool {
	return p.Get(eq.key) == eq.value
}

// Equal evaluates to true if the value pointed to by key is equal in value to
// v.
func Equal(k, v string) Exp {
	return expEq{k, v}
}

// Eq is an alias for Equal.
func Eq(k, v string) Exp {
	return Equal(k, v)
}

// NotEqual is a shorthand for Not(Eq(k, v)).
func NotEqual(k, v string) Exp {
	return Neq(k, v)
}

// Neq is an alias for NotEqual.
func Neq(k, v string) Exp {
	return Not(Eq(k, v))
}

// Gt

type expGt struct {
	key, value string
}

func (gt expGt) Eval(p Params) bool {
	ia, ea := strconv.ParseFloat(p.Get(gt.key), 32)
	ib, eb := strconv.ParseFloat(gt.value, 32)
	if ea != nil || eb != nil {
		return false
	}
	return ia > ib
}

// GreaterThan evaluates to true if the value pointed to by key is greater in
// value than v. The value is parsed as float before performing the comparison.
func GreaterThan(k, v string) Exp {
	return expGt{k, v}
}

// Gt is an alias for GreaterThan.
func Gt(k, v string) Exp {
	return GreaterThan(k, v)
}

// GreaterOrEqual is a shorthand for Or(Gt(k, v), Eq(k, v)).
func GreaterOrEqual(k, v string) Exp {
	return Or(Gt(k, v), Eq(k, v))
}

// Gte is an alias for GreaterOrEqual.
func Gte(k, v string) Exp {
	return GreaterOrEqual(k, v)
}

// Lt

type expLt struct {
	key, value string
}

func (lt expLt) Eval(p Params) bool {
	ia, ea := strconv.ParseFloat(p.Get(lt.key), 32)
	ib, eb := strconv.ParseFloat(lt.value, 32)
	if ea != nil || eb != nil {
		return false
	}
	return ia < ib
}

// LessThan evaluates to true if the value pointed to by key is less in value
// than v. The value is parsed as float before performing the comparison.
func LessThan(k, v string) Exp {
	return expLt{k, v}
}

// Lt is an alias for LessThan.
func Lt(k, v string) Exp {
	return LessThan(k, v)
}

// LessOrEqual is a shorthand for Lt(Gt(k, v), Eq(k, v)).
func LessOrEqual(k, v string) Exp {
	return Or(Lt(k, v), Eq(k, v))
}

// Lte is an alias for LessOrEqual.
func Lte(k, v string) Exp {
	return LessOrEqual(k, v)
}
