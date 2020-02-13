// Package example should serve as an example to how functional
// programming in Go looks like. These examples are mostly a direct
// translations from Haskell programs.
package example

// Examples taken from Sheet 2 of Functional Programming course.
type Fraction struct {
	Numerator   int
	Denominator int
}

// the way to implement Sum types would be with an interface.
// kinda ugly, and no compiler support for pattern-matching
// (e.g. no check for exhaustive), but works.
type MultResult interface {
	MultResult()
}

type Success int

func (s Success) MultResult() {}

type Failure struct{}

func (f Failure) MultResult() {}

func Mult(f Fraction, n int) MultResult {
	if (n*f.Numerator)%f.Denominator != 0 {
		return Failure{}
	}
	return Success((n * f.Numerator) / f.Denominator)
}

type Program []Fraction

func Exec(p Program, input int) []int {
	// we need to define the function fist, if we want to
	// call it recursively
	var run func([]Fraction, []int) []int
	run = func(fracs []Fraction, n []int) []int {
		switch len(fracs) {
		case 0:
			return n
		default:
			switch x := Mult(fracs[0], n[0]).(type) {
			case Success:
				// prepending something to a list (Haskells 'x:l'),
				// has to be done with the append function, which,
				// used like this, will always allocate a new arrray...
				// a good list implementation would be useful
				return run([]Fraction(p), append([]int{int(x)}, n...))
			case Failure:
				return run(fracs[1:], n)
			}
			// the compiler does not know / ensure that we have
			// exhausted all patterns in the matching. this is
			// why we need something here.
			panic("should never get here")
		}
	}

	return run(p, []int{input})
}

// Examples taken from Sheet 1 of Functional Programming course
func sumFn(f func(int) int, x, y int) int {
	switch {
	case x == y:
		return f(x)
	default:
		return f(x) * sumFn(f, x+1, y)
	}
}

// there's no generic map implementation, so we'd need to define a
// map function for _every_ type combination. Currently, things like
// this (maps (e.g. map[string]int), and slices are done within the
// compiler. A generic map implementation would need to be done in
// the compiler too, or with Go 2 (I _think_ there is an example
// implementation of the Go 2 Generics proposal somewhere, but absolutely
// alpha...).
// A generic map implementation would have the function header
//     func map(x []<T1>, func(<T1>) <T2>) []<T2>
// to give an example (named for clarity):
func MapStringToInt(x []string, f func(string) int) []int {
	var l []int
	for _, elem := range x {
		l = append(l, f(elem))
	}
	return l
}

// there is also no list implementation. There are slices ("views"
// on arrays), but they are 'arrays on steroids'. The only function
// that exists for them is append, which adds to a slice / array.
// For functional programming, we could at least need some helper
// functions on slices; prepend (Haskells 'a:l' -> l = prepend(a, l)),
// reverse, head, last?
// prepend _could_ be defined like this (without regard to runtime,
// and only for one specific type, because again, no generics):
func Prepend(x int, l []int) []int {
	return append([]int{x}, l...)
}
