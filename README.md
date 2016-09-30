# Binary Expression Tree [![wercker status](https://app.wercker.com/status/3627f2113c06b84a316c4d3ab59b414c/s/master "wercker status")](https://app.wercker.com/project/byKey/3627f2113c06b84a316c4d3ab59b414c) [![Code Climate](https://codeclimate.com/repos/57ee74cb0cee2109cb001a8d/badges/df8b36b023b964ac23ca/gpa.svg)](https://codeclimate.com/repos/57ee74cb0cee2109cb001a8d/feed)

Package exp implements a binary expression tree which can be used to evaluate
arbitrary binary expressions. You can use this package to build your own
expressions however a few expressions are provided out of the box.

## Installation

```
$ go get github.com/alexkappa/exp
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

API documentation is available at [godoc](https://godoc.org/github.com/alexkappa/exp).
