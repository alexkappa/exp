package exp

import "testing"

func TestNumbers(t *testing.T) {
	var p = Map{
		"foo": "23",
		"bar": "5",
	}
	for _, test := range []struct {
		exp Exp
		out bool
	}{
		{Eq("foo", 23.00), true},
		{Gt("bar", 4.9), true},
		{Gt("bar", 5.1), false},
		{Lt("bar", 6.2), true},
		{Lt("bar", 4.2), false},
		{Neq("bar", 6), true},
		{Gte("bar", 4.9), true},
		{Gte("bar", 5), true},
		{Gte("bar", 6), false},
		{Lte("foo", 23), true},
		{Lte("foo", 22), false},
		{Lte("foo", 24), true},
	} {
		if test.exp.Eval(p) != test.out {
			t.Errorf("%s should evaluate to %t.", test.exp, test.out)
		}
	}
}
