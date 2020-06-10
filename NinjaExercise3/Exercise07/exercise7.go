package main

import "fmt"

func main() {
	x := 100

	if x == 101 {
		fmt.Println("oh yeah")
	} else if x == 100 {
		fmt.Println("oh yes too")
	} else {
		fmt.Println("why not")
	}
}
