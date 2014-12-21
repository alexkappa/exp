package exp

import "testing"

func TestNotTrue(t *testing.T) {
	if Not(True).Eval(nil) {
		t.Error("NOT T should evaluate to F")
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
