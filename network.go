package exp

import "net"

// Conatins IP

type expContainsIP struct {
	key  string
	cidr *net.IPNet
}

func (e expContainsIP) Eval(p Params) bool {
	return e.cidr.Contains(net.ParseIP(p.Get(e.key)))
}

func (e expContainsIP) String() string {
	return sprintf("[%s∋%s]", e.key, e.cidr)
}

// ContainsIP is an expression that evaluates to true if an ip falls within the
// CIDR range pointed to by key.
//
//	m := Map{
//		"ip1": "192.168.1.1",
//		"ip2": "192.168.32.128",
//	}
//
//	_, cidr, err := net.ParseCIDR("192.168.1.0/24")
//	if err != nil {
//		// handle err
//	}
//
//	ContainsIP("ip1", cidr).Eval(m) // true
//	ContainsIP("ip2", cidr).Eval(m) // false
func ContainsIP(key string, cidr *net.IPNet) Exp {
	return expContainsIP{key, cidr}
}
