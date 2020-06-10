package main

import "fmt"

func main() {
	xii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := foo(xii...)
	s2 := bar(xii)

	fmt.Println(s1, s2)
}

func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
