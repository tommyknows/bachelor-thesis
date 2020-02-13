// This file demonstrates Sum types in Go.
// It is questionable if support is needed, as
// returning multiple values could be used somehow.
package main

// There is no built in support, so we need to "hack around".
// This hack consists of declaring an interface and implementing
// a necessary dummy-function.
type MaybeInt interface {
	MaybeInt()
}

type MyInt int

func (m MyInt) MaybeInt() {}

type Nothing struct{}

func (n Nothing) MaybeInt() {}

func Example(x MaybeInt) MyInt {
	switch y := x.(type) {
	case MyInt:
		return y + MyInt(3)
	case Nothing:
		return MyInt(0)
	}
	// No compiler support for "exhaustive" pattern matching checks.
	// The compiler does not know that only two types implement that
	// interface (which is technically true, for interfaces). For
	// Sum-Types, it should / would know that MaybeInt only consists
	// of either a MyInt or Nothing.
	panic("We will never get here")
}
