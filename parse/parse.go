// Copyright (c) 2016 Alex Kalyvitis

package parse

import (
	"fmt"
	"io"
)

type parser struct {
	lexer *lexer
	buf   []token
	ast   Exp
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
// coundn't read all n tokens, for example if a tokenEOF was returned by the
// lexer, an error is returned and the returned slice will have all tokens read
// until that point, including tokenEOF.
func (p *parser) readn(n int) ([]token, error) {
	tokens := make([]token, 0, n) // make a slice capable of storing up to n tokens
	for i := 0; i < n; i++ {
		tokens = append(tokens, p.read())
		if tokens[i].typ == tokenEOF {
			return tokens, io.EOF
		}
	}
	return tokens, nil
}

// readt returns the tokens starting from the current position until the first
// match of t. Similar to readn it will return an error if a tokenEOF was
// returned by the lexer before a match was made.
func (p *parser) readt(t tokenType) ([]token, error) {
	var tokens []token
	for {
		token := p.read()
		tokens = append(tokens, token)
		switch token.typ {
		case tokenEOF:
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
		if t.typ == tokenEOF {
			return p.buf, io.EOF
		}
	}
	return p.buf, nil
}

// peekt returns the tokens from the current postition until the first match of
// t. it will not advance the cursor.
func (p *parser) peekt(t tokenType) ([]token, error) {
	for i := 0; i < len(p.buf); i++ {
		switch p.buf[i].typ {
		case t:
			return p.buf[:i], nil
		case tokenEOF:
			return p.buf[:i], io.EOF
		}
	}
	for {
		token := p.lexer.token()
		p.buf = append(p.buf, token)
		switch token.typ {
		case t:
			return p.buf, nil
		case tokenEOF:
			break
		}
	}
	return p.buf, io.EOF
}

func (p *parser) errorf(t token, format string, v ...interface{}) error {
	return fmt.Errorf("%d:%d syntax error: %s", t.line, t.col, fmt.Sprintf(format, v...))
}

// parse begins parsing based on tokens read from the lexer.
func (p *parser) parse() (Exp, error) {
	var exp Exp

loop:
	for {
		token := p.read()
		switch token.Type {
		case T_EOF:
			break loop
		case T_ERR:
			return nil, p.errorf(token, "%s", token.Value)
		case T_IDENTIFIER:

		case T_LEFT_PAREN:
			t, err := p.readt(T_RIGHT_PAREN)
			if err != nil {
				return nil, p.errorf(t[len(t)-1], "%s", token.Value)
			}
			p = subParser(t)
			e, err := p.parse()
			if err != nil {
				return nil, p.errorf(t[len(t)-1], "%s", token.Value)
			}
		}
	}

	// 	var nodes []node
	// loop:
	// 	for {
	// 		token := p.read()
	// 		switch token.typ {
	// 		case tokenEOF:
	// 			break loop
	// 		case tokenError:
	// 			return nil, p.errorf(token, "%s", token.val)
	// 		case tokenText:
	// 			nodes = append(nodes, textNode(token.val))
	// 		case tokenLeftDelim:
	// 			node, err := p.parseTag()
	// 			if err != nil {
	// 				return nodes, err
	// 			}
	// 			nodes = append(nodes, node)
	// 		case tokenRawStart:
	// 			node, err := p.parseRawTag()
	// 			if err != nil {
	// 				return nodes, err
	// 			}
	// 			nodes = append(nodes, node)
	// 		case tokenSetDelim:
	// 			nodes = append(nodes, new(delimNode))
	// 		}
	// 	}
	return exp, nil
}

// newParser creates a new parser using the suppliad lexer.
func newParser(l *lexer) *parser {
	return &parser{lexer: l}
}

// subParser creates a new parser with a pre-defined token buffer.
func subParser(b []token) *parser {
	return &parser{buf: append(b, token{Type: T_EOF})}
}
