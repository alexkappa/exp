package exp

import "testing"

func TestBool(t *testing.T) {
	for _, test := range []struct {
		exp Exp
		out bool
	}{
		{True, true},
		{False, false},
	} {
		if test.exp.Eval(nil) != test.out {
			t.Error("unexpected output")
		}
	}
}
