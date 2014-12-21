package exp

import "strings"

// Contains

type expContains struct {
	key, substr string
}

func (e expContains) Eval(p Params) bool {
	return strings.Contains(p.Get(e.key), e.substr)
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

// Count evaluates to true if the number of non-overlapping instances of sep in
// the value pointed to by key is equal to count.
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

// EqualFold evaluates to true if the value pointed to by key and s, interpreted
// as UTF-8 strings, are equal under Unicode case-folding.
func EqualFold(key, s string) Exp {
	return expEqualFold{key, s}
}
