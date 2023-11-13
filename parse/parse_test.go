// Copyright (c) 2014 Alex Kalyvitis

package parse

import (
	"testing"
)

func TestParser(t *testing.T) {
	for _, test := range []struct {
		exp string
		ast *tree
	}{
		{
			"foo > bar",
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
			"((f > 1) && (f < 3)) && ((b > 1) && (b < 3))",
			&tree{
				value: token{Type: T_LOGICAL_AND, Value: "&&"},
				left: &tree{
					value: token{Type: T_LOGICAL_AND, Value: "&&"},
					left: &tree{
						value: token{Type: T_IS_GREATER, Value: ">"},
						left: &tree{
							value: token{Type: T_IDENTIFIER, Value: "f"},
						},
						right: &tree{
							value: token{Type: T_NUMBER, Value: "1"},
						},
					},
					right: &tree{
						value: token{Type: T_IS_SMALLER, Value: "<"},
						left: &tree{
							value: token{Type: T_IDENTIFIER, Value: "f"},
						},
						right: &tree{
							value: token{Type: T_NUMBER, Value: "3"},
						},
					},
				},
				right: &tree{
					value: token{Type: T_LOGICAL_AND, Value: "&&"},
					left: &tree{
						value: token{Type: T_IS_GREATER, Value: ">"},
						left: &tree{
							value: token{Type: T_IDENTIFIER, Value: "b"},
						},
						right: &tree{
							value: token{Type: T_NUMBER, Value: "1"},
						},
					},
					right: &tree{
						value: token{Type: T_IS_SMALLER, Value: "<"},
						left: &tree{
							value: token{Type: T_IDENTIFIER, Value: "b"},
						},
						right: &tree{
							value: token{Type: T_NUMBER, Value: "3"},
						},
					},
				},
			},
		},
		{
			"((f > 1) && (f < 3))",
			&tree{
				value: token{Type: T_LOGICAL_AND, Value: "&&"},
				left: &tree{
					value: token{Type: T_IS_GREATER, Value: ">"},
					left: &tree{
						value: token{Type: T_IDENTIFIER, Value: "f"},
					},
					right: &tree{
						value: token{Type: T_NUMBER, Value: "1"},
					},
				},
				right: &tree{
					value: token{Type: T_IS_SMALLER, Value: "<"},
					left: &tree{
						value: token{Type: T_IDENTIFIER, Value: "f"},
					},
					right: &tree{
						value: token{Type: T_NUMBER, Value: "3"},
					},
				},
			},
		},
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

		if !treeEquals(test.ast, ast) {
			t.Fatalf("trees are not equal.\n\twant: %s\n\thave: %s", test.ast, ast)
		}
	}
}

func treeEquals(a, b *tree) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.value.Type != b.value.Type {
		return false
	}

	if a.value.Value != b.value.Value {
		return false
	}

	if !treeEquals(a.left, b.left) || !treeEquals(a.right, b.right) {
		return false
	}

	return true
}
