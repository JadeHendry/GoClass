package main

import "fmt"

func main() {
	x := [5]int{}
	x[0] = 10
	x[1] = 20
	x[2] = 30
	x[3] = 40
	x[4] = 50

	for v, i := range x {
		fmt.Println(v, i)
	}

	fmt.Printf("%T", x)
}
