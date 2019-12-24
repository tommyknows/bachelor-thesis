# Bachelor Thesis Proposal: Functional Go

## Goal

The goal is to implement a functional language based on the currently existing
language `Go`. With Go supporting first-class functions, making it purely functional
would consist mainly of changes to the linter and possibly compiler (for optimisations).

## Why

Last semester, I had a course on functional programming. This involved learning Haskell.
Although I am a big fan of Haskell, getting started with the functional primitives
_and_ learning a new language and syntax at the same time can be difficult
(and Haskell is not really known for its simplicity :-) ).

While Haskell's syntax is great if you are used to functional programming, what
threw me off quite a few times was that I didn't know what was being passed to
which function, with which parameters.
To ease the transition into functional programming (from a imperative or object
oriented approach), Functional Go would stand in between both. While forcing you
to write functional code, the syntax is still "imperative" (and quite verbose),
which should make it easier to understand what is going on.

Thus, Functional Go should mostly be an "educational" language, and it is not the
goal to make it "production ready" (although you may define production ready
however you want).

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

## Further Ideas and Notes

Functional Go files would likely have the file ending `.fgo`. It would be nice
if a project, maybe even a single package, could consist of both `.go` and
`.fgo` files, with zero interoperability overhead.

Once bootstrapped, the compiler optimisations could maybe also be (re)written
in Functional Go.
