package main

import "fmt"

func main() {
	pointers()
	h := "some text"
	var ptr *string = &h
	fmt.Println(ptr)
	var ptr2 *string = &h
	fmt.Println(ptr2)
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
