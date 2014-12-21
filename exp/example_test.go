package exp

import "fmt"

func Example() {
	conjunction := And(True, True, True)
	disjunction := Or(True, False)
	negation := Not(False)

	complex := Or(And(conjunction, disjunction), negation)

	fmt.Printf("%t\n", complex.Eval(p))

	// Output: true
}
