// Package exp implements a binary expression tree which can be used to evaluate
// arbitrary binary expressions. You can use this package to build your own
// expressions however a few expressions are provided out of the box.
package exp

import (
	"strconv"
	"strings"
	"time"
)

// The Exp interface represents a tree node. There are several implementations
// of the interface in this package, but one may define custom Exp's as long as
// they implement the Eval function.
type Exp interface {
	Eval(Params) bool
}

// Params defines the interface needed by Exp in order to be able to validate
// conditions. An example implementation of this interface would be
// https://golang.org/pkg/net/url/#Values.
//
// A simple implementation of Params can be described as a map of strings.
//
// 	type Params map[string]string
//
// 	func (p Params) Get(s string) string {
// 		return p[s]
// 	}
type Params interface {
	Get(string) string
}

// And -------------------------------------------------------------------------

type expAnd struct{ elems []Exp }

func (a expAnd) Eval(p Params) bool {
	for _, elem := range a.elems {
		if !elem.Eval(p) {
			return false
		}
	}
	return true
}

// Or --------------------------------------------------------------------------

type expOr struct{ elems []Exp }

func (o expOr) Eval(p Params) bool {
	for _, elem := range o.elems {
		if elem.Eval(p) {
			return true
		}
	}
	return false
}

// Not -------------------------------------------------------------------------

type expNot struct{ elem Exp }

func (n expNot) Eval(p Params) bool {
	return !n.elem.Eval(p)
}

// True / False ----------------------------------------------------------------

type expBool bool

func (b expBool) Eval(p Params) bool { return bool(b) }

// Eq --------------------------------------------------------------------------

type expEq struct{ Key, Value string }

func (eq expEq) Eval(p Params) bool { return p.Get(eq.Key) == eq.Value }

// Gt --------------------------------------------------------------------------

type expGt struct{ Key, Value string }

func (gt expGt) Eval(p Params) bool {
	ia, ea := strconv.ParseFloat(p.Get(gt.Key), 32)
	ib, eb := strconv.ParseFloat(gt.Value, 32)
	if ea != nil || eb != nil {
		return false
	}
	return ia > ib
}

// Lt --------------------------------------------------------------------------

type expLt struct{ Key, Value string }

func (lt expLt) Eval(p Params) bool {
	ia, ea := strconv.ParseFloat(p.Get(lt.Key), 32)
	ib, eb := strconv.ParseFloat(lt.Value, 32)
	if ea != nil || eb != nil {
		return false
	}
	return ia < ib
}

// Like ------------------------------------------------------------------------

type expLike struct{ Key, Value string }

func (l expLike) Eval(p Params) bool {
	return strings.Contains(p.Get(l.Key), l.Value)
}

// Before ----------------------------------------------------------------------

type expBefore struct{ Key, Value string }

func (b expBefore) Eval(p Params) bool {
	ta, ea := time.Parse(dateFormat, p.Get(b.Key))
	tb, eb := time.Parse(dateFormat, b.Value)
	if ea != nil || eb != nil {
		return false
	}
	return ta.Before(tb)
}

// After -----------------------------------------------------------------------

type expAfter struct{ Key, Value string }

func (a expAfter) Eval(p Params) bool {
	ta, ea := time.Parse(dateFormat, p.Get(a.Key))
	tb, eb := time.Parse(dateFormat, a.Value)
	if ea != nil || eb != nil {
		return false
	}
	return ta.After(tb)
}

// Public API ------------------------------------------------------------------

// And evaluates to true if all t's are true.
func And(t ...Exp) Exp { return expAnd{t} }

// Or evaluates to true if any t's are true.
func Or(t ...Exp) Exp { return expOr{t} }

// Not evaluates to the opposite of t.
func Not(t Exp) Exp { return expNot{t} }

// Equals is an expression that evaluates to true if the evaluated key is equal
// in value to v.
func Equals(k, v string) Exp { return expEq{k, v} }

// GreaterThan is an expression that evaluates to true if the evaluated key is
// greater in value than v. The value is parsed as float before performing the
// comparison.
func GreaterThan(k, v string) Exp { return expGt{k, v} }

// LessThan is an expression that evaluates to true if the evaluated key is less
// in value than v. The value is parsed as float before performing the
// comparison.
func LessThan(k, v string) Exp { return expLt{k, v} }

// Like is an expression that evaluates to true if v is contained within the
// value of the evaluated key.
func Like(k, v string) Exp { return expLike{k, v} }

// Before is an expression that evaluates to true if v is a date before the
// evaluated date. The value is parsed to a time.Time before comparing.
func Before(k, v string) Exp { return expBefore{k, v} }

// After is an expression that evaluates to true if v is a date after the
// evaluated date. The value is parsed to a time.Time before comparing.
func After(k, v string) Exp { return expAfter{k, v} }

// Eq is an alias for Equals.
func Eq(k, v string) Exp { return Equals(k, v) }

// Neq is a shorthand for Not(Eq(k, v)).
func Neq(k, v string) Exp { return Not(Eq(k, v)) }

// Gt is an alias for GreaterThan.
func Gt(k, v string) Exp { return GreaterThan(k, v) }

// Lt is an alias for LessThan.
func Lt(k, v string) Exp { return LessThan(k, v) }

// Gte is a shorthand for Or(Gt(k, v), Eq(k, v)).
func Gte(k, v string) Exp { return Or(Gt(k, v), Eq(k, v)) }

// Lte is a shorthand for Lt(Gt(k, v), Eq(k, v)).
func Lte(k, v string) Exp { return Or(Lt(k, v), Eq(k, v)) }

// The default format used to parse dates.
var dateFormat = "2006-01-02"

// DateFormat changes the date format used to parse dates.
func DateFormat(f string) { dateFormat = f }

var (
	// True is an expression that always evaluates to true.
	True expBool = true
	// False is an expression that always evaluates to false.
	False expBool = false
)
