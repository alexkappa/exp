package exp

import "testing"

func TestParse(t *testing.T) {
	m := Map{
		"foo": "124",
		"bar": "x",
		"b-z": "z",
	}
	for _, s := range []string{
		`true`,
		`(foo == 124)`,
		`(foo >= 123)`,
		`(foo < 456)`,
		`((foo > 200) || (bar == "x"))`,
		`((foo > 100) && (bar == "x"))`,
		`('b-z' == "z")`,
	} {
		exp, err := Parse(s)
		if err != nil {
			t.Fatal(err)
		}
		if success := exp.Eval(m); !success {
			t.Errorf("unexpected output %t", success)
		}
		t.Logf("%s", exp)
	}
}
