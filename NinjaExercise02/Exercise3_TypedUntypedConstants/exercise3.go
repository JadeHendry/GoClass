package main

import "fmt"

const (
	//typed constant
	x = 42
	//untyped constant
	y int = 43
)

func main() {
	fmt.Printf("%T\t%v\n", x, x)
	fmt.Printf("%T\t%v", y, y)
}
