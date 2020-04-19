package main

import "fmt"

func main() {
	m := MyStruct{
		x: "struct",
		y: 42,
	}

	fmt.Println(m) // {struct 42}
	mutateNoPointer(m)
	fmt.Println(m) // {struct 42}
	mutatePointer(&m)
	fmt.Println(m) // {changed 0}
}

type MyStruct struct {
	x string
	y int
}

func mutateNoPointer(m MyStruct) {
	m.x = "changed"
	m.y = 0
}
func mutatePointer(m *MyStruct) {
	m.x = "changed"
	m.y = 0
}
