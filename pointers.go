package main

import "fmt"

func main() {
	pointers()
}

func pointers() {
	a := "hello world"
	b := 42
	fmt.Println(a)
	fmt.Println(b)
	p := &a
	fmt.Println(p)
	fmt.Println(*p)
	g := &b
	*g = *g / 2
	fmt.Println(b)
}
