package main

import "fmt"

const (
	a = (2020 + iota + 1)
	b = (2020 + iota + 1)
	c = (2020 + iota + 1)
	d = (2020 + iota + 1)
)

func main() {
	fmt.Print(a, b, c, d)
}
