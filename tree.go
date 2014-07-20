package tree

// The Node interface represents a tree node. There are several implementations
// of the interface in this package, but one may define custom Node's as long as
// they implement the Eval function.
type Node interface {
	Eval(p Params) bool
}

// Wraps up a map
type Params map[string]string
