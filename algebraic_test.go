package tree

import "testing"

var p = Params{
	"foo": "bar",
	"bar": "baz",
}

func TestEq(t *testing.T) {
	op := Eq{"foo", "bar"}
	if !op.Eval(p) {
		t.Error("parameter foo should be equal bar")
	}
}

func TestGt(t *testing.T) {
	op := Gt{"bar", "aaa"}
	if !op.Eval(p) {
		t.Error("parameter bar should be greater than aaa")
	}
}

func TestLt(t *testing.T) {
	op := Lt{"bar", "xxx"}
	if !op.Eval(p) {
		t.Error("parameter bar should be less than xxx")
	}
}
