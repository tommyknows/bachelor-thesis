// BA Proposal:
// Functional Go
//
// Goal: Rewrite / extend the Go compiler with logic to
// have a functional first programming language.
// Work that could be done:
// - Use build tags to guard functional code and
//   make it possible to combine functional with normal code?
//   (like cgo build tags, have fgo build tags)
// - Extend the compiler to warn on non-functional statements
//   -> for-loops, if-statements (use switch!), channels (?)
// - "Native" partial application?
//   Right now, partial application could be done like the
//   `partialFunc` example in this file.
//   This is, however, not really beautiful ;-)
//   Also, make all variable declarations "const"
// - Will Variadic functions still be possible somehow?
// - Pointers in functional programming?

package main

import "fmt"

// Haskell:
// a -> a -> a -> a
// Function that allows for partial application
func partialFunc(x int) func(int) func(int) int {
	// technically, this is it
	// but not really readable :-(
	// the program would always (?) need this and have the logic
	// in the innermost return
	return func(y int) func(z int) int {
		return func(z int) int {
			return x + y + z
		}
	}
}

// it would be cool if the syntax could be "standard" go:
// func (x, y, z int) int {
//   return x + y + z
// }
// and have this to be rewritten to the `partialFunc` example

func main() {
	const a = 3
	x := partialFunc(3)(5)
	fmt.Println(x(1))
	fmt.Println(x(2))
}

// x := "abc" -> const x := "abc"
// sadly, assigned return values from functions cannot be
// made constant:
// const x := partialFunc(3)(5) -> `const initializer [x] is not a constant
// and function literals cannot be made constant either:
// const y = func() { fmt.Println("hello fgo") } -> `const initializer func literal is not a constant`
//
// From the docs (https://golang.org/ref/spec#Constants):
// > A constant value is represented by a rune, integer, floating-point,
// > imaginary, or string literal, an identifier denoting a constant, a
// > constant expression, a conversion with a result that is a constant,
// > or the result value of some built-in functions  [...]
// also interesting: https://groups.google.com/forum/#!topic/golang-nuts/w5KFr9WfTAM
