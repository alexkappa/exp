# Binary Expression Tree

[![Go Reference](https://pkg.go.dev/badge/github.com/alexkappa/exp.svg)](https://pkg.go.dev/github.com/alexkappa/exp)
[![example workflow](https://github.com/alexkappa/exp/actions/workflows/go.yml/badge.svg)
](https://github.com/alexkappa/exp/actions/workflows/go.yml)
[![Code Climate](https://codeclimate.com/repos/57ee74cb0cee2109cb001a8d/badges/df8b36b023b964ac23ca/gpa.svg)](https://codeclimate.com/repos/57ee74cb0cee2109cb001a8d/feed)

Package exp implements a binary expression tree which can be used to evaluate
arbitrary binary expressions. You can use this package to build your own
expressions however a few expressions are provided out of the box.

## Installation

```
$ go get github.com/alexkappa/exp/...
```

## Usage

Expressions can be used as a rule evaluation engine, where rules are created and
evaluated against a target object. A simple example is checking whether a value
falls under a pre-defined threshold. We could use an expression to create the
rule.

```Go
condition := exp.LessThan("value", 10.0)

if condition.Eval(exp.Map{"value": 9.99}) {
	alert.Dispatch()
}
```

```Go
exp.True.Eval(nil)                     // true
exp.False.Eval(nil)                    // false
exp.Not(exp.False).Eval(nil)           // true
exp.Or(exp.True, exp.False).Eval(nil)  // true
exp.And(exp.True, exp.False).Eval(nil) // false

exp.And(
	exp.LessThan("value", 10)
	exp.GreaterThan("value", 5)
).Eval(
	exp.Map{"value": 7}
) // true
```

It is also possible to use text to describe expressions.

**Warning** this feature is not battle-tested so use it with caution.

```Go
x, err := exp.Parse(`(foo >= 100.00)`)
if err != nil {
	// handle error
}
x.Eval(exp.Map{"foo": "150.00"}) // true
```

At this time, the following operators are supported. More data types and
operators will be added in the future.

| Expression                             | Operator                   | Data Type |
| -------------------------------------- | -------------------------- | --------- |
| `true`                                 | `True`                     | `bool`    |
| `false`                                | `False`                    | `false`   |
| `(true && true)`                       | `And`                      | `any`     |
| <code>(false &#124;&#124; true)</code> | `Or`                       | `any`     |
| `foo == "xxx"`                         | `Match`                    | `string`  |
| `foo != "xxx"`                         | `Not(Match)`               | `string`  |
| `foo == 123`                           | `Equal`, `Eq`              | `float64` |
| `foo != 123`                           | `NotEqual`, `Neq`          | `float64` |
| `foo > 123`                            | `GreaterThan `, `Gt`       | `float64` |
| `foo >= 123`                           | `GreaterThanEqual `, `Gte` | `float64` |
| `foo < 123`                            | `LessThan `, `Lt`          | `float64` |
| `foo <= 123`                           | `LessThanEqual `, `Lte`    | `float64` |
