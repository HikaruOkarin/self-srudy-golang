package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	structs()
}

func structs() {
	var p1 Point = Point{100, 5}
	p := Point{x: 12, y: 4545}
	fmt.Println(p)

	var ptr2 *Point = &p
	ptr2.y = 667
    
	fmt.Println(p.x)
	fmt.Println(p1.x, p1.y)

	fmt.Println(p1)
	fmt.Println(p)
	p3 := Point{
		x: 123,
	}
	fmt.Println(p3)
}

func change(ptr *Point) {
	ptr.x = 100
}
