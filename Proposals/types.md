# Bachelor Thesis Proposal: Types

## Goal

The goal is to implement Algebraic Data Types in Go. It should concentrate
mainly on Unions. The reason for this is that tagged unions already exist in
the form of structs and tuples as a first class-type can be implemented with
structs too.

## Description

The thesis will start off with the 'alpha'-implementation for parametric
polymorphism (see [design](https://go.googlesource.com/proposal/+/master/design/go2draft-generics-overview.md)).
Leveraging this, Unions will be added to this design. Most probably, the
type system will have to be (re)written too, to support the additional
complexity.
Go is known for its simple syntax. This should be kept as far as possible.
Here is an example for a union in Go:

```go
type Nothing struct{}
type Maybe union {
    int
    Nothing
}
```

This would implement the `Maybe Int` type in Haskell. To allow for generic
data types, as for example in Haskell `Maybe`, using the
contracts proposal:

```go
type Nothing struct{}
type Maybe(type T) union {
    T
    Nothing
}
```

Unions would leverage the built-in type switch:

```go
func Concrete(type T)(t Maybe T) (T, error) {
    switch v := t.(type) {
    case T:
        return v, nil
    case Nothing:
        var zero T
        return zero, ErrNoConcrete
    }
}
```

The compiler would enforce the checking of every type of a union in a type switch.

## Problems

There will be a lot of unknown aspects of actually implementing this. Especially
considering the type system and the concrete implementation under the hood.

Also, it has to be checked if this design actually makes sense and solves real world
problems. Is there a benefit in implementing this that offsets the cost of:

- possibly compile-speed losses because of a more complex type system
- increasing the complexity of the language

Another thing that needs to be checked is _if_ this work would need the contracts
design to be implemented. There is a [draft](https://go-review.googlesource.com/c/go/+/187317),
which is only a prototype and contains missing features. It is unclear if this work
could depend on the early implementation, if needed.

## Related Notes

- Discussion of Sum Types / Union Types in Go can be found [here](https://github.com/golang/go/issues/19412)
