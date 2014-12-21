package exp

import "time"

// On

type expOn struct {
	key  string
	date time.Time
}

func (on expOn) Eval(p Params) bool {
	date, err := time.Parse(dateFormat, p.Get(on.key))
	if err != nil {
		return false
	}
	return date.Equal(on.date)
}

// On evaluates to true if date is equal to the date pointed to by key. The
// value is parsed to a time.Time before comparing. In case of a parse error
// false is returned.
func On(key string, date time.Time) Exp {
	return expOn{key, date}
}

// Before

type expBefore struct {
	key  string
	date time.Time
}

func (b expBefore) Eval(p Params) bool {
	date, err := time.Parse(dateFormat, p.Get(b.key))
	if err != nil {
		return false
	}
	return date.Before(b.date)
}

// Before evaluates to true if date is before the date pointed to by key. The
// value is parsed to a time.Time before comparing. In case of a parse error
// false is returned.
func Before(key string, date time.Time) Exp {
	return expBefore{key, date}
}

// After

type expAfter struct {
	key  string
	date time.Time
}

func (a expAfter) Eval(p Params) bool {
	date, err := time.Parse(dateFormat, p.Get(a.key))
	if err != nil {
		return false
	}
	return date.After(a.date)
}

// After is an expression that evaluates to true if v is a date after the
// evaluated date. The value is parsed to a time.Time before comparing.
func After(key string, date time.Time) Exp {
	return expAfter{key, date}
}

// The default format used to parse dates.
var dateFormat = "2006-01-02"

// DateFormat changes the date format used to parse dates.
func DateFormat(f string) {
	dateFormat = f
}
