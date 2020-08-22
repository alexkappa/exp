package exp

import (
	"testing"
)

type testcase struct {
	key  string
	exp  map[string]string
	data []Data
}

type Data struct {
	InputData Map
	Result    bool
}

func TestTimeEqual(t *testing.T) {
	param := "timestamp"
	testCases := []testcase{
		{param, map[string]string{param: "2020-04-07T20:05:00+05:30"},
			[]Data{
				{map[string]string{param: "2020-04-07T20:05:00+05:30"}, true},
				{map[string]string{param: "2020-05-07T20:05:00+05:30"}, false},
				{map[string]string{param: "2020-03-07T20:05:00+05:30"}, false},
			}}}

	for _, itr := range testCases {
		for _, d := range itr.data {
			if !d.Result == TimeEqual(itr.key, itr.exp[param]).Eval(d.InputData) {
				t.Errorf("onTestTimeEqual(%q = %q) should evaluate to %v", d.InputData.Get(param), itr.exp[param], d.Result)
			}
		}
	}
}

func TestTimeGreaterThan(t *testing.T) {
	param := "timestamp"
	testCases := []testcase{
		{param, map[string]string{param: "2020-04-07T20:05:00+05:30"},
			[]Data{
				{map[string]string{param: "2020-04-07T20:05:00+05:30"}, false},
				{map[string]string{param: "2020-04-07T20:05:01+05:30"}, true},
				{map[string]string{param: "2020-04-07T20:04:59+05:30"}, false},
			}}}

	for _, itr := range testCases {
		for _, d := range itr.data {
			if !d.Result == TimeGreaterThan(itr.key, itr.exp[param]).Eval(d.InputData) {
				t.Errorf("OnTestTimeGreaterThan(%q > %q) should evaluate to %v", d.InputData.Get(param), itr.exp[param], d.Result)
			}
		}
	}
}

func TestTimeGreaterOrEqual(t *testing.T) {
	param := "timestamp"
	testCases := []testcase{
		{param, map[string]string{param: "2020-04-07T20:05:00+05:30"},
			[]Data{
				{map[string]string{param: "2020-04-07T20:05:00+05:30"}, true},
				{map[string]string{param: "2020-04-07T20:05:01+05:30"}, true},
				{map[string]string{param: "2020-04-07T20:04:00+05:30"}, false},
			}}}

	for _, itr := range testCases {
		for _, d := range itr.data {
			if !d.Result == TimeGreaterOrEqual(itr.key, itr.exp[param]).Eval(d.InputData) {
				t.Errorf("OnTimeGreaterOrEqual(%q >= %q) should evaluate to %v", d.InputData.Get(param), itr.exp[param], d.Result)
			}
		}
	}
}

func TestOnTimeLessThan(t *testing.T) {
	param := "timestamp"
	testCases := []testcase{
		{param, map[string]string{param: "2020-04-07T20:05:00+05:30"},
			[]Data{
				{map[string]string{param: "2020-04-07T20:05:00+05:30"}, false},
				{map[string]string{param: "2020-04-07T20:05:01+05:30"}, false},
				{map[string]string{param: "2020-04-07T20:04:00+05:30"}, true},
			}}}

	for _, itr := range testCases {
		for _, d := range itr.data {
			if !d.Result == TimeLessThan(itr.key, itr.exp[param]).Eval(d.InputData) {
				t.Errorf("OnTimeLessThan(%q < %q) should evaluate to %v", d.InputData.Get(param), itr.exp[param], d.Result)
			}
		}
	}
}

func TestOnTimeLessOrEqual(t *testing.T) {
	param := "timestamp"
	testCases := []testcase{
		{param, map[string]string{param: "2020-04-07T20:05:00+05:30"},
			[]Data{
				{map[string]string{param: "2020-04-07T20:05:00+05:30"}, true},
				{map[string]string{param: "2020-04-07T20:05:01+05:30"}, false},
				{map[string]string{param: "2020-04-07T20:04:00+05:30"}, true},
			}}}

	for _, itr := range testCases {
		for _, d := range itr.data {
			if !d.Result == TimeLessOrEqual(itr.key, itr.exp[param]).Eval(d.InputData) {
				t.Errorf("OnTimeLessOrEqual(%q <= %q) should evaluate to %v", d.InputData.Get(param), itr.exp[param], d.Result)
			}
		}
	}
}

func TestOnTimeContains(t *testing.T) {
	param := "timestamp"
	testCases := []testcase{
		{param, map[string]string{param: "2020-04-07T20:05:00+05:30"},
			[]Data{
				{map[string]string{param: "This 2020-04-07T20:05:00+05:30 is matching timestamp "}, true},
				{map[string]string{param: "2020-04-07T20:05:00+05:30"}, true},
				{map[string]string{param: "This 2020-04-07T20:04:00+05:30 is not matching timestamp"}, false},
			}}}

	for _, itr := range testCases {
		for _, d := range itr.data {
			if !d.Result == TimeContains(itr.key, itr.exp[param]).Eval(d.InputData) {
				t.Errorf("OnTimeContains(%q : %q) should evaluate to %v", d.InputData.Get(param), itr.exp[param], d.Result)
			}
		}
	}
}
