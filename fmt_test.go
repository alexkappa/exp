package exp

import "testing"

func TestString(t *testing.T) {
	for _, test := range []struct {
		exp Exp
		str string
	}{
		{True, "T"},
		{False, "F"},
		{Not(True), "¬T"},
		{And(True), "(T)"},
		{And(True, False), "(T∧F)"},
		{Or(False), "(F)"},
		{Or(True, False), "(T∨F)"},
		{Eq("foo", 10), "[foo==10.00]"},
		{Neq("foo", 5), "¬[foo==5.00]"},
		{Gt("bar", 10), "[bar>10.00]"},
		{Gte("bar", 10), "([bar>10.00]∨[bar==10.00])"},
		{Lt("bar", 5), "[bar<5.00]"},
		{Lte("bar", 5), "([bar<5.00]∨[bar==5.00])"},
		{Match("baz", "abc"), "[baz==abc]"},
		{MatchAny("baz", "abc", "bcd"), "([baz==abc]∨[baz==bcd])"},
		{Contains("foo", "bc"), "[foo∋bc]"},
		{ContainsAny("foo", "bc"), "[foo∋bc]"},
		{ContainsRune("foo", 'a'), "[foo∋a]"},
		{Len("bar", 3), "[len(bar)==3]"},
		{Count("bar", "a", 2), "[count(bar,a)==2]"},
		{EqualFold("bar", "AbC"), "[bar≈AbC]"},
	} {
		if sprintf("%s", test.exp) != test.str {
			t.Errorf("unexpected string %q != %q", test.exp, test.str)
		}
	}
}
