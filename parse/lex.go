// Copyright (c) 2016 Alex Kalyvitis
// Portions Copyright (c) 2011 The Go Authors

package parse

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// token represents a token or text string returned from the scanner.
type token struct {
	Type  tokenType
	Value string
	Line  int
	Col   int
}

// String satisfies the fmt.Stringer interface making it easier to print tokens.
func (i token) String() string {
	return fmt.Sprintf("%s:%q", i.Type, i.Value)
}

// tokenType identifies the type of lex tokens.
type tokenType int

const (
	T_UNKNOWN tokenType = iota

	T_ERR
	T_EOF

	T_IDENTIFIER

	T_NUMBER
	T_STRING
	T_BOOLEAN

	T_LOGICAL_AND
	T_LOGICAL_OR
	T_LOGICAL_NOT

	T_LEFT_PAREN
	T_RIGHT_PAREN

	T_IS_EQUAL
	T_IS_NOT_EQUAL
	T_IS_GREATER
	T_IS_GREATER_OR_EQUAL
	T_IS_SMALLER
	T_IS_SMALLER_OR_EQUAL
)

var tokenName = map[tokenType]string{
	T_UNKNOWN:             "T_UNKNOWN",
	T_ERR:                 "T_ERR",
	T_EOF:                 "T_EOF",
	T_IDENTIFIER:          "T_IDENTIFIER",
	T_NUMBER:              "T_NUMBER",
	T_STRING:              "T_STRING",
	T_BOOLEAN:             "T_BOOLEAN",
	T_LOGICAL_AND:         "T_LOGICAL_AND",
	T_LOGICAL_OR:          "T_LOGICAL_OR",
	T_LOGICAL_NOT:         "T_LOGICAL_NOT",
	T_LEFT_PAREN:          "T_LEFT_PAREN",
	T_RIGHT_PAREN:         "T_RIGHT_PAREN",
	T_IS_EQUAL:            "T_IS_EQUAL",
	T_IS_NOT_EQUAL:        "T_IS_NOT_EQUAL",
	T_IS_GREATER:          "T_IS_GREATER",
	T_IS_GREATER_OR_EQUAL: "T_IS_GREATER_OR_EQUAL",
	T_IS_SMALLER:          "T_IS_SMALLER",
	T_IS_SMALLER_OR_EQUAL: "T_IS_SMALLER_OR_EQUAL",
}

// String satisfies the fmt.Stringer interface making it easier to print tokens.
func (i tokenType) String() string {
	s := tokenName[i]
	if s == "" {
		return fmt.Sprintf("T_UNKNOWN_%d", int(i))
	}
	return s
}

const eof = -1

// stateFn represents the state of the scanner as a function that returns the
// next state.
type stateFn func(*lexer) stateFn

// lexer holds the state of the scanner.
type lexer struct {
	input  string     // the string being scanned.
	state  stateFn    // the next lexing function to enter.
	pos    int        // current position in the input.
	start  int        // start position of this token.
	width  int        // width of last rune read from input.
	tokens chan token // channel of scanned tokens.
}

// next returns the next rune in the input.
func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

// seek advances the pointer by n spaces.
func (l *lexer) seek(n int) {
	l.pos += n
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

// buffer returns the string consumed by the lexer after the last emit.
func (l *lexer) buffer() string {
	return l.input[l.start:l.pos]
}

// emit passes an token back to the client.
func (l *lexer) emit(t tokenType) {
	l.tokens <- token{
		t,
		l.buffer(),
		l.lineNum(),
		l.columnNum(),
	}
	l.start = l.pos
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}

// lineNum reports which line we're on. Doing it this way
// means we don't have to worry about peek double counting.
func (l *lexer) lineNum() int {
	return 1 + strings.Count(l.input[:l.pos], "\n")
}

// columnNum reports the character of the current line we're on.
func (l *lexer) columnNum() int {
	if lf := strings.LastIndex(l.input[:l.pos], "\n"); lf != -1 {
		return len(l.input[lf+1 : l.pos])
	}
	return len(l.input[:l.pos])
}

// error returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.token.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.tokens <- token{
		T_ERR,
		fmt.Sprintf(format, args...),
		l.lineNum(),
		l.columnNum(),
	}
	return nil
}

// token returns the next token from the input.
func (l *lexer) token() token {
	for {
		select {
		case t := <-l.tokens:
			return t
		default:
			l.state = l.state(l)
		}
	}
}

func (l *lexer) String() string {
	w := bytes.NewBuffer(nil)
	fmt.Fprintf(w, "Expression: %q\n", l.input)
	fmt.Fprintf(w, "Index     : %d\n", l.pos)
	fmt.Fprintf(w, "Current   : %q\n", l.input[l.pos])
	fmt.Fprintf(w, "Buffer    : %q\n", l.input[l.start:l.pos])
	return w.String()
}

// newLexer creates a new scanner for the input string.
func newLexer(input string) *lexer {
	l := &lexer{
		input:  input,
		tokens: make(chan token, 2),
	}
	l.state = stateInit
	return l
}

// state functions.

// stateInit is the initial state of the lexer.
func stateInit(l *lexer) stateFn {
start:
	switch r := l.next(); {
	case isWhitespace(r):
		l.ignore()
		goto start
	case isNumeric(r):
		return stateNumber
	case isAlphanum(r):
		return stateIdentifier
	case isOperator(r):
		return stateOperator
	case r == '\'':
		return stateSingleQuote
	case r == '"':
		return stateDoubleQuote
	case r == '(':
		l.emit(T_LEFT_PAREN)
		goto start
	case r == ')':
		l.emit(T_RIGHT_PAREN)
		goto start
	}
	return stateEnd
}

// stateEnd is the final state of the lexer. After this state is entered no more
// tokens can be requested as it will result in a nil pointer dereference.
func stateEnd(l *lexer) stateFn {
	// Always end with EOF token. The parser will keep asking for tokens until
	// an T_EOF or T_ERR token are encountered.
	l.emit(T_EOF)

	return nil
}

// stateIdentifier scans an indentifier from the input stream. An identifier is
// a variable which will be substituted with a concrete value during expression
// evaluation.
func stateIdentifier(l *lexer) stateFn {
loop:
	for {
		switch r := l.next(); {
		case isAlphanum(r):
		default:
			break loop
		}
	}

	l.backup()

	switch l.buffer() {
	case "true", "false":
		l.emit(T_BOOLEAN)
	default:
		l.emit(T_IDENTIFIER)
	}

	return stateInit
}

// stateOperator scans an operator from the input stream.
func stateOperator(l *lexer) stateFn {
	r := l.next()
	for isOperator(r) {
		r = l.next()
	}

	l.backup()

	switch l.buffer() {
	case "!":
		l.emit(T_LOGICAL_NOT)
	case "&&":
		l.emit(T_LOGICAL_AND)
	case "||":
		l.emit(T_LOGICAL_OR)
	case "==":
		l.emit(T_IS_EQUAL)
	case "!=":
		l.emit(T_IS_NOT_EQUAL)
	case ">":
		l.emit(T_IS_GREATER)
	case ">=":
		l.emit(T_IS_GREATER_OR_EQUAL)
	case "<":
		l.emit(T_IS_SMALLER)
	case "<=":
		l.emit(T_IS_SMALLER_OR_EQUAL)
	}

	return stateInit
}

// stateSingleQuote scans an identifier enclosed in single quotes from the input
// stream.
func stateSingleQuote(l *lexer) stateFn {
	l.ignore()
loop:
	for {
		switch l.next() {
		case '\'':
			l.backup()
			l.emit(T_IDENTIFIER)
			l.next()
			l.ignore()
			break loop
		case eof:
			return l.errorf("unexpected EOF")
		}
	}

	return stateInit
}

// stateDoubleQuote scans a string enclosed in double quotes from the input
// stream.
func stateDoubleQuote(l *lexer) stateFn {
	l.ignore()
loop:
	for {
		switch l.next() {
		case '"':
			l.backup()
			l.emit(T_STRING)
			l.next()
			l.ignore()
			break loop
		case eof:
			return l.errorf("unexpected EOF")
		}
	}

	return stateInit
}

// stateNumber scans a numeric value from the input stream.
func stateNumber(l *lexer) stateFn {
loop:
	switch r := l.next(); {
	case isNumeric(r) || r == '.':
		goto loop
	}

	l.backup()
	l.emit(T_NUMBER)

	return stateInit
}

// isWhitespace reports whether r is a space character.
func isWhitespace(r rune) bool {
	switch r {
	case ' ', '\t', '\n', '\r':
		return true
	}
	return false
}

// isNumeric reports whether r is a digit.
func isNumeric(r rune) bool {
	return unicode.IsDigit(r)
}

// isAlphanum reports whether r is an alphabetic, digit, or underscore.
func isAlphanum(r rune) bool {
	return r == '_' || r == '.' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

// isOperator reports whether r is one of the predefined operators.
func isOperator(r rune) bool {
	return r == '=' || r == '!' || r == '>' || r == '<' || r == '&' || r == '|'
}
