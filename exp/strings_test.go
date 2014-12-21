package exp

import "testing"

var m = Map{
	"foo": "bar",
	"bar": "baz",
	"baz": "booyah",
}

func TestContains(t *testing.T) {
	for key, substr := range map[string][]string{
		"foo": {"ar", "ba", "bar"},
		"bar": {"ba", "az", "baz"},
	} {
		for _, s := range substr {
			if !Contains(key, s).Eval(m) {
				t.Errorf("Contains(%q, %q) should evaluate to true", key, s)
			}
		}
	}
}

func TestContainsAny(t *testing.T) {
	for key, chars := range map[string][]string{
		"foo": {"abc", "abr", "rtu"},
		"bar": {"ab", "zax", "wea"},
		"baz": {"oq", "ya", "ha"},
	} {
		for _, c := range chars {
			if !ContainsAny(key, c).Eval(m) {
				t.Errorf("ContainsAny(%q, %q) should evaluate to true", key, c)
			}
		}
	}
}

func TestContainsRune(t *testing.T) {
	for key, runes := range map[string][]rune{
		"foo": {'b', 'a', 'r'},
		"bar": {'b', 'a', 'z'},
		"baz": {'b', 'o', 'y', 'a', 'h'},
	} {
		for _, r := range runes {
			if !ContainsRune(key, r).Eval(m) {
				t.Errorf("ContainsRune(%q, %q) should evaluate to true", key, r)
			}
		}
	}
}

func TestLen(t *testing.T) {
	for key, length := range map[string]int{
		"foo": 3,
		"bar": 3,
		"baz": 6,
	} {
		if !Len(key, length).Eval(m) {
			t.Errorf("Len(%q, %d) should evaluate to true", key, length)
		}
	}
}

func TestCount(t *testing.T) {
	for key, tests := range map[string]map[string]int{
		"foo": {"b": 1, "a": 1, "r": 1},
		"bar": {"b": 1, "a": 1, "z": 1},
		"baz": {"b": 1, "o": 2, "y": 1},
	} {
		for sep, length := range tests {
			if !Count(key, sep, length).Eval(m) {
				t.Errorf("Count(%q, %q, %d) should evaluate to true", key, sep, length)
			}
		}
	}
}

func TestEqualFold(t *testing.T) {
	for key, fold := range map[string]string{
		"foo": "Bar",
		"bar": "Baz",
		"baz": "BooYah",
	} {
		if !EqualFold(key, fold).Eval(m) {
			t.Errorf("EqualFold(%q, %q) should evaluate to true", key, fold)
		}
	}
}
