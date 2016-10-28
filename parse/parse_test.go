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
		assertEqual(t, test.ast, ast)
	}
}
