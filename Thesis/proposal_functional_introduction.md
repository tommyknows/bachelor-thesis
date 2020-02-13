# Functional Go

## Situation

With the rise of Javascript, Rust, and Go, the functional programming paradigm has
gained popularity too. Though none of these programming languages are purely functional,
they all share the common feature of having the possibility to use functional concepts.
However, a lot of programmers struggle initially with the concept of functional
programming. Learning a purely functional programming language is extremely useful
to gain familiarity with these concepts. Purely functional programming languages
like Haskell though are not known for their beginner-friendliness.
What makes learning a functional language difficult is that not only does the
programmer have to learn an entirely different paradigm, but also a syntax that
is uncommon for people coming from imperative or object-oriented languages.

## Objective

The objective is to ease the entry into functional programming by providing a
"harness" for Go that enforces a purely functional style. This harness can either
be a separate stand-alone tool (like `gofmt` or `gopls`) or built into the Go
compiler directly.
The functional Go code could still be valid Go code, but the harness could also apply
simple transformations to the code to make it more performant.
Things to consider while implementing that harness would be immutability (without
introducing unnecessary complications), the possibility to make functional Go code
run and / or compile with regular Go code (code translation?).
Another possibility would be to implement something like "functional-check". What
`shellcheck` is to shell-scripts, `funccheck` is to Go code. It lints existing
Go code and points out the bits that are not functional, displaying rules and
examples. While being conceptually simpler, this would leave most options to the
programmer.

In the end, functional Go should be syntactically familiar to people that have
worked with Go (or C in that regard). With that, one can learn the concepts related
to functional programming, without also needing to learn a new language and syntax.
