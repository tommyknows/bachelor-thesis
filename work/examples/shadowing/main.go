package main

import "fmt"

func main() {
	x := 5
	fmt.Println(x) // 5
	// introducing a new scope
	{
		// this assignment would be forbidden, as it
		// overwrites the parent block's value value.
		x = 3
		fmt.Println(x) // 3
	}
	fmt.Println(x) // 3
	// introducing a new scope
	{
		// this redeclares the variable x, effectively
		// shadowing it. This will not change the parent
		// block's variable.
		x := 4
		fmt.Println(x) // 4
	}
	fmt.Println(x) // 3
}
