package exp

import "net"

// Matches

type expContainsIp struct {
	key, str string
}

func (e expContainsIp) Eval(p Params) bool {
  _, cidrnet, err := net.ParseCIDR(p.Get(e.key))
  if err != nil {
		return false
	}
  testIp := net.ParseIP(e.str)
  if err != nil {
		return false
	}

	return cidrnet.Contains(testIp)
}


func (e expContainsIp) String() string {
	return sprintf("[%s∋%s]", e.key, e.str)
}

// Contains is an expression that evaluates to true if substr is within the
// value pointed to by key.
func ContainsIp(key, substr string) Exp {
	return expContainsIp{key, substr}
}
