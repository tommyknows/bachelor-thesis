package main

import (
	"fmt"
	"strconv"
)

type parity bool

const (
	even parity = true
	odd  parity = false
)

// shouldBe returns a function that returns true if an int is of the
// given parity
func shouldBe(p parity) func(i int) bool {
	return func(i int) bool {
		return (i%2 == 0) == p
	}
}

func main() {
	l := []int{1, 2, 3, 4, 5}
	l5 := fmap(func(i int) int { return i * 5 }, prepend(0, l))

	// fold over even / odd numbers and add them to a string
	evens := foldl(
		func(s string, i int) string { return s + strconv.Itoa(i) + " " },
		"even: ",
		filter(shouldBe(even), l5),
	)
	odds := foldl(
		func(s string, i int) string { return s + strconv.Itoa(i) + " " },
		"odd: ",
		filter(shouldBe(odd), l5),
	)

	fmt.Println(evens, odds)
}
