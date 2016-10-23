// Copyright (c) 2016 Alex Kalyvitis

package parse

import (
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	for _, test := range []struct {
		exp    string
		tokens []token
	}{
		{
			`foo > bar`,
			[]token{
				{Type: T_IDENTIFIER, Value: "foo"},
				{Type: T_IS_GREATER, Value: ">"},
				{Type: T_IDENTIFIER, Value: "bar"},
				{Type: T_EOF},
			},
		},
		{
			`bar == 'baz'`,
			[]token{
				{Type: T_IDENTIFIER, Value: "bar"},
				{Type: T_IS_EQUAL, Value: "=="},
				{Type: T_IDENTIFIER, Value: "baz"},
				{Type: T_EOF},
			},
		},
		{
			`(bar!="baz")&&foo==123.00`,
			[]token{
				{Type: T_LEFT_PAREN, Value: "("},
				{Type: T_IDENTIFIER, Value: "bar"},
				{Type: T_IS_NOT_EQUAL, Value: "!="},
				{Type: T_STRING, Value: "baz"},
				{Type: T_RIGHT_PAREN, Value: ")"},
				{Type: T_LOGICAL_AND, Value: "&&"},
				{Type: T_IDENTIFIER, Value: "foo"},
				{Type: T_IS_EQUAL, Value: "=="},
				{Type: T_NUMBER, Value: "123.00"},
				{Type: T_EOF},
			},
		},
		{
			`!true`,
			[]token{
				{Type: T_LOGICAL_NOT, Value: "!"},
				{Type: T_BOOLEAN, Value: "true"},
				{Type: T_EOF},
			},
		},
	} {
		var tokens []token
		lexer := newLexer(test.exp)

	loop:
		for {
			token := lexer.token()
			tokens = append(tokens, token)

			if token.Type == T_EOF || token.Type == T_ERR {
				break loop
			}
		}

		// t.Logf("%v\n", tokens)

		assertEqual(t, test.tokens, tokens)
	}
}

func assertEqual(t *testing.T, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Failed asserting that items are equal.\n\twant: %s\n\thave: %s", a, b)
	}
}
