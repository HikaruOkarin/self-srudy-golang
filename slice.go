package main

import "fmt"

func main() {
	array()
}

type SLice struct {
	ptr      *int
	len, cap int
}

func array() {
	var array1 []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	name1 := array1[2:7]

	fmt.Println(array1)
	fmt.Println(name1)

	fmt.Println(len(array1), cap(array1))
	fmt.Println(len(name1), cap(name1))
	// d := make([]int, 10)
	// fmt.Println(len(d), cap(d))
	// d = append(d, 2)
	// fmt.Println(len(d), cap(d))
}
