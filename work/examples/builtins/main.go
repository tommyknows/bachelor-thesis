package main

import (
	"fmt"
	"strconv"
)

// Counterpart to Haskell's `derive Show`
//go:generate stringer -type parity

type parity int

const even parity = 0
const odd parity = 1

// shouldBe returns a function that returns true if an int is of the
// given parity
func shouldBe(p parity) func(i int) bool {
	return func(i int) bool {
		return i%2 == int(p)
	}
}

func main() {
	l := []int{1, 2, 3, 4, 5}
	l5 := fmap(func(i int) int { return i * 5 }, prepend(0, l))

	addToString := func(s string, i int) string {
		return s + strconv.Itoa(i) + " "
	}
	// fold over even / odd numbers and add them to a string
	evens := foldl(addToString, even.String()+": ",
		filter(shouldBe(even), l5))
	odds := foldl(addToString, odd.String()+": ",
		filter(shouldBe(odd), l5))

	fmt.Println(evens, odds)
}
