package exp

import (
	"net"
	"testing"
)

var ipMap = Map{
	"ip1": "192.168.1.32",
	"ip2": "192.168.2.64",
}

func TestContainsIP(t *testing.T) {
	_, cidr, err := net.ParseCIDR("192.168.1.0/24")
	if err != nil {
		t.Error(err)
	}

	for _, test := range []struct {
		exp Exp
		out bool
	}{
		{ContainsIP("ip1", cidr), true},
		{ContainsIP("ip2", cidr), false},
	} {
		if test.exp.Eval(ipMap) != test.out {
			t.Errorf("%s should evaluate to %t.", test.exp, test.out)
		}
	}
}
