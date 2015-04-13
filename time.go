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
//
// 	m := Map{
// 		"past":   "1995-01-01",
// 		"future": "2045-01-01",
// 	}
// 	Before("past", time.Now()).Eval(m) // true
// 	Before("future", time.Now()).Eval(m) // false
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

// After is an expression that evaluates to true if date is a time after the
// evaluated date. The value is parsed to a time.Time before comparing.
func After(key string, date time.Time) Exp {
	return expAfter{key, date}
}

// Weekday

type expWeekday struct {
	key     string
	weekday time.Weekday
}

func (w expWeekday) Eval(p Params) bool {
	date, err := time.Parse(dateFormat, p.Get(w.key))
	if err != nil {
		return false
	}
	return date.Weekday() == w.weekday
}

// Weekday is an expression that evaluates to true if the date pointed to by key
// is on the specified weekday.
func Weekday(key string, weekday time.Weekday) Exp {
	return expWeekday{key, weekday}
}

// Day

type expDay struct {
	key string
	day int
}

func (d expDay) Eval(p Params) bool {
	date, err := time.Parse(dateFormat, p.Get(d.key))
	if err != nil {
		return false
	}
	return date.Day() == d.day
}

// Day is an expression that evaluates to true if the date pointed to by key is
// on the specified day.
func Day(key string, day int) Exp {
	return expDay{key, day}
}

// Month

type expMonth struct {
	key   string
	month time.Month
}

func (m expMonth) Eval(p Params) bool {
	date, err := time.Parse(dateFormat, p.Get(m.key))
	if err != nil {
		return false
	}
	return date.Month() == m.month
}

// Month is an expression that evaluates to true if the date pointed to by key
// is on the specified month.
func Month(key string, month time.Month) Exp {
	return expMonth{key, month}
}

// Year

type expYear struct {
	key  string
	year int
}

func (y expYear) Eval(p Params) bool {
	date, err := time.Parse(dateFormat, p.Get(y.key))
	if err != nil {
		return false
	}
	return date.Year() == y.year
}

// Year is an expression that evaluates to true if the date pointed to by key
// is on the specified year.
func Year(key string, year int) Exp {
	return expYear{key, year}
}

// The default format used to parse dates.
var dateFormat = "2006-01-02"

// DateFormat changes the date format used to parse dates and returnes the
// previous format in case you need to revert back in the future.
func DateFormat(f string) string {
	var previous = dateFormat
	dateFormat = f
	return previous
}
