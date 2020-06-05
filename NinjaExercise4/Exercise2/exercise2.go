package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for v, i := range x {
		fmt.Println(v, i)
	}

	fmt.Printf("%T", x)
}
