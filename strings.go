package exp

import "strings"

// Match

type expMatch struct {
	key, str string
}

func (e expMatch) Eval(p Params) bool {
	return p.Get(e.key) == e.str
}

func (e expMatch) String() string {
	return sprintf("[%s==%s]", e.key, e.str)
}

// Match is an expression that evaluates to true if str is equal to the value
// pointed to by key.
//
// 	m := Map{"foo": "bar"}
// 	Match("foo", "bar").Eval(m) // true
// 	Match("foo", "baz").Eval(m) // false
func Match(key, str string) Exp {
	return expMatch{key, str}
}

// MatchAny is an expression that evaluates to true if any of the strs are equal
// to the value pointed to by key.
//
// 	m := Map{"foo": "bar"}
// 	MatchAny("foo", "bar", "baz", "cux").Eval(m) // true
// 	MatchAny("foo", "baf", "baz", "lux").Eval(m) // false
func MatchAny(key string, strs ...string) Exp {
	exp := make([]Exp, len(strs))
	for i, str := range strs {
		exp[i] = Match(key, str)
	}
	return Or(exp...)
}

// Contains

type expContains struct {
	key, substr string
}

func (e expContains) Eval(p Params) bool {
	return strings.Contains(p.Get(e.key), e.substr)
}

func (e expContains) String() string {
	return sprintf("[%s∋%s]", e.key, e.substr)
}

// Contains is an expression that evaluates to true if substr is within the
// value pointed to by key.
func Contains(key, substr string) Exp {
	return expContains{key, substr}
}

// ContainsAny

type expContainsAny struct {
	key, chars string
}

func (e expContainsAny) Eval(p Params) bool {
	return strings.ContainsAny(p.Get(e.key), e.chars)
}

func (e expContainsAny) String() string {
	return sprintf("[%s∋%s]", e.key, e.chars)
}

// ContainsAny evaluates to true if any Unicode code points in chars are within
// the value pointed to by key.
func ContainsAny(key, chars string) Exp {
	return expContainsAny{key, chars}
}

// ContainsRune

type expContainsRune struct {
	key string
	r   rune
}

func (e expContainsRune) Eval(p Params) bool {
	return strings.ContainsRune(p.Get(e.key), e.r)
}

func (e expContainsRune) String() string {
	return sprintf("[%s∋%c]", e.key, e.r)
}

// ContainsRune evaluates to true if the Unicode code point r is within the
// value pointed to by key.
func ContainsRune(key string, r rune) Exp {
	return expContainsRune{key, r}
}

// Len

type expLen struct {
	key    string
	length int
}

func (e expLen) Eval(p Params) bool {
	return len(p.Get(e.key)) == e.length
}

func (e expLen) String() string {
	return sprintf("[len(%s)==%d]", e.key, e.length)
}

// Len evalates to true if the length of the string pointed to by key is equal
// to length.
func Len(key string, length int) Exp {
	return expLen{key, length}
}

// Count

type expCount struct {
	key, sep string
	count    int
}

func (e expCount) Eval(p Params) bool {
	return strings.Count(p.Get(e.key), e.sep) == e.count
}

func (e expCount) String() string {
	return sprintf("[count(%s,%s)==%d]", e.key, e.sep, e.count)
}

// Count evaluates to true if the number of non-overlapping instances of sep in
// the value pointed to by key is equal to count.
//
// 	m := Map{
// 		"foo": "bar",
// 		"bar": "zoo"
// 	}
// 	Count("foo", "b", 1).Eval(m) // true
// 	Count("foo", "a", 1).Eval(m) // true
// 	Count("foo", "r", 1).Eval(m) // true
// 	Count("bar", "o", 2).Eval(m) // true
func Count(key, sep string, count int) Exp {
	return expCount{key, sep, count}
}

// EqualFold

type expEqualFold struct {
	key, s string
}

func (e expEqualFold) Eval(p Params) bool {
	return strings.EqualFold(p.Get(e.key), e.s)
}

func (e expEqualFold) String() string {
	return sprintf("[%s≈%s]", e.key, e.s)
}

// EqualFold evaluates to true if the value pointed to by key and s, interpreted
// as UTF-8 strings, are equal under Unicode case-folding.
//
// 	m := Map{
// 		"foo": "Bar",
// 		"bar": "BaZ"
// 	}
// 	EqualFold("foo", "bar").Eval(m) // true
// 	EqualFold("bar", "baz").Eval(m) // true
func EqualFold(key, s string) Exp {
	return expEqualFold{key, s}
}
