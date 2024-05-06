package main

import "fmt"

type arr []int

func (p *arr) app(n int) {
	*p = append(*p, n)
}

func main() {
	var id arr = arr{1, 2}
	// id.app(3)
	// fmt.Println(id)
	var ptr *arr = &id
	id = append(id, 123)
	fmt.Println(ptr)
}
