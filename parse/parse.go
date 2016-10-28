// Copyright (c) 2016 Alex Kalyvitis

package parse

import (
	"fmt"
	"io"
)

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

// readn returns the next n tokens from the lexer and advances the cursor. If it
// coundn't read all n tokens, for example if a T_EOF was returned by the
// lexer, an error is returned and the returned slice will have all tokens read
// until that point, including T_EOF.
func (p *parser) readn(n int) ([]token, error) {
	tokens := make([]token, 0, n) // make a slice capable of storing up to n tokens
	for i := 0; i < n; i++ {
		tokens = append(tokens, p.read())
		if tokens[i].Type == T_EOF {
			return tokens, io.EOF
		}
	}
	return tokens, nil
}

// readt returns the tokens starting from the current position until the first
// match of t. Similar to readn it will return an error if a T_EOF was
// returned by the lexer before a match was made.
func (p *parser) readt(t tokenType) ([]token, error) {
	var tokens []token
	for {
		token := p.read()
		tokens = append(tokens, token)
		switch token.Type {
		case T_EOF:
			return tokens, io.EOF
		case t:
			return tokens, nil
		default:
			continue
		}
	}
	return tokens, fmt.Errorf("token %q not found", t)
}

// peek returns the next token without advancing the cursor. Consecutive calls
// of peek would result in the same token being retuned. To advance the cursor,
// a read must be made.
func (p *parser) peek() token {
	if len(p.buf) > 0 {
		return p.buf[0]
	}
	t := p.lexer.token()
	p.buf = append(p.buf, t)
	return t
}

// peekn returns the next n tokens without advancing the cursor.
func (p *parser) peekn(n int) ([]token, error) {
	if len(p.buf) > n {
		return p.buf[:n], nil
	}
	for i := len(p.buf) - 1; i < n; i++ {
		t := p.lexer.token()
		p.buf = append(p.buf, t)
		if t.Type == T_EOF {
			return p.buf, io.EOF
		}
	}
	return p.buf, nil
}

// peekt returns the tokens from the current postition until the first match of
// t. it will not advance the cursor.
func (p *parser) peekt(t tokenType) ([]token, error) {
	for i := 0; i < len(p.buf); i++ {
		switch p.buf[i].Type {
		case t:
			return p.buf[:i], nil
		case T_EOF:
			return p.buf[:i], io.EOF
		}
	}
	for {
		token := p.lexer.token()
		p.buf = append(p.buf, token)
		switch token.Type {
		case t:
			return p.buf, nil
		case T_EOF:
			break
		}
	}
	return p.buf, io.EOF
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
			node.value = token
			node.right = newTree()
			stack.push(node)
			node = node.right
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
