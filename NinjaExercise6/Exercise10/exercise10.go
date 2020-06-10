package main

import "fmt"

func main() {

	f := foo()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	g := foo()
	for i := 0; i < 10; i++ {
		fmt.Println(g())
	}

}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
