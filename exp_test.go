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

func TestParse(t *testing.T) {
	m := Map{
		"foo": "124",
		"bar": "x",
		"b-z": "z",
	}
	for _, s := range []string{
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
		if !exp.Eval(m) {
			t.Error("unexpected output")
		}
		t.Logf("%s", exp)
	}
}

func TestTimestampParse(t *testing.T) {
	m := Map{
		"timestamp": "2020-04-07T20:05:00+05:30",
	}
	for _, s := range []string{
		`(timestamp == "2020-04-07T20:05:00+05:30")`,
		`(timestamp > "2020-04-07T20:04:00+05:30")`,
		`(timestamp < "2020-04-07T20:06:00+05:30")`,
		`(timestamp <= "2020-04-07T20:05:00+05:30")`,
		`(timestamp >= "2020-04-07T20:05:00+05:30")`,
	} {
		exp, err := Parse(s)
		if err != nil {
			t.Fatal(err)
		}
		if !exp.Eval(m) {
			t.Error("unexpected output")
		}
		t.Logf("%s", exp)
	}
}
