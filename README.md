# Binary Expression Tree

Package exp implements a binary expression tree which can be used to evaluate
arbitrary binary expressions. You can use this package to build your own
expressions however a few expressions are provided out of the box.

## Installation

```
$ go get github.com/alexkappa/go-exp/exp
```

## Usage

```
conjunction := And(True, True, True)
disjunction := Or(True, False)
negation := Not(False)

complex := Or(And(conjunction, disjunction), negation)

fmt.Printf("%t\n", complex.Eval(p)) // true
```

## Documentation

API documentation is available at [godoc](https://godoc.org/github.com/alexkappa/go-exp/exp).