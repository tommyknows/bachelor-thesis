package main

import (
	"fmt"
	"os/exec"
)

func reassignments() {
	// START OMIT

	//f := func() {
	//	f() // undeclared name: f
	//}

	var f func()
	f = func() {
		f()
	}
	// END OMIT
}
func main() {
	cmd := exec.Command("/go/bin/funcheck", "./code/reassignments/funcliterals")
	out, err := cmd.CombinedOutput()
	fmt.Printf("%s\n%s", out, err)
}
