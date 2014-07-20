# tree

Binary Expression Tree

```
$ go get go.alxkp.co/tree.v0
```

### Boolean

```go
package main

import _ "go.alxkp.co/tree.v0"

func main() {
    and := And{True{}, True{}}
    and.Eval(nil) // true

    or := Or{True{}, False{}}
    or.Eval(nil) // true

    tree := Or{And{True{}, True{}}, Or{True{}, False{}}}
    tree.Eval(nil) // true
}
```

### Algebraic

```go
func main() {
    eq := Eq{"foo", "bar"}
    eq.Eval(Params{"foo": "bar"}) // true
    eq.Eval(Params{"bar": "baz"}) // false
}
```

### Combined

```go
func main() {
    // true if foo=bar and bar=baz
    tree := And{Eq{"foo", "bar"}, Eq{"bar", "baz"}}
    tree.Eval(Params{"foo": "bar", "bar": "baz"}) // true
}
```

### Custom Node

```go
import "strings"

type Like struct {
    Key, Value string
}

func (l Like) Eval(p Params) bool {
    if val, found := p[l.Key]; found {
        return strings.Contains(val, l.Value)
    }
    return false
}

func main() {
    l := Like{"foo", "bar"}
    if !l.Eval(Params{"foo": "barracuda"}) {
        t.Error("the value of foo (barracuda) should contain bar")
    }
}
```