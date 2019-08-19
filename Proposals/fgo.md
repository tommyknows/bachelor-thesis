# Bachelor Thesis Proposal: Functional Go

## Goal

The goal is to implement a functional language based on the currently existing
language `Go`. With Go supporting first-class functions, making it purely functional
would consist mainly of changes to the linter and possibly compiler (for optimisations).

## Description

Go is mainly an imperative language. While allowing functional constructs (which
is a very useful property, if not overused), it does not aim to be a purely functional
language like Haskell or F#.
While being an extremely sophisticated language, Haskell is also known for its complexity.
Functional Go should serve as an introduction to functional languages, while keeping
the syntax of "classical" imperative programming languages. This would mean that
learners could concentrate on _how to solve problems_, rather than _how to write
code_.

For Go to be a functional language, a few things would need to be implemented:

- Disallowing `for`-loops
- Disallowing mutation of state (everything immutable)
- Enforcing pure functions - no global state. -> What about methods on types?
- Only allow [purely functional data structures](https://en.wikipedia.org/wiki/Purely_functional_data_structure)
  (this would most probably require a powerful `List` implementation, see [Problems](#problems))

In addition, these could also be added:

- Tail Call Optimisation (not implemented in Go because of Stack Trace)
- Converting recursive calls into for-loops
- As nothing mutates data, pass pointers around instead of coping the data

## Problems

Functional Go would depend on the implementation of Generics (parametric
polymorphism, mainly). Go does not have any for of generics currently, but a
[proposal](https://go.googlesource.com/proposal/+/master/design/go2draft-generics-overview.md)
exists and has been [partially implemented](https://go-review.googlesource.com/c/go/+/187317).

The Bachelor Thesis could not depend on this implementation however (it is
unclear when Go 2 will be released, even an alpha state).
This would mean that functional Go would only reach a state of being practical
once Go 2 is released (and changes in fgo rebased onto that).

A list-implementation without generics could be done with the help of the compiler
(as maps and slices currently do already, too).
