package exp

import (
	"fmt"
	"strings"
)

func sprintf(format string, v ...interface{}) string {
	return fmt.Sprintf(format, v...)
}

func join(elems []Exp, sep string) string {
	s := make([]string, len(elems))
	for i, elem := range elems {
		s[i] = sprintf("%s", elem)
	}
	return strings.Join(s, sep)
}
