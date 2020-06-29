package main

import (
	"fmt"
	"os/exec"
)

func reassignments() {
	// START OMIT

	var a int
	a = 5

	var b int = 5

	c := 5

	a = 6
	a |= 3
	a += 7

	a++

	// END OMIT
	fmt.Println(a, b, c)
}

func main() {
	cmd := exec.Command("/go/bin/funcheck", "./code/reassignments/")
	out, err := cmd.CombinedOutput()
	fmt.Printf("%s\n%s", out, err)
}
