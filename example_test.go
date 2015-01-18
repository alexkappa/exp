package exp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func Example() {
	conjunction := And(True, True, True)
	disjunction := Or(True, False)
	negation := Not(False)

	complex := Or(And(conjunction, disjunction), negation)

	fmt.Printf("%t\n", complex.Eval(nil))

	// Output: true
}

func Example_uRL() {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "http://example.com?event=signup&date=2014-10-03&basket_value=199.90&referrer=https://google.com", nil)

	handler := func(w http.ResponseWriter, r *http.Request) {
		exp := And(
			Match("event", "signup"),
			Before("date", time.Now()),
			GreaterThan("basket_value", 99.99),
			Contains("referrer", "google.com"))

		if exp.Eval(r.URL.Query()) {
			fmt.Fprintf(w, "Expression evaluates")
		}
	}

	handler(writer, request)

	fmt.Println(writer.Body.String())
	// Output: Expression evaluates
}
