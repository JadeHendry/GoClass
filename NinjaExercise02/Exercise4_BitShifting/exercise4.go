package main

import "fmt"

func main() {
	var x int = 1
	fmt.Printf("%d\t%b\t%#X\n", x, x, x)
	var y int = x << 1
	fmt.Printf("%d\t%b\t%#X", y, y, y)

}
