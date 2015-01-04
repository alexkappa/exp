package exp

import (
	"testing"
	"time"
)

var d = Map{
	"past":   "1969-07-10",
	"now":    "2014-12-15",
	"future": "2022-03-09",
	"ansic":  "02 Mar 2001 14:00 UTC",
}

func TestOn(t *testing.T) {
	date := time.Date(2014, time.December, 15, 0, 0, 0, 0, time.UTC)
	if !On("now", date).Eval(d) {
		t.Errorf("On(%q, %q) should evaluate to true", d.Get("now"), date)
	}
}

func TestBefore(t *testing.T) {
	now := time.Now()
	if !Before("past", now).Eval(d) {
		t.Errorf("Before(%q, %q)", d.Get("past"), now)
	}
}

func TestAfter(t *testing.T) {
	now := time.Now()
	if !After("future", now).Eval(d) {
		t.Errorf("After(%q, %q)", d.Get("future"), now)
	}
}

func TestDateFormat(t *testing.T) {
	prev := DateFormat(time.ANSIC) // "02 Jan 06 15:04 MST"
	date := time.Date(2001, time.March, 2, 14, 0, 0, 0, time.UTC)
	if On("ansic", date).Eval(d) {
		t.Errorf("On(%q, %q) should evauate to true", d.Get("ansic"), date)
	}
	DateFormat(prev)
}
