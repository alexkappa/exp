// Copyright (c) 2016 Alex Kalyvitis

package parse

import "fmt"

type parser struct {
	lexer *lexer
	buf   []token
}

// read returns the next token from the lexer and advances the cursor. This
// token will not be available by the parser after it has been read.
func (p *parser) read() token {
	if len(p.buf) > 0 {
		r := p.buf[0]
		p.buf = p.buf[1:]
		return r
	}
	return p.lexer.token()
}

// errorf creates a parsing error which describes the token currently being
// processed as well as line and column numbers from the input stream.
func (p *parser) errorf(t token, format string, v ...interface{}) error {
	return fmt.Errorf("%d:%d syntax error: %s", t.line, t.col, fmt.Sprintf(format, v...))
}

// parse requests tokens from the lexer and generates
func (p *parser) parse() (*tree, error) {

	tree := newTree()
	node := tree

	stack := newStack()
	stack.push(tree)

	var err error

loop:
	for {
		token := p.read()
		switch token.Type {
		case T_LEFT_PAREN:
			node.left = newTree()
			stack.push(node)
			node = node.left
		case T_RIGHT_PAREN:
			node, err = stack.pop()
			if err != nil {
				break loop
			}
		case T_LOGICAL_AND, T_LOGICAL_OR, T_LOGICAL_NOT, T_IS_EQUAL, T_IS_NOT_EQUAL, T_IS_GREATER, T_IS_GREATER_OR_EQUAL, T_IS_SMALLER, T_IS_SMALLER_OR_EQUAL:
			if node.value.Type != T_ERR {
				tree = newTree()
				tree.left = node
				tree.right = newTree()
				tree.value = token
				stack.push(tree)
				node = tree.right
			} else {
				node.value = token
				node.right = newTree()
				stack.push(node)
				node = node.right
			}
		case T_IDENTIFIER, T_NUMBER, T_STRING, T_BOOLEAN:
			node.value = token
			node, err = stack.pop()
			if err != nil {
				break loop
			}
		case T_EOF:
			break loop
		case T_ERR:
			return nil, p.errorf(token, "%s", token.Value)
		default:
			return nil, p.errorf(token, "unknown token %s", token.Type)
		}
	}

	return tree, err
}

// newParser creates a new parser using the suppliad lexer.
func newParser(l *lexer) *parser {
	return &parser{lexer: l}
}

// Parse parses an expression in text format and returns the parse tree.
func Parse(s string) (Tree, error) {
	l := newLexer(s)
	p := newParser(l)
	return p.parse()
}
