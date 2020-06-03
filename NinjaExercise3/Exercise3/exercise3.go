package main

import "fmt"

func main() {
	i := 0
	for i < 24 {
		i++
		fmt.Println("Year:", i)
	}

	fmt.Println("I am", i, "years old!")
}
