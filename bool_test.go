package exp

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
