package exp

import (
	"testing"
	"time"
)

var d = Map{
	"past":    "1969-07-10",
	"present": "2014-12-15",
	"future":  "2022-03-09",
	"ansic":   "02 Mar 2001 14:00 UTC",
}

func TestOn(t *testing.T) {
	today := time.Date(2014, time.December, 15, 0, 0, 0, 0, time.UTC)
	if !On("present", today).Eval(d) {
		t.Errorf("On(%q, %q) should evaluate to true", d.Get("present"), today)
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

func TestWeekday(t *testing.T) {
	for key, weekday := range map[string]time.Weekday{
		"past":    time.Thursday,
		"present": time.Monday,
		"future":  time.Wednesday,
	} {
		if !Weekday(key, weekday).Eval(d) {
			t.Errorf("Weekday(%q, %s)", key, weekday)
		}
	}
}

func TestDay(t *testing.T) {
	for key, day := range map[string]int{
		"past":    10,
		"present": 15,
		"future":  9,
	} {
		if !Day(key, day).Eval(d) {
			t.Errorf("Day(%q, %d)", key, day)
		}
	}
}

func TestMonth(t *testing.T) {
	for key, month := range map[string]time.Month{
		"past":    time.July,
		"present": time.December,
		"future":  time.March,
	} {
		if !Month(key, month).Eval(d) {
			t.Errorf("Month(%q, %s)", key, month)
		}
	}
}

func TestYear(t *testing.T) {
	for key, year := range map[string]int{
		"past":    1969,
		"present": 2014,
		"future":  2022,
	} {
		if !Year(key, year).Eval(d) {
			t.Errorf("Year(%q, %d)", key, year)
		}
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
