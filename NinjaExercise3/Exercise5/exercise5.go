package main

import "fmt"

func main() {

	for x := 10; x <= 100; x++ {
		fmt.Println("Value:", x, "\tMod 4:", x%4)
	}
}
