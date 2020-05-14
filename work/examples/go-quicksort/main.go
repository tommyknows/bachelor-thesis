package main

import "fmt"

func main() {
	fmt.Println(quicksort([]int{1, 8, 5, 3, 4, 9}))
}

// start-quicksort
func quicksort(p []int) []int {
	if len(p) == 0 {
		return []int{}
	}

	lesser := filter(func(x int) bool { return p[0] > x }, p[1:])
	greater := filter(func(x int) bool { return p[0] <= x }, p[1:])

	return append(quicksort(lesser), prepend(p[0], quicksort(greater))...)
}
// end-quicksort
