package main

import "fmt"

func main() {
	g := func() {
		fmt.Println("This is func g.")
	}

	superFunc(g)
}

func superFunc(f func()) {
	fmt.Println("This is Super Func.")
	f()
}
