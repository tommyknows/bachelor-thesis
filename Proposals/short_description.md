Functional languages like Haskell often come with a steep learning curve for
beginners of functional programming. Not only is the paradigm completely
different to what is being taught in imperative and object-oriented languages,
the syntax also differs considerably.

To ease the entry to functional programming, a dialect of Go (golang) should
be developed that only allows functional programming constructs.
The goal is to develop this either with a transpiler to Go, or build it into
the Go compiler directly.

The resulting language should resemble Go as much as possible, maybe even be
directly compatible and compilable with the standard Go compiler.

Then, some compiler optimisations should be implemented (possibly behind a
build-flag). These optimisations, targeted on the functional aspects, should
then be benchmarked against the non-optimised binary and, perhaps, the same
functionality implemented in an imperative paradigm.
To name a few of these optimisations:
- Tail Call Optimisation
- Copy by reference for larger values, instead of copy by reference
  - find the "sweet spot" where a value is considered large and it is cheaper
    to take the reference instead of copying it
- other optimisations can be researched and benchmarked
