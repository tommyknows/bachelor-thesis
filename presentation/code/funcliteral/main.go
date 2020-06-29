package main

func main() {
	// START OMIT
	f := func() {
		f() // invalid: f declared but not used
	}

	var g func()
	g = func() {
		g() // valid: g has been declared
	}
	// END OMIT
}
