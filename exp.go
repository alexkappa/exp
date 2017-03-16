// Package exp implements a binary expression tree which can be used to evaluate
// arbitrary binary expressions. You can use this package to build your own
// expressions however a few expressions are provided out of the box.
package exp

import (
	"errors"
	"strconv"

	"github.com/alexkappa/exp/parse"
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
type Params interface {
	Get(string) string
}

// Map is a simple implementation of Params using a map of strings.
type Map map[string]string

// Get returns the value pointed to by key.
func (m Map) Get(key string) string {
	return m[key]
}

func Parse(s string) (Exp, error) {
	t, err := parse.Parse(s)
	if err != nil {
		return nil, err
	}
	return visit(t)
}

func visit(t parse.Tree) (Exp, error) {
	token := t.Value()
	switch token.Type {
	case parse.T_BOOLEAN:
		switch token.Value {
		case "true":
			return True, nil
		case "false":
			return False, nil
		}
	case parse.T_LOGICAL_AND:
		l, err := visit(t.Left())
		if err != nil {
			return nil, err
		}
		r, err := visit(t.Right())
		if err != nil {
			return nil, err
		}
		return And(l, r), nil
	case parse.T_LOGICAL_OR:
		l, err := visit(t.Left())
		if err != nil {
			return nil, err
		}
		r, err := visit(t.Right())
		if err != nil {
			return nil, err
		}
		return Or(l, r), nil
	case parse.T_IS_EQUAL:
		var (
			k, v string
			f    float64
			l    = t.Left()
			r    = t.Right()
		)
		switch l.Value().Type {
		case parse.T_IDENTIFIER:
			k = l.Value().Value
		case parse.T_STRING:
			v = l.Value().Value
		case parse.T_NUMBER:
			var err error
			f, err = strconv.ParseFloat(l.Value().Value, 10)
			if err != nil {
				return nil, err
			}
		}
		switch r.Value().Type {
		case parse.T_IDENTIFIER:
			k = r.Value().Value
		case parse.T_STRING:
			v = r.Value().Value
		case parse.T_NUMBER:
			var err error
			f, err = strconv.ParseFloat(r.Value().Value, 10)
			if err != nil {
				return nil, err
			}
		}
		if v == "" {
			return Equal(k, f), nil
		}
		return Match(k, v), nil
	case parse.T_IS_NOT_EQUAL:
		var (
			k, v string
			f    float64
			l    = t.Left()
			r    = t.Right()
		)
		switch l.Value().Type {
		case parse.T_IDENTIFIER:
			k = l.Value().Value
		case parse.T_STRING:
			v = l.Value().Value
		case parse.T_NUMBER:
			var err error
			f, err = strconv.ParseFloat(l.Value().Value, 10)
			if err != nil {
				return nil, err
			}
		}
		switch r.Value().Type {
		case parse.T_IDENTIFIER:
			k = r.Value().Value
		case parse.T_STRING:
			v = r.Value().Value
		case parse.T_NUMBER:
			var err error
			f, err = strconv.ParseFloat(r.Value().Value, 10)
			if err != nil {
				return nil, err
			}
		}
		if v == "" {
			return NotEqual(k, f), nil
		}
		return Not(Match(k, v)), nil
	case parse.T_IS_GREATER:
		var (
			k string
			v float64
			l = t.Left()
			r = t.Right()
		)
		switch l.Value().Type {
		case parse.T_IDENTIFIER:
			k = l.Value().Value
		case parse.T_NUMBER:
			var err error
			v, err = strconv.ParseFloat(l.Value().Value, 10)
			if err != nil {
				return nil, err
			}
		}
		switch r.Value().Type {
		case parse.T_IDENTIFIER:
			k = r.Value().Value
		case parse.T_NUMBER:
			var err error
			v, err = strconv.ParseFloat(r.Value().Value, 10)
			if err != nil {
				return nil, err
			}
		}
		return GreaterThan(k, v), nil
	case parse.T_IS_GREATER_OR_EQUAL:
		var (
			k string
			v float64
			l = t.Left()
			r = t.Right()
		)
		switch l.Value().Type {
		case parse.T_IDENTIFIER:
			k = l.Value().Value
		case parse.T_NUMBER:
			var err error
			v, err = strconv.ParseFloat(l.Value().Value, 10)
			if err != nil {
				return nil, err
			}
		}
		switch r.Value().Type {
		case parse.T_IDENTIFIER:
			k = r.Value().Value
		case parse.T_NUMBER:
			var err error
			v, err = strconv.ParseFloat(r.Value().Value, 10)
			if err != nil {
				return nil, err
			}
		}
		return GreaterOrEqual(k, v), nil
	case parse.T_IS_SMALLER:
		var (
			k string
			v float64
			l = t.Left()
			r = t.Right()
		)
		switch l.Value().Type {
		case parse.T_IDENTIFIER:
			k = l.Value().Value
		case parse.T_NUMBER:
			var err error
			v, err = strconv.ParseFloat(l.Value().Value, 10)
			if err != nil {
				return nil, err
			}
		}
		switch r.Value().Type {
		case parse.T_IDENTIFIER:
			k = r.Value().Value
		case parse.T_NUMBER:
			var err error
			v, err = strconv.ParseFloat(r.Value().Value, 10)
			if err != nil {
				return nil, err
			}
		}
		return LessThan(k, v), nil
	case parse.T_IS_SMALLER_OR_EQUAL:
		var (
			k string
			v float64
			l = t.Left()
			r = t.Right()
		)
		switch l.Value().Type {
		case parse.T_IDENTIFIER:
			k = l.Value().Value
		case parse.T_NUMBER:
			var err error
			v, err = strconv.ParseFloat(l.Value().Value, 10)
			if err != nil {
				return nil, err
			}
		}
		switch r.Value().Type {
		case parse.T_IDENTIFIER:
			k = r.Value().Value
		case parse.T_NUMBER:
			var err error
			v, err = strconv.ParseFloat(r.Value().Value, 10)
			if err != nil {
				return nil, err
			}
		}
		return LessOrEqual(k, v), nil
	}

	return nil, errors.New("invalid expression")
}

// And

type expAnd struct{ elems []Exp }

func (a expAnd) Eval(p Params) bool {
	for _, elem := range a.elems {
		if !elem.Eval(p) {
			return false
		}
	}
	return true
}

func (a expAnd) String() string {
	return sprintf("(%s)", join(a.elems, "∧"))
}

// And evaluates to true if all t's are true.
func And(t ...Exp) Exp {
	return expAnd{t}
}

// Or

type expOr struct{ elems []Exp }

func (o expOr) Eval(p Params) bool {
	for _, elem := range o.elems {
		if elem.Eval(p) {
			return true
		}
	}
	return false
}

func (o expOr) String() string {
	return sprintf("(%s)", join(o.elems, "∨"))
}

// Or evaluates to true if any t's are true.
func Or(t ...Exp) Exp {
	return expOr{t}
}

// Not

type expNot struct{ elem Exp }

func (n expNot) Eval(p Params) bool {
	return !n.elem.Eval(p)
}

func (n expNot) String() string {
	return sprintf("¬%s", n.elem)
}

// Not evaluates to the opposite of t.
func Not(t Exp) Exp {
	return expNot{t}
}
