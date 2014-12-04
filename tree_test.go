package tree

import "testing"

func TestTrue(t *testing.T) {
	if !True.Eval(nil) {
		t.Error("T should evaluate to T")
	}
}

func TestFalse(t *testing.T) {
	if False.Eval(nil) {
		t.Error("F should evaluate to F")
	}
}

func TestAndTrue(t *testing.T) {
	exp := And(True, True)
	if !exp.Eval(nil) {
		t.Error("T AND T should evaluate to T")
	}
}

func TestAndFalse(t *testing.T) {
	exp := And(True, False)
	if exp.Eval(nil) {
		t.Error("T AND F should evaluate to F")
	}
}

func TestOrTrue(t *testing.T) {
	exp := Or(True, False)
	if !exp.Eval(nil) {
		t.Error("T OR F should evaluate to T")
	}
}

func TestOrFalse(t *testing.T) {
	exp := Or(False, False)
	if exp.Eval(nil) {
		t.Error("F OR F should evaluate to F")
	}
}

type params map[string]string

func (p params) Get(k string) string { return p[k] }

var p = params{
	"foo": "23",
	"bar": "5",
}

func TestEq(t *testing.T) {
	op := Eq("foo", "23")
	if !op.Eval(p) {
		t.Error("parameter foo should be equal to 23")
	}
}

func TestGt(t *testing.T) {
	op := Gt("bar", "4")
	if !op.Eval(p) {
		t.Error("parameter bar should be greater than 6")
	}
}

func TestLt(t *testing.T) {
	op := Lt("bar", "6")
	if !op.Eval(p) {
		t.Error("parameter bar should be less than 4")
	}
}
