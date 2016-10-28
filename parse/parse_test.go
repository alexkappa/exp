// Copyright (c) 2014 Alex Kalyvitis

package parse

import "testing"

func TestParser(t *testing.T) {
	for _, test := range []struct {
		exp string
		ast *tree
	}{
		{
			"(foo > bar)",
			&tree{
				value: token{Type: T_IS_GREATER, Value: ">"},
				left: &tree{
					value: token{Type: T_IDENTIFIER, Value: "foo"},
				},
				right: &tree{
					value: token{Type: T_IDENTIFIER, Value: "bar"},
				},
			},
		},
		{
			"((foo > bar) && true)",
			&tree{
				value: token{Type: T_LOGICAL_AND, Value: "&&"},
				left: &tree{
					value: token{Type: T_IS_GREATER, Value: ">"},
					left:  &tree{value: token{Type: T_IDENTIFIER, Value: "foo"}},
					right: &tree{value: token{Type: T_IDENTIFIER, Value: "bar"}},
				},
				right: &tree{value: token{Type: T_BOOLEAN, Value: "true"}},
			},
		},
		{
			"(foo > bar)",
			&tree{
				value: token{Type: T_IS_GREATER, Value: ">"},
				left:  &tree{value: token{Type: T_IDENTIFIER, Value: "foo"}},
				right: &tree{value: token{Type: T_IDENTIFIER, Value: "bar"}},
			},
		},
	} {
		ast, err := newParser(newLexer(test.exp)).parse()
		if err != nil {
			t.Fatal(err)
		}

		assertTreeEqual(t, test.ast, ast)
	}
}

func assertTreeEqual(t *testing.T, a, b *tree) {
	t.Logf("\n\ta: %s\n\tb: %s", a, b)

	if a == nil && b == nil {
		return
	}

	if a == nil || b == nil {
		t.Fatalf("unexpected nil tree.\n\twant: %s\n\thave: %s", a, b)
	}

	if a.value.Type != b.value.Type {
		t.Fatalf("unexpected tree value type.\n\twant: %s\n\thave: %s", a.value, b.value)
	}

	if a.value.Value != b.value.Value {
		t.Fatalf("unexpected tree value.\n\twant: %s\n\thave: %s", a.value.Value, b.value.Value)
	}

	assertTreeEqual(t, a.left, b.left)
	assertTreeEqual(t, a.right, b.right)
}
