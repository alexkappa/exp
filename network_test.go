package exp

import (
    "testing"
)

var ipMap = Map{
	"src_ip": "192.168.1.0/24",
}

func TestContainsIp(t *testing.T) {
	for key, value := range map[string]string{
		"src_ip": "192.168.1.61",
	} {
		if !ContainsIp(key, value).Eval(ipMap) {
			t.Errorf("Match(%q, %q) should evaluate to true", key, value)
		}
	}
}
