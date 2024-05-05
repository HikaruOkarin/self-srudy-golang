package main

import "fmt"

type Points struct {
	x, y int
	s    string
}

func main() {
	p := Points{
		x: 12,
		y: 24,
		s: "hello world",
	}
	p.print()
}

func (p Points) print() {
	fmt.Print(p.x)
}
