# Bachelor Thesis: Functional Go - Draft

NOTE:
when talking about the transpiler, a piece of software is meant that translates
".fgo" functional go code to normal go code (".go").
The compiler then takes ".go" code and compiles it to machine level language.

These do not need to be different software products, but could be merged
in the already existing compiler, for example creating a workflow like:

```sh
# translate example.fgo to a valid example.go file
$ go functional build example.fgo
# build the example.go (with potentially other .go files in this directory)
$ go build .
```


## Mandatory Goal

The goal is to define and implement a functional language, based on Go (golang).
There are a few decisions that need to be taken:

- will / should fgo be valid go code, and just subject to further restrictions
  (to enforce a functional programming style) and optimisations, or should we
  diverge and make functional Go non-valid Go code (under some circumstances)?[1]
- implement it in a separate package or directly in the go tool / compiler (main repository)?
  When doing compiler optimisations, we'd need to touch the Go tool anyway. However,
  it could be more difficult to get the whole project into that repo. Existing
  abstractions and internal code could be used though.
- how do you achieve true immutability (e.g. x = x + 1) without introducing
  unnecessary complications? (x = func(x)?)
- make it possible to write "normal" go and functional go in a single:
  - file? (build tags)
  - package (separate files, e.g. ".fgo" & ".go")
  or not at all? if compilation is done ".fgo" -> ".go", package-level should
  certainly be possible. However, this has some implications on optimisations
  (they'd either need to be done when transpiling, or built into the compiler
  in such a way that we'd not "damage" existing Go code)

## Stretch Goals

### Immutability & pointers

As the language is now immutable, we could pass everything as pointers. This will
be more expensive though for smaller types (smaller then a pointer for sure), as
referencing and dereferencing costs too.
Benchmark different solutions (copy-vs-pointer-pass-sizes)

### Compiler optimisations

- Tail Call Optimisation (get rid of the stack). Needs to be built into the compiler,
  meaning that there should be a flag to enable / disable it.
- conversion of recursive calls to either a for-loop or goto (Single static assignment, SSA).
  This _could_ (?) be implemented in the compiler when using SSA, or in the transpiler
  by using for-loops
- more optimisations for recursive calls could be researched, implemented and benchmarked.
  This means doing some research on these kind of optimisations, finding suitable ones,
  implement them and do a performance test.

Further compiler- or even runtime-optimisations?
Go has nice support for concurrency. Can we leverage that in fgo?




[1]
Example are recursive calls, which would currently look like this:

```go
func example() {
    var x func(int) int // we need to pre-declare it so that we can recursively call it
    x = func(a int) int {
        if a == 0 {
            return 0
        }
        return a + (x(a-1))
    }
    fmt.Println(x(5))
}
```

However, trading off the "Go compatibility" just to save one line is questionable.
