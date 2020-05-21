package list

func listOps(list []int) {
	// collect all elements in a list and increment
	// each element by one
	var incremented []int
	for _, n := range list {
		incremented = append(incremented, n+1)
	}

	// same, but with new builtin
	incremented2 := fmap(
		func(x int) int { return x + 1 },
		list)

	// scale each element by factor s
	// and add t, collect results in
	// new list
	s := 2
	t := 1
	scaled := fmap(
		func(x int) int { return s*x + t },
		list)

	// filter all even elements and
	// collect them in a new list
	even := filter(
		func(x int) bool { return x%2 == 0 },
		list)

	// sum all elements in the list
	sum_reduce := 0
	for _, n := range list {
		sum_reduce += n
	}

	// same, but with builtin
	sum_reduce2 := foldl(
		func(a, b int) int { return a + b },
		0, list)
}
