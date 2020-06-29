package main

import (
	"fmt"
	"os/exec"
)

func reassignments() {
	// START OMIT
	x, err := do()
	if err != nil {
		panic(err)
	}

	y, err := do()
	if err != nil {
		panic(err)
	}

	_ = "hello" // blank identifiers

	var a float64 = 5

	// END OMIT
	fmt.Println(a, x, y)
}

func do() (int, error) {
	return 0, nil
}

func main() {
	cmd := exec.Command("/go/bin/funcheck", "./code/reassignments/issues")
	out, err := cmd.CombinedOutput()
	fmt.Printf("%s\n%s", out, err)
}
