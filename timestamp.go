package exp

import (
	"regexp"
	"strings"
	"time"
)

/*
Perform Timestamp eq(==), gt(>), lt(<), gte(>) and lte(<=) operations.
time.RFC3339 Timestamp is supported. Timestamp will be converted to epcoh value for operations.
*/

// Eq
type expTimeEq struct {
	key   string
	value string
}

func (eq expTimeEq) Eval(p Params) bool {
	var pValue, cValue int64
	var err error
	pValue, err = getEpochValue(p.Get(eq.key))
	if err != nil {
		return false
	}
	cValue, err = getEpochValue(eq.value)
	if err != nil {
		return false
	}

	return pValue == cValue
}

func (eq expTimeEq) String() string {
	return sprintf("(%s==%s)", eq.key, eq.value)
}

// TimeEqual ...
func TimeEqual(k string, v string) Exp {
	return expTimeEq{k, v}
}

// TimeEq is an alias for TimeEqual ...
func TimeEq(k string, v string) Exp {
	return TimeEqual(k, v)
}

// Gt
type expTimeGt struct {
	key   string
	value string
}

func (gt expTimeGt) Eval(p Params) bool {
	var pValue, cValue int64
	var err error
	pValue, err = getEpochValue(p.Get(gt.key))
	if err != nil {
		return false
	}
	cValue, err = getEpochValue(gt.value)
	if err != nil {
		return false
	}

	return pValue > cValue
}

func (gt expTimeGt) String() string {
	return sprintf("(%s>%s)", gt.key, gt.value)
}

// TimeGreaterThan ...
func TimeGreaterThan(k string, v string) Exp {
	return expTimeGt{k, v}
}

// TimeGt is an alias for TimeGreaterThan ...
func TimeGt(k string, v string) Exp {
	return TimeGreaterThan(k, v)
}

type expTimeGte struct {
	key   string
	value string
}

func (gte expTimeGte) Eval(p Params) bool {
	var pValue, cValue int64
	var err error
	pValue, err = getEpochValue(p.Get(gte.key))
	if err != nil {
		return false
	}
	cValue, err = getEpochValue(gte.value)
	if err != nil {
		return false
	}
	return pValue > cValue || pValue == cValue
}

func (gte expTimeGte) String() string {
	return sprintf("(%s>=%s)", gte.key, gte.value)
}

// TimeGreaterOrEqual ...
func TimeGreaterOrEqual(k string, v string) Exp {
	return expTimeGte{k, v}
}

// TimeGte is an alias for TimeGreaterOrEqual ...
func TimeGte(k string, v string) Exp {
	return TimeGreaterOrEqual(k, v)
}

// Lt ...
type expTimeLt struct {
	key   string
	value string
}

func (lt expTimeLt) Eval(p Params) bool {
	var pValue, cValue int64
	var err error
	pValue, err = getEpochValue(p.Get(lt.key))
	if err != nil {
		return false
	}
	cValue, err = getEpochValue(lt.value)
	if err != nil {
		return false
	}
	return pValue < cValue
}

func (lt expTimeLt) String() string {
	return sprintf("(%s<%s)", lt.key, lt.value)
}

// TimeLessThan ...
func TimeLessThan(k string, v string) Exp {
	return expTimeLt{k, v}
}

// TimeLt is an alias for TimeLessThan ...
func TimeLt(k string, v string) Exp {
	return TimeLessThan(k, v)
}

type expTimeLte struct {
	key   string
	value string
}

func (lte expTimeLte) Eval(p Params) bool {
	var pValue, cValue int64
	var err error
	pValue, err = getEpochValue(p.Get(lte.key))
	if err != nil {
		return false
	}
	cValue, err = getEpochValue(lte.value)
	if err != nil {
		return false
	}
	return pValue < cValue || pValue == cValue
}

func (lte expTimeLte) String() string {
	return sprintf("(%s<=%s)", lte.key, lte.value)
}

// TimeLessOrEqual ...
func TimeLessOrEqual(k string, v string) Exp {
	return expTimeLte{k, v}
}

// TimeLte is an alias for TimeLessOrEqual ...
func TimeLte(k string, v string) Exp {
	return TimeLessOrEqual(k, v)
}

func getEpochValue(pVal string) (value int64, err error) {
	if pVal == "" {
		value = int64(0)
	} else {
		value, err = convertTimestampToEpoch(pVal)
		if err != nil {
			return
		}
	}
	return
}

// convertTimestampToEpoch convert RFC3339 timestamp to epoch value
func convertTimestampToEpoch(val string) (value int64, err error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return
	}
	return t.UTC().Unix(), nil
}

// Regex
type expTimeContains struct {
	key, substr string
	substrRegex *regexp.Regexp
}

// TimeContains ...
//Contains is an expression that evaluates to true if timestamp as a substr is within the value pointed to by key.
func TimeContains(key, str string) Exp {
	regex := regexp.MustCompile("(?i)" + strings.ReplaceAll(str, "+", "\\+"))
	return expTimeContains{key, str, regex}
}

func (e expTimeContains) Eval(p Params) bool {
	return e.substrRegex.MatchString(p.Get(e.key))
}

func (e expTimeContains) String() string {
	return sprintf("(%s:\"%s\")", e.key, e.substr)
}
