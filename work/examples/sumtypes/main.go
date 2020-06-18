package main

import "fmt"

func main() {
	x := get(false)

	switch m := x.(type) {
	case myInt:
		fmt.Printf("Is integer: %v\n", m)
	case myString:
		fmt.Printf("Is string: %s\n", m)
	}
}

func get(b bool) MySumType {
	if b {
		return myInt(0)
	}
	return myString("zero")
}

type MySumType interface {
	mysumtype()
}

type myInt int

func (m myInt) mysumtype() {}

type myString string

func (m myString) mysumtype() {}
