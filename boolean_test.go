package tree

import "testing"

type True struct{}

func (t True) Eval(p map[string]string) bool { return true }

type False struct{}

func (f False) Eval(p map[string]string) bool { return false }

func TestAndTrue(t *testing.T) {
	tree := And{True{}, True{}}
	if !tree.Eval(nil) {
		t.Error("T AND T should evaluate to T")
	}
}

func TestAndFalse(t *testing.T) {
	tree := And{True{}, False{}}
	if tree.Eval(nil) {
		t.Error("T AND F should evaluate to F")
	}
}

func TestOrTrue(t *testing.T) {
	tree := Or{True{}, False{}}
	if !tree.Eval(nil) {
		t.Error("T OR F should evaluate to T")
	}
}

func TestOrFalse(t *testing.T) {
	tree := Or{False{}, False{}}
	if tree.Eval(nil) {
		t.Error("F OR F should evaluate to F")
	}
}
