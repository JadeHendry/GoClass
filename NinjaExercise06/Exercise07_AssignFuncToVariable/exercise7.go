package main

import "fmt"

func main() {
	x := func() {
		for i := 1; i <= 100; i++ {
			fmt.Println(i)
		}
	}

	x()

	fmt.Println("Done.")
}
