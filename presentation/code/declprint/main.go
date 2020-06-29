package main

import (
	"fmt"
	"os/exec"
)

func reassignments() {
	// START OMIT
	a, b := 5, 1
	a, c := 5, 2
	// END OMIT
	fmt.Println(a, b, c)
}

func main() {
	cmd := exec.Command("/go/bin/prettyprint", "./code/declprint")
	out, err := cmd.CombinedOutput()

	fmt.Printf("%s\n", out[:findEnd(out)])
	if err != nil {
		fmt.Printf("%s", err)
	}
}

func findEnd(b []byte) int {
	var occurrence int
	for i, r := range b {
		if r == '\n' {
			occurrence++
		}
		if occurrence == 10 {
			return i
		}
	}
	return -1
}
