package main

import (
	"fmt"
	"os/exec"
)

func reassignments() {
	var x int
	// START OMIT
	for {

	}
	for x != 0 {

	}
	for i, elem := range []int{1, 2, 3} {
		fmt.Println(i, elem)
	}
	for x := 0; x < 10; x++ {

	}
	// END OMIT
}
func main() {
	cmd := exec.Command("/go/bin/funcheck", "./code/reassignments/for")
	out, err := cmd.CombinedOutput()
	fmt.Printf("%s\n%s", out, err)

}
