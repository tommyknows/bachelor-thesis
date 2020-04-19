package main

import "fmt"

func main() {
	zero, one, two, three := 0, 1, 2, 3
	list := []*int{&one, &two, nil, &three, &zero}

	// this will work, as the values will be evaluated
	// "lazily" - the nested functions will never
	// be executed, thus it will never panic.
	fmt.Printf("%v\n", myFold(mulLazy, 1, list))
	// This will panic.
	fmt.Printf("%v\n", foldl(mul, 1, list))
}

func mul(x int, y *int) int {
	if *y == 0 {
		return 0
	}
	return x * *y
}

func mulLazy(x func() int, y *int) func() int {
	return func() int {
		if *y == 0 {
			return 0
		}
		return x() * *y
	}
}

func myFold(f func(func() int, *int) func() int, acc int, list []*int) int {
	a := func() int { return acc }
	a = foldl(f, a, list)
	return a()
}
