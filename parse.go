package exp

import (
	"fmt"
	"strconv"

	"github.com/alexkappa/exp/parse"
)

func Parse(s string) (Exp, error) {
	t, err := parse.Parse(s)
	if err != nil {
		return nil, err
	}
	return compile(t)
}

func left(t parse.Tree) (string, error) {
	switch t.Value().Type {
	case parse.T_IDENTIFIER:
		return t.Value().Value, nil
	default:
		return "", fmt.Errorf("expected identifier but have %s instead", t.Value().Type)
	}
}

func right(t parse.Tree) (any, error) {
	switch t.Value().Type {
	case parse.T_STRING:
		return t.Value().Value, nil
	case parse.T_NUMBER:
		f, err := strconv.ParseFloat(t.Value().Value, 64)
		if err != nil {
			return nil, err
		}
		return f, nil
	default:
		return nil, fmt.Errorf("expected string or number but have %s instead", t.Value().Type)
	}
}

func compile(t parse.Tree) (Exp, error) {
	token := t.Value()
	switch token.Type {
	case parse.T_ERR:
		return nil, fmt.Errorf("parse error %v at line %d col %d", token, token.Line, token.Col)
	case parse.T_UNKNOWN:
		return nil, fmt.Errorf("unknown token at line %d col %d", token.Line, token.Col)
	case parse.T_BOOLEAN:
		switch token.Value {
		case "true":
			return True, nil
		case "false":
			return False, nil
		}
	case parse.T_LOGICAL_AND:
		l, err := compile(t.Left())
		if err != nil {
			return nil, err
		}
		r, err := compile(t.Right())
		if err != nil {
			return nil, err
		}
		return And(l, r), nil
	case parse.T_LOGICAL_OR:
		l, err := compile(t.Left())
		if err != nil {
			return nil, err
		}
		r, err := compile(t.Right())
		if err != nil {
			return nil, err
		}
		return Or(l, r), nil
	case
		parse.T_IS_EQUAL,
		parse.T_IS_NOT_EQUAL,
		parse.T_IS_GREATER,
		parse.T_IS_GREATER_OR_EQUAL,
		parse.T_IS_SMALLER,
		parse.T_IS_SMALLER_OR_EQUAL:

		var (
			k   string
			v   any
			err error
		)
		k, err = left(t.Left())
		if err != nil {
			return nil, fmt.Errorf("invalid expression. %w", err)
		}
		v, err = right(t.Right())
		if err != nil {
			return nil, fmt.Errorf("invalid expression. %w", err)
		}

		switch vv := v.(type) {
		case float64:
			switch token.Type {
			case parse.T_IS_EQUAL:
				return Equal(k, vv), nil
			case parse.T_IS_NOT_EQUAL:
				return Not(Equal(k, vv)), nil
			case parse.T_IS_GREATER:
				return GreaterThan(k, vv), nil
			case parse.T_IS_GREATER_OR_EQUAL:
				return GreaterOrEqual(k, vv), nil
			case parse.T_IS_SMALLER:
				return LessThan(k, vv), nil
			case parse.T_IS_SMALLER_OR_EQUAL:
				return LessOrEqual(k, vv), nil
			default:
				return nil, fmt.Errorf("%v is not allowed in %s expressions", vv, token.Type)
			}
		case string:
			switch token.Type {
			case parse.T_IS_EQUAL:
				return Match(k, vv), nil
			case parse.T_IS_NOT_EQUAL:
				return Not(Match(k, vv)), nil
			default:
				return nil, fmt.Errorf("%s is not allowed in %s expressions", vv, token.Type)
			}
		}
	}

	return nil, fmt.Errorf("unexpected %s:%s at line %d col %d", token.Type, token.Value, token.Line, token.Col)
}
