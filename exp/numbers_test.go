package exp

import "testing"

var p = Map{
	"foo": "23",
	"bar": "5",
}

func TestEq(t *testing.T) {
	if !Eq("foo", "23").Eval(p) {
		t.Error("parameter foo should be equal to 23")
	}
}

func TestGt(t *testing.T) {
	if !Gt("bar", "4").Eval(p) {
		t.Error("parameter bar should be greater than 6")
	}
}

func TestLt(t *testing.T) {
	if !Lt("bar", "6").Eval(p) {
		t.Error("parameter bar should be less than 4")
	}
}

func TestNeq(t *testing.T) {
	if !NotEqual("bar", "6").Eval(p) {
		t.Error("parameter bar should not be equal to 6")
	}
}

func TestGte(t *testing.T) {
	if !Gte("bar", "5").Eval(p) {
		t.Error("parameter bar should be greater than or equal to 5")
	}
}

func TestLte(t *testing.T) {
	if !Lte("foo", "23").Eval(p) {
		t.Error("parameter foo should be less than or equal to 23")
	}
}
