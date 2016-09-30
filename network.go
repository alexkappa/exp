package exp

import "net"

// Conatins IP

type expContainsIp struct {
	cidr, ip string
}

func (e expContainsIp) Eval(p Params) bool {
	_, cidrnet, err := net.ParseCIDR(p.Get(e.cidr))
	if err != nil {
		return false
	}
	testIp := net.ParseIP(e.ip)

	return cidrnet.Contains(testIp)
}

func (e expContainsIp) String() string {
	return sprintf("[%s∋%s]", e.cidr, e.ip)
}

// Contains is an expression that evaluates to true if substr falls within the cidr range
// given example:
//
// 192.168.1.0/24 will match all IPs that fall between
// 192.168.1.1 and 	192.168.1.254
//
// 192.168.1.0/32 will only match 192.168.1.0
func ContainsIp(cidr, ip string) Exp {
	return expContainsIp{cidr, ip}
}
