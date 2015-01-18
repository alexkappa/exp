package exp

import "testing"

func TestExp(t *testing.T) {
	for _, test := range []struct {
		exp Exp
		out bool
	}{
		{Not(True), false},
		{Not(False), true},
		{And(True, True), true},
		{And(True, False), false},
		{And(False, False), false},
		{And(True, True, True), true},
		{And(True, False, True), false},
		{Or(True, True), true},
		{Or(True, False), true},
		{Or(False, False), false},
		{Or(False, False, True), true},
	} {
		if test.exp.Eval(nil) != test.out {
			t.Error("unexpected output")
		}
	}
}
